package main

import (
	"errors"
	"fmt"
)

type Game struct {
	board  *Board
	p1, p2 *Player
}

func NewGame(p1, p2 *Player) *Game {

	println("generating new board...")

	board := [8][8]Movable{}

	for i := range 8 {
		blackPos := Position{Row: 6, Col: i}
		whitePos := Position{Row: 1, Col: i}
		board[6][i] = NewPawn(false, blackPos) // black
		board[1][i] = NewPawn(true, whitePos)  // white
	}

	return &Game{
		board: &Board{grid: board},
		p1:    p1,
		p2:    p2,
	}
}

func (g *Game) GetPiece(pos Position, player *Player) (Movable, error) {

	piece := g.board.GetPiece(pos)

	if piece.IsWhite() != player.White {
		return nil, fmt.Errorf("Not your piece, %s.", player.Name)
	}

	return piece, nil
}

func (g *Game) MovePiece(piece Movable, pos Position, player *Player) error {
	// asegurar que pos este dentro del tablero
	if pos.Col < 0 || pos.Col > 7 || pos.Row < 0 || pos.Row > 7 {
		return errors.New("Position out of bounds.")
	}

	if player.White != piece.IsWhite() {
		return fmt.Errorf("Not your piece, %s.", player.Name)
	}

	// verificar si pieza puede moverse a pos
	if !ContainsPosition(piece.PossibleMoves(g.board), pos) {
		return fmt.Errorf("%s cant move to row=%d col=%d.", piece.String(), pos.Row, pos.Col)
	}

	currPos := piece.GetPosition()

	if pos == currPos {
		return errors.New("Cannot move to the same position.")
	}

	g.board.MovePiece(piece, pos)

	return nil
}
