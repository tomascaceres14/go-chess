package main

import "fmt"

func main() {
	p1 := NewPlayer("White", true)
	p2 := NewPlayer("Black", false)

	game := NewGame(p1, p2)

	pawn, err := game.GetPiece(Pos("E2"), p1)
	if err != nil {
		PrintError(err)
	}

	fmt.Println(pawn.LegalMoves(game.board))

	if err := game.MovePiece(Pos("a2"), Pos("a4"), p1); err != nil {
		PrintError(err)
	}

	if err := game.MovePiece(Pos("f7"), Pos("f6"), p2); err != nil {
		PrintError(err)
	}

	if err := game.MovePiece(Pos("e5"), Pos("f6"), p1); err != nil {
		PrintError(err)
	}

	fmt.Println(game.board)
}
