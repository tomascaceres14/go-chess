package main

import (
	"errors"
	"fmt"
)

type Board struct {
	grid *[8][8]Movable
}

func (b *Board) InsertPiece(piece Movable) bool {
	if b.isOccupied(piece.GetPosition()) {
		return false
	}

	b.grid[piece.GetPosition().Row][piece.GetPosition().Col] = piece
	return true
}

func (b *Board) MovePiece(piece Movable, pos Position) error {
	currPos := piece.GetPosition()

	if pos == currPos {
		return errors.New("Cannot move to the same position")
	}

	b.grid[pos.Row][pos.Col] = piece
	b.grid[currPos.Row][currPos.Col] = nil

	piece.SetPosition(pos)

	return nil
}

func (b *Board) GetPiece(pos Position) Movable {
	return b.grid[pos.Row][pos.Col]
}

func (b *Board) isOccupied(pos Position) bool {
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
