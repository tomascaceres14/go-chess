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

	game.MovePiece(Pos("D", 2), Pos("D", 4), p1)
	game.MovePiece(Pos("E", 5), Pos("D", 4), p2)

	fmt.Println(game.board)

	fmt.Println("White Pieces:", game.PWhite.Pieces)
	fmt.Println("Black Pieces:", game.PBlack.Pieces)
}
