package engine

import (
	"errors"
	"fmt"
	"strings"
)

var games map[string]*Game

// Error printing for debugging
func PrintError(err error) {
	fmt.Printf("--- ERROR: %v\n", err)
}

func StartGame(whiteName, blackName string) (string, error) {

	if len(games) <= 5 {
		return "", errors.New("Maximum of 5 simultaneous games reached")
	}

	whiteName = strings.TrimSpace(whiteName)
	blackName = strings.TrimSpace(blackName)

	pWhite := NewPlayer(whiteName, true)
	pBlack := NewPlayer(blackName, false)

	game := NewGame(pWhite, pBlack)

	games[game.id] = game

	return game.id, nil
}

func MakeMove(id string, player *Player, from, to Position, playerColor bool) error {

	// Validate id belonging to player and game state = playing

	game := games[id]

	if err := game.MovePiece(from, to, player); err != nil {
		return err
	}

	return nil
}
