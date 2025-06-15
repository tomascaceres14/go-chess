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

	game.MovePiece(Pos(7, "G"), Pos(5, "G"), p2)
	game.MovePiece(Pos(2, "D"), Pos(3, "D"), p1)
	game.MovePiece(Pos(1, "C"), Pos(5, "G"), p1)

	fmt.Println(game.board)
}
