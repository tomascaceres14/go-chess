package main

import "fmt"

func main() {
	p1 := NewPlayer("White", true)
	p2 := NewPlayer("Black", false)

	game := NewGame(p1, p2)

	if err := game.MovePiece(Pos("E", 7), Pos("E", 5), p2); err != nil {
		PrintError(err)
	}

	if err := game.MovePiece(Pos("F", 1), Pos("C", 4), p1); err != nil {
		PrintError(err)
	}

	if err := game.MovePiece(Pos("F", 3), Pos("F", 7), p1); err != nil {
		PrintError(err)
	}

	if err := game.MovePiece(Pos("G", 1), Pos("F", 3), p1); err != nil {
		PrintError(err)
	}

	if err := game.MovePiece(Pos("B", 1), Pos("C", 3), p1); err != nil {
		PrintError(err)
	}

	if err := game.MovePiece(Pos("E", 7), Pos("E", 5), p2); err != nil {
		PrintError(err)
	}

	fmt.Println(game.board)
}
