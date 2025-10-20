package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tomascaceres14/go-chess/backend/internal/engine"
)

// Error printing for debugging
func PrintError(err error) {
	fmt.Printf("--- ERROR: %v\n", err)
}

func SwitchTurns(p1, p2 *engine.Player, white bool) *engine.Player {
	if p1.White == white {
		return p1
	}

	return p2
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("White player's name: ")
	p1Name, _ := reader.ReadString('\n')
	p1Name = strings.TrimSpace(p1Name)

	fmt.Print("Black player's name: ")
	p2Name, _ := reader.ReadString('\n')
	p2Name = strings.TrimSpace(p2Name)

	p1 := engine.NewPlayer(p1Name, true)
	p2 := engine.NewPlayer(p2Name, false)

	game := engine.NewGame(p1, p2)

	currPlayer := p1

	for true {
		fmt.Printf("\n%s's turn! (%s)\n", currPlayer.Name, engine.ColorToString(currPlayer.White))

		fmt.Print("From: ")
		from, _ := reader.ReadString('\n')
		from = strings.TrimSpace(from)

		fmt.Print("To: ")
		to, _ := reader.ReadString('\n')
		to = strings.TrimSpace(to)

		if err := game.MovePiece(engine.Pos(from), engine.Pos(to), currPlayer); err != nil {
			PrintError(err)
			continue
		}

		currPlayer = SwitchTurns(p1, p2, game.WhiteTurn)
	}
}
