package main

import "fmt"

func main() {
	p1 := NewPlayer("White", true)
	p2 := NewPlayer("Black", false)

	game := NewGame(p1, p2)

	if err := game.MovePiece(Pos("D", 2), Pos("D", 4), p1); err != nil {
		PrintError(err)
	}

	if err := game.MovePiece(Pos("E", 4), Pos("D", 5), p1); err != nil {
		PrintError(err)
	}

	fmt.Println(game.board)
}
