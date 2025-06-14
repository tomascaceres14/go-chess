package main

type Game struct {
	board *Board
}

func NewGame() *Game {

	println("generating new board...")

	board := [8][8]Movable{}

	for i := range 8 {
		blackPos := Position{Row: 6, Col: i}
		whitePos := Position{Row: 1, Col: i}
		board[6][i] = NewPawn(false, blackPos) // black
		board[1][i] = NewPawn(true, whitePos)  // white
	}

	return &Game{board: &Board{grid: board}}
}
