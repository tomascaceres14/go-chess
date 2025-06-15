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

	// if err := game.MovePiece(Pos(1, "B"), Pos(3, "A"), p1); err != nil {
	// 	PrintError(err)
	// }

	if err := game.MovePiece(Pos(1, "B"), Pos(4, "C"), p1); err != nil {
		PrintError(err)
	}

	fmt.Println(game.board)

}
