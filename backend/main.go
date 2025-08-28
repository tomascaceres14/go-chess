package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("White player's name: ")
	p1Name, _ := reader.ReadString('\n')
	p1Name = strings.TrimSpace(p1Name)

	fmt.Print("Black player's name: ")
	p2Name, _ := reader.ReadString('\n')
	p2Name = strings.TrimSpace(p2Name)

	p1 := NewPlayer(p1Name, true)
	p2 := NewPlayer(p2Name, false)

	game := NewGame(p1, p2)

	currPlayer := p1
	fmt.Println(game.board)

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

		fmt.Println(game.board)

		currPlayer = SwitchTurns(p1, p2, game.WhiteTurn)
	}
}
