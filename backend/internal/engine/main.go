package engine

import (
	"errors"
	"fmt"
	"time"
	"strings"
)

var games map[string]*Game

// Error printing for debugging
func PrintError(err error) {
	fmt.Printf("--- ERROR: %v\n", err)
}

func SwitchTurns(p1, p2 *Player, white bool) *Player {
	if p1.White == white {
		return p1
	}

	return p2
}

func StartGame(whiteName, blackName string) (string, error) {

	if len(games) <= 5 {
		return "", errors.New("Maximum of 5 simultaneous games reached")
	}

	whiteName = strings.TrimSpace(whiteName)
	blackName = strings.TrimSpace(blackName)

	pWhite := NewPlayer(whiteName, true)
	pBlack := NewPlayer(blackName, false)

	now := time.Now()
	timestamp := now.Format("20060201150405")

	id := whiteName + "_" + blackName + "_" + timestamp
	
	if _, exists := games[id]; exists {
		return id, nil
	}

	games[id] = NewGame(id, pWhite, pBlack)

	return id, nil
}

func MakeMove(id, from, to string, playerColor bool) {

	// Validate id belonging to player and game state = playing

	game := games[id]
	
}
