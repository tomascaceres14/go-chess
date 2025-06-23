package main

import "fmt"

func main() {
	p1 := NewPlayer("White", true)
	p2 := NewPlayer("Black", false)

	game := NewGame(p1, p2)

	game.MovePiece(Pos("C", 3), Pos("C", 4), p2)

	fmt.Println("ORIGINAL")
	fmt.Println(game.board)

	Pw, _ := game.board.GetPiece(Pos("D", 2))
	fmt.Println(game.MoveUnchecksPlayer(Pw, Pos("D", 4), p1, p2))
}
