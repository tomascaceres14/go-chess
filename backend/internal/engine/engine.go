package engine

import (
	"errors"
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

	return game.id, nil
}

func (e *ChessEngine) MakeMove(from, to position, pColor bool) error {

	if e.game == nil {
		return errors.New("Game not initialized.")
	}

	if err := e.game.movePiece(from, to, pColor); err != nil {
		return err
	}

	return nil
}
