package gochess

import (
	"fmt"
)

type board struct {
	grid *[8][8]movable
}

// Get piece at given pos. Returns piece and wether square is occupied or not.
// (should refactor, second bool not necessary as piece can be nil, not necessarily zero value)
func (b *board) getPiece(pos position) (movable, bool) {
	piece := b.grid[pos.Row][pos.Col]
	return piece, b.IsOccupied(pos)
}

// Inserts piece in piece.Pos
func (b *board) insertPiece(piece movable) {
	b.grid[piece.getPosition().Row][piece.getPosition().Col] = piece
}

func (b *board) InsertPieceList(pieces []movable) {
	for _, v := range pieces {
		b.insertPiece(v)
	}
}

func (b *board) MovePieceSim(from, to position) {
	piece, _ := b.getPiece(from)

	b.grid[from.Row][from.Col] = nil
	b.grid[to.Row][to.Col] = piece

}

// Moves piece on board. Does not update piece, only relocates in grid
func (b *board) movePiece(piece movable, pos position) movable {
	capture, _ := b.getPiece(pos)
	b.grid[piece.getPosition().Row][piece.getPosition().Col] = nil
	b.grid[pos.Row][pos.Col] = piece
	return capture
}

func (b *board) isKingInCheck(pos position, color bool) bool {
	return b.isSquareAttacked(pos, color)
}

func (b *board) isSquareAttacked(pos position, color bool) bool {
	return b.attackedByColor(!color)[pos]
}

// Calculates which squares are attacked by color
func (b *board) attackedByColor(white bool) map[position]bool {
	threats := make(map[position]bool)
	for i := range 8 {
		for j := range 8 {
			p := (*b.grid)[i][j]
			if p != nil && p.isWhite() == white {
				for sq := range p.visibleSquares(b) {
					threats[sq] = true
				}
			}
		}
	}
	return threats
}

// Find king by color
func (b *board) findKingPos(white bool) (position, bool) {
	for i := range 8 {
		for j := range 8 {
			p := (*b.grid)[i][j]
			if p != nil && p.isWhite() == white {
				if _, ok := p.(*king); ok {
					return p.getPosition(), true
				}
			}
		}
	}
	return position{}, false
}

func (b *board) IsOccupied(pos position) bool {
	return b.grid[pos.Row][pos.Col] != nil
}

func (b *board) clone() *board {
	newGrid := &[8][8]movable{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			piece := (*b.grid)[i][j]
			if piece != nil {
				newGrid[i][j] = piece.clone()
			}
		}
	}
	return &board{grid: newGrid}
}

func (b *board) clearSquare(pos position) {
	b.grid[pos.Row][pos.Col] = nil
}

func (b *board) isRowPathClear(from, to position) bool {
	if from.Row != to.Row {
		return false
	}

	step := 1
	if from.Col > to.Col {
		step = -1
	}

	for col := from.Col + step; col != to.Col; col += step {
		if b.IsOccupied(position{Row: from.Row, Col: col}) {
			return false
		}
	}

	return true
}

func (b *board) String() string {
	output := ""

	for row := 7; row >= 0; row-- {
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

	output += "  A  B  C  D  E  F  G  H\n"

	return output
}
