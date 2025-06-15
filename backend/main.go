package main

import (
	"fmt"
)

func main() {

	//color := randomBool()
	//p1 := NewPlayer("P1", color)
	//p2 := NewPlayer("P2", !color)
	p1 := NewPlayer("White", true)
	p2 := NewPlayer("Black", false)

	game := NewGame(p1, p2)

	fmt.Println(game.board)

	G7Pawn, err := game.GetPiece(Pos(7, "G"), p2)
	if err != nil {
		PrintError(err)
	}

	fmt.Println(G7Pawn)
	game.MovePiece(G7Pawn, Pos(5, "G"), p2)

	fmt.Println(game.board)

	game.MovePiece(G7Pawn, Pos(6, "G"), p2)
	fmt.Println(game.board)
	// fmt.Println(WBishop.PossibleMoves(game.board))
}
