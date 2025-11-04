package engine

import (
	"fmt"
	"strings"
)

type Engine struct {
	game *Game
}

func NewEngine() *Engine {
	return &Engine{game: &Game{}}
}

// Error printing for debugging
func PrintError(err error) {
	fmt.Printf("--- ERROR: %v\n", err)
}

func (e *Engine) StartGame(whiteName, blackName string) (string, error) {

	whiteName = strings.TrimSpace(whiteName)
	blackName = strings.TrimSpace(blackName)

	game := NewGame(whiteName, blackName)
	e.game = game

	return game.id, nil
}

func (e *Engine) GetGame() *Game {
	return e.game
}

func (e *Engine) MakeMove(from, to Position, pColor bool) error {

	if err := e.game.MovePiece(from, to, pColor); err != nil {
		return err
	}

	return nil
}
