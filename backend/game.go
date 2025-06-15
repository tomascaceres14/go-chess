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

	game := &Game{
		board: &Board{grid: &board},
		p1:    p1,
		p2:    p2,
	}

	// Pawns
	for i := range 8 {
		game.board.InsertPiece(NewPawn(false, Position{Row: 6, Col: i})) // black
		game.board.InsertPiece(NewPawn(true, Position{Row: 1, Col: i}))  // white
	}

	// Rooks
	game.board.InsertPiece(NewRook(false, Position{Row: 7, Col: 0})) // black
	game.board.InsertPiece(NewRook(false, Position{Row: 7, Col: 7})) // black
	game.board.InsertPiece(NewRook(true, Position{Row: 0, Col: 0}))  // white
	game.board.InsertPiece(NewRook(true, Position{Row: 0, Col: 7}))  // white

	// Bishops
	game.board.InsertPiece(NewBishop(false, Position{Row: 7, Col: 2})) // black
	game.board.InsertPiece(NewBishop(false, Position{Row: 7, Col: 5})) // black
	game.board.InsertPiece(NewBishop(true, Position{Row: 0, Col: 2}))  // white
	game.board.InsertPiece(NewBishop(true, Position{Row: 0, Col: 5}))  // white

	return game
}

func (g *Game) getPiece(pos Position, player *Player) (Movable, error) {

	piece := g.board.GetPiece(pos)

	if piece.IsWhite() != player.White {
		return nil, fmt.Errorf("Not your piece, %s.", player.Name)
	}

	return piece, nil
}

func (g *Game) MovePiece(from, to Position, player *Player) error {
	// asegurar que pos este dentro del tablero
	if !to.InBounds() {
		return errors.New("Position out of bounds.")
	}

	piece, err := g.getPiece(from, player)
	if err != nil {
		return err
	}

	// verificar si pieza puede moverse a pos
	if !ContainsPosition(piece.PossibleMoves(g.board), to) {
		return fmt.Errorf("%s cant move to row=%d col=%d.", piece.String(), to.Row, to.Col)
	}

	if piece.GetPosition() == to {
		return errors.New("Cannot move to the same position.")
	}

	g.board.MovePiece(piece, to)

	return nil
}
