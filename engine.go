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

	fmt.Println("FEN STRING:", game.GetFENString())

	return game.id, nil
}

func (e *ChessEngine) NewGameFENString(FENString string, whiteName, blackName string) (*game, error) {

	game, err := newGameFENString(FENString, whiteName, blackName)
	if err != nil {
		return nil, err
	}

	return game, nil
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
