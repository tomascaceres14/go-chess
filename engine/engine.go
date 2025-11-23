package gochess

import (
	"errors"
	"fmt"
)

type ChessEngine struct {
	game *game
}

func NewChessEngine() *ChessEngine {
	return &ChessEngine{}
}

func (e *ChessEngine) NewGame(whiteName, blackName string) (string, error) {

	game, err := newGameClassic(whiteName, blackName)
	if err != nil {
		return "", err
	}

	e.game = game

	return game.id, nil
}

func (e *ChessEngine) NewGameFENString(whiteName, blackName string, FENString string) (string, error) {

	game, err := newGameFENString(FENString, whiteName, blackName)
	if err != nil {
		return "", err
	}

	e.game = game

	return game.id, nil
}

func (e *ChessEngine) Move(from, to string, pColor bool) ([8][8]Movable, error) {
	if e.game == nil {
		return [8][8]Movable{}, errors.New("Game not initialized.")
	}

	fromPos := pos(from)
	if !fromPos.isValid() {
		return [8][8]Movable{}, fmt.Errorf("%s", "'"+from+"' is an invalid position.")
	}

	toPos := pos(to)
	if !toPos.isValid() {
		return [8][8]Movable{}, fmt.Errorf("%s", "'"+to+"' is an invalid position.")
	}

	if err := e.game.makeMove(fromPos, toPos, pColor); err != nil {
		return [8][8]Movable{}, err
	}

	return *e.game.gameBoard.GetGrid(), nil
}

func (e *ChessEngine) GetFENString() string {
	return e.game.GetFENString()
}

func (e *ChessEngine) GetTurn() bool {
	return e.game.WhiteTurn
}

func (e *ChessEngine) GetBoard() Board {
	return e.game.gameBoard.clone()
}
