package main

import (
	"fmt"
)

type Board struct {
	grid *[8][8]Movable
}

// Returns piece in given pos and wether it's occupied or not.
func (b *Board) GetPiece(pos Position) (Movable, bool) {
	piece := b.grid[pos.Row][pos.Col]
	return piece, b.IsOccupied(pos)
}

// Inserts piece in piece.Pos
func (b *Board) InsertPiece(piece Movable) bool {
	if b.IsOccupied(piece.GetPosition()) {
		return false
	}

	b.grid[piece.GetPosition().Row][piece.GetPosition().Col] = piece
	return true
}

func (b *Board) InsertPieces(pieces []Movable) {
	for _, v := range pieces {
		b.InsertPiece(v)
	}
}

// Moves piece from piece.Pos to pos
func (b *Board) MovePiece(piece Movable, pos Position) Movable {
	currPos := piece.GetPosition()

	capture, _ := b.GetPiece(pos)

	b.grid[pos.Row][pos.Col] = piece
	b.grid[currPos.Row][currPos.Col] = nil

	piece.SetPosition(pos)

	return capture
}

func (b *Board) IsOccupied(pos Position) bool {
	return b.grid[pos.Row][pos.Col] != nil
}

func (b *Board) String() string {
	output := ""

	for row := 7; row >= 0; row-- { // Mostrar del 8 al 1
		output += fmt.Sprintf("%d ", row+1)
		for col := 0; col < 8; col++ {
			piece := b.grid[row][col]
			if piece != nil {
				output += fmt.Sprintf("%-3s", piece.String())
			} else {
				output += fmt.Sprintf("%-3s", "--")
			}
		}
		output += "\n"
	}

	output += "   A  B  C  D  E  F  G  H\n"

	return output
}
