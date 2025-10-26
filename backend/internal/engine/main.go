package engine

import (
	"fmt"
	"strings"
)

var game *Game

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

func StartGame(p1Name, p2Name string) {

	p1Name = strings.TrimSpace(p1Name)
	p2Name = strings.TrimSpace(p2Name)

	p1 := NewPlayer(p1Name, true)
	p2 := NewPlayer(p2Name, false)

	game = NewGame(p1, p2)

	currPlayer := p1

	for true {
		fmt.Printf("\n%s's turn! (%s)\n", currPlayer.Name, ColorToString(currPlayer.White))

		fmt.Print("From: ")
		from, _ := reader.ReadString('\n')
		from = strings.TrimSpace(from)

		fmt.Print("To: ")
		to, _ := reader.ReadString('\n')
		to = strings.TrimSpace(to)

		if err := game.MovePiece(Pos(from), Pos(to), currPlayer); err != nil {
			PrintError(err)
			continue
		}

		currPlayer = SwitchTurns(p1, p2, game.WhiteTurn)
	}
}
