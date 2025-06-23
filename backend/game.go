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
	game.board.InsertPiece(NewQueen(Pos("C", 3), blacks)) // black
	game.board.InsertPiece(NewQueen(Pos("D", 1), whites)) // white

	// Kings
	bKing := NewKing(Pos("E", 8), blacks)
	blacks.King = bKing
	game.board.InsertPiece(bKing)

	wKing := NewKing(Pos("E", 4), whites)
	whites.King = wKing
	game.board.InsertPiece(wKing)

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

	// Obtains opponent
	opponent := g.PBlack
	if opponent.White == player.White {
		opponent = g.PWhite
	}

	// Obtains piece to move
	piece, err := g.getPiece(from, player)
	if err != nil {
		return err
	}

	// if player.Checked && g.MoveUnchecksPlayer(piece, to, player, opponent) {

	// }

	// Check if piece can move to desired position or if is trying to move in-place
	legalMoves := piece.LegalMoves(g.board)
	if !legalMoves[to] || piece.GetPosition() == to {
		return fmt.Errorf("%s can't move from %s to %s.", piece.String(), from, to)
	}

	capture := g.board.MovePiece(piece, to)

	if capture != nil {
		opponent.Pieces = DeletePiece(opponent.Pieces, capture)
	}

	attackedSquares := player.AttackedSquares(g.board)

	// Update threats map of opponent and flag as checked or not
	opponent.Threats = attackedSquares
	opponent.Checked = attackedSquares[opponent.King.Pos]

	return nil
}

func (g *Game) MoveUnchecksPlayer(piece Movable, to Position, player, oppponent *Player) bool {
	cloneBoard := g.board.Clone()
	cloneBoard.MovePiece(piece, to)
	return !oppponent.AttackedSquares(cloneBoard)[player.King.Pos]
}
