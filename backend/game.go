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

	// Pawns
	for i := range 8 {
		blackRookPos := Position{Row: 6, Col: i}
		whiteRookPos := Position{Row: 1, Col: i}
		board[6][i] = NewPawn(false, blackRookPos) // black
		board[1][i] = NewPawn(true, whiteRookPos)  // white
	}

	// Rooks
	board[7][0] = NewRook(false, Position{Row: 7, Col: 0}) // black
	board[7][7] = NewRook(false, Position{Row: 7, Col: 7}) // black

	board[0][0] = NewRook(true, Position{Row: 0, Col: 0}) // white
	board[0][7] = NewRook(true, Position{Row: 0, Col: 7}) // white

	return &Game{
		board: &Board{grid: &board},
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
	if !pos.InBounds() {
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
