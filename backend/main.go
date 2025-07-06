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

	if err := game.MovePiece(Pos("E7"), Pos("E5"), p2); err != nil {
		PrintError(err)
	}

	if err := game.MovePiece(Pos("F1"), Pos("C4"), p1); err != nil {
		PrintError(err)
	}

	if err := game.MovePiece(Pos("F3"), Pos("F7"), p1); err != nil {
		PrintError(err)
	}

	if err := game.MovePiece(Pos("G1"), Pos("F3"), p1); err != nil {
		PrintError(err)
	}

	if err := game.MovePiece(Pos("B1"), Pos("C3"), p1); err != nil {
		PrintError(err)
	}

	if err := game.MovePiece(Pos("E7"), Pos("E5"), p2); err != nil {
		PrintError(err)
	}

	fmt.Println(game.board)
}
