package main

import (
	"errors"
	"fmt"
)

type Game struct {
	board          *Board
	PWhite, PBlack *Player
	Captures       []Movable
}

// Generates a new board with classic chess configuration
func NewGame(whites, blacks *Player) *Game {

	println("generating new board...")

	board := [8][8]Movable{}

	game := &Game{
		board:    &Board{grid: &board},
		PWhite:   whites,
		PBlack:   blacks,
		Captures: []Movable{},
	}

	// Pawns
	for i := range 8 {
		game.board.InsertPiece(NewPawn(Pos(GetCol(i), 7), blacks)) // black
		game.board.InsertPiece(NewPawn(Pos(GetCol(i), 2), whites)) // white
	}

	// Rooks
	game.board.InsertPiece(NewRook(Pos("A", 8), blacks)) // black
	game.board.InsertPiece(NewRook(Pos("H", 8), blacks)) // black
	game.board.InsertPiece(NewRook(Pos("A", 1), whites)) // white
	game.board.InsertPiece(NewRook(Pos("H", 1), whites)) // white

	// Knights
	game.board.InsertPiece(NewKnight(Pos("B", 8), blacks)) // black
	game.board.InsertPiece(NewKnight(Pos("G", 8), blacks)) // black
	game.board.InsertPiece(NewKnight(Pos("B", 1), whites)) // white
	game.board.InsertPiece(NewKnight(Pos("G", 1), whites)) // white

	// Bishops
	game.board.InsertPiece(NewBishop(Pos("C", 8), blacks)) // black
	game.board.InsertPiece(NewBishop(Pos("F", 8), blacks)) // black
	game.board.InsertPiece(NewBishop(Pos("C", 1), whites)) // white
	game.board.InsertPiece(NewBishop(Pos("F", 1), whites)) // white

	// Queens
	game.board.InsertPiece(NewQueen(Pos("D", 8), blacks)) // black
	game.board.InsertPiece(NewQueen(Pos("D", 1), whites)) // white

	// Kings
	game.board.InsertPiece(NewKing(Pos("E", 8), blacks)) // black
	game.board.InsertPiece(NewKing(Pos("E", 1), whites)) // white

	return game
}

// Internal use only. Obtains piece at given position if player is owner of piece
func (g *Game) getPiece(pos Position, player *Player) (Movable, error) {

	piece, ok := g.board.GetPiece(pos)
	if !ok {
		return nil, fmt.Errorf("Position %v is empty.", pos)
	}

	// validar que pieza pertenezca a player
	if piece.IsWhite() != player.White {
		return nil, fmt.Errorf("Not your piece, %s.", player.Name)
	}

	return piece, nil
}

// Moves piece in position `from` to position `to` if player is owner of piece
func (g *Game) MovePiece(from, to Position, player *Player) error {
	// Check if positions are in bounds
	if !from.InBounds() || !to.InBounds() {
		return errors.New("Position out of bounds.")
	}

	// Obtains piece to move
	piece, err := g.getPiece(from, player)
	if err != nil {
		return err
	}

	// Check if piece can move to desired position
	possibleMoves := piece.PossibleMoves(g.board)
	if !possibleMoves[to] {
		return fmt.Errorf("%s cant move from %s to %s.", piece.String(), from, to)
	}

	// Check if it's trying to move in place
	if piece.GetPosition() == to {
		return errors.New("Cannot move to the same position.")
	}

	capture := g.board.MovePiece(piece, to)

	opponent := g.PBlack

	if opponent.White == player.White {
		opponent = g.PWhite
	}

	if capture != nil {
		opponent.Pieces = DeletePiece(opponent.Pieces, capture)
	}

	// Update threats map of opponent
	opponent.Threats = player.CalculateThreats(g.board)

	return nil
}
