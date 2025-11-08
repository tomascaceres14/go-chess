package engine

import (
	"errors"
	"fmt"
	"strings"
)

type ChessEngine struct {
	game *game
}

func NewChessEngine() *ChessEngine {
	return &ChessEngine{}
}

func (e *ChessEngine) StartGame(whiteName, blackName string) (string, error) {

	whiteName = strings.TrimSpace(whiteName)
	blackName = strings.TrimSpace(blackName)

	game := NewGame(whiteName, blackName)
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

	if err := e.game.movePiece(fromPos, toPos, pColor); err != nil {
		return err
	}

	fmt.Println("FEN STRING:", e.game.GetFENString())

	return nil
}
