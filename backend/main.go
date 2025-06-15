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

}
