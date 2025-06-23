package main

import "fmt"

func main() {

	//color := randomBool()
	//p1 := NewPlayer("P1", color)
	//p2 := NewPlayer("P2", !color)
	p1 := NewPlayer("White", true)
	p2 := NewPlayer("Black", false)

	game := NewGame(p1, p2)

	game.MovePiece(Pos("D", 7), Pos("D", 5), p2)

	// cloneBoard := game.board.Clone()
	// pieceClone, _ := cloneBoard.GetPiece(Pos("E", 2))
	// cloneBoard.MovePiece(pieceClone, Pos("E", 4))

	fmt.Println("ORIGINAL")
	fmt.Println(game.board)
	// fmt.Println("CLONE")
	// fmt.Println(cloneBoard)

	fmt.Println(p1)
}
