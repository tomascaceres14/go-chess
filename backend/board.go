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

// Used for simulating movement when cloning board.
func (b *Board) MovePieceSim(piece Movable, pos Position) {
	prevPos := piece.GetPosition()
	b.grid[prevPos.Row][prevPos.Col] = nil
	b.grid[pos.Row][pos.Col] = piece
}

// Moves piece from piece.Pos to pos and returns captured piece
func (b *Board) MovePiece(piece Movable, pos Position) Movable {
	prevPos := piece.GetPosition()

	capture, _ := b.GetPiece(pos)

	b.grid[prevPos.Row][prevPos.Col] = nil
	b.grid[pos.Row][pos.Col] = piece

	piece.SetPosition(pos)

	return capture
}

func (b *Board) IsOccupied(pos Position) bool {
	return b.grid[pos.Row][pos.Col] != nil
}

func (b *Board) Clone() *Board {
	grid := *b.grid
	return &Board{grid: &grid}
}

func (b *Board) IsChecked(p *Player) bool {
	return p.Threats[p.King.Pos]
}

func (b *Board) IsRowPathClear(from, to Position) bool {
	if from.Row != to.Row {
		return false
	}

	step := 1
	if from.Col > to.Col {
		step = -1
	}

	for col := from.Col + step; col != to.Col; col += step {
		if b.IsOccupied(Position{Row: from.Row, Col: col}) {
			return false
		}
	}

	return true
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
