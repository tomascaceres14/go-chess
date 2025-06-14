package main

import (
	"fmt"
)

type Player struct {
	Name   string
	White  bool
	Points int
}

func NewPlayer(name string, isWhite bool) *Player {
	return &Player{
		Name:   name,
		White:  isWhite,
		Points: 0,
	}
}

func main() {

	//p1 := NewPlayer("P1", randomBool())
	//p2 := NewPlayer("P2", randomBool())
	p1 := NewPlayer("White", true)
	p2 := NewPlayer("Black", false)

	game := NewGame(p1, p2)

	fmt.Println(game.board)

	pawn, err := game.GetPiece(Position{Row: 6, Col: 2}, p2)
	if err != nil {
		PrintError(err)
	}

	fmt.Println(pawn)

	if err := game.MovePiece(pawn, Position{Row: 7, Col: 2}, p2); err != nil {
		PrintError(err)
	}

	fmt.Println(game.board)
}
