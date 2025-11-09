package gochess

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type ChessEngine struct {
	game *game
}

func NewChessEngine() *ChessEngine {
	return &ChessEngine{}
}

func (e *ChessEngine) NewGame(whiteName, blackName string) (string, error) {

	game, err := NewDefaultGame(whiteName, blackName)
	if err != nil {
		return "", err
	}

	e.game = game

	fmt.Println("FEN STRING:", game.GetFENString())

	return game.id, nil
}

func (e *ChessEngine) Move(from, to string, pColor bool) error {
	if e.game == nil {
		return errors.New("Game not initialized.")
	}

	fromPos := pos(from)
	if !fromPos.isValid() {
		return fmt.Errorf("%s", "'"+from+"' is an invalid position.")
	}

	toPos := pos(to)
	if !toPos.isValid() {
		return fmt.Errorf("%s", "'"+to+"' is an invalid position.")
	}

	if err := e.game.makeMove(fromPos, toPos, pColor); err != nil {
		return err
	}

	fmt.Println("FEN STRING:", e.game.GetFENString())

	return nil
}

func (e *ChessEngine) NewGameFromFENString(FENString string, whiteName, blackName string) (*game, error) {

	gameBoard := [8][8]movable{}

	// Generate ID for match
	now := time.Now()
	timestamp := now.Format("20060201150405")
	id := whiteName + "_" + blackName + "_" + timestamp

	game := game{
		id:            id,
		gameBoard:     &board{grid: &gameBoard},
		moveHistory:   []move{},
		status:        playing,
		halfmoveClock: 0,
	}

	if pWhite, err := newPlayerWhite(whiteName); err != nil {
		return nil, err
	} else {
		game.pWhite = pWhite
	}

	if pBlack, err := newPlayerBlack(blackName); err != nil {
		return nil, err
	} else {
		game.pBlack = pBlack
	}

	FENSplit := strings.Split(FENString, " ")
	if len(FENSplit) != 6 {
		return nil, errors.New("Error reading FEN String: Incorrect format. FEN string should contain 6 tramos")
	}

	FENPosition := strings.Split(FENSplit[0], "/")
	if len(FENPosition) != 8 {
		return nil, errors.New("Error reading FEN String: Incorrect format for position. Should contain 8 tramos")
	}

	if err := game.setupPositionFENString(FENPosition); err != nil {
		return nil, err
	}

	FENTurn := FENSplit[1]
	if len(FENTurn) != 1 || (FENTurn != "w" && FENTurn != "b") {
		return nil, errors.New("Incorrect format for turn. Should be 1 character, either 'b' or 'w'.")
	}

	game.WhiteTurn = FENTurn == "w"

	FENCastling := FENSplit[2]
	if err := game.setupCastlingFENString(FENCastling); err != nil {
		return nil, err
	}

	FENEnPassant := FENSplit[3]
	if err := game.setupEnPassantFENString(FENEnPassant, game.WhiteTurn); err != nil {
		return nil, err
	}

	fmt.Println(game.gameBoard)

	fmt.Println("white castling opt long", game.pWhite.king.longCastlingOpt)
	fmt.Println("white castling opt short", game.pWhite.king.shortCastlingOpt)
	fmt.Println("black castling opt long", game.pBlack.king.longCastlingOpt)
	fmt.Println("black castling opt short", game.pBlack.king.shortCastlingOpt)

	fmt.Println("Current turn:", game.WhiteTurn)

	

	return &game, nil
}
