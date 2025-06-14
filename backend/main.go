package main

import "fmt"

type Player struct {
	name   string
	white  bool
	points int
}

func main() {

	g := NewGame()

	fmt.Println(g.board)

	pawn := g.board.GetPiece(Position{Row: 1, Col: 2})
	fmt.Println(pawn)

	if err := g.board.MovePiece(pawn, Position{Row: 3, Col: 2}); err != nil {
		fmt.Printf("--- ERROR: %v\n", err)
	}

	fmt.Println(g.board)
}
