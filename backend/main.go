package main

import (
	"fmt"
)

func main() {

	//p1 := NewPlayer("P1", randomBool())
	//p2 := NewPlayer("P2", randomBool())
	p1 := NewPlayer("White", true)
	p2 := NewPlayer("Black", false)

	game := NewGame(p1, p2)

	fmt.Println(game.board)

	getPos, err := Pos(2, "E")
	if err != nil {
		PrintError(err)
	}

	pawn, err := game.GetPiece(getPos, p1)
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Println(pawn)

	movePos, err := Pos(4, "e")
	if err != nil {
		PrintError(err)
		return
	}

	if err := game.MovePiece(pawn, movePos, p1); err != nil {
		PrintError(err)
		return
	}

	fmt.Println(game.board)
}
