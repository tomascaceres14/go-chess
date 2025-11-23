package gochess

import (
	"fmt"
)

type Board struct {
	grid            *[8][8]Movable
	enPassantTarget *Position
}

// Get piece at given pos. Returns piece and wether square is occupied or not.
// (should refactor, second bool not necessary as piece can be nil, not necessarily zero value)
func (b *Board) getPiece(pos Position) (Movable, bool) {
	piece := b.grid[pos.row][pos.col]
	return piece, b.IsOccupied(pos)
}

// Inserts piece in piece.Pos
func (b *Board) insertPiece(piece Movable) {
	b.grid[piece.getPosition().row][piece.getPosition().col] = piece
}

func (b *Board) InsertPieceList(pieces []Movable) {
	for _, v := range pieces {
		b.insertPiece(v)
	}
}

func (b *Board) MovePieceSim(from, to Position) {
	piece, _ := b.getPiece(from)
	b.grid[from.row][from.col] = nil
	b.grid[to.row][to.col] = piece
}

// Moves piece on board. Does not update piece, only relocates in grid
func (b *Board) movePiece(piece Movable, to Position) Movable {
	capture, _ := b.getPiece(to)
	b.grid[piece.getPosition().row][piece.getPosition().col] = nil
	b.grid[to.row][to.col] = piece
	return capture
}

func (b *Board) isKingInCheck(pos Position, color bool) bool {
	return b.isSquareAttacked(pos, color)
}

func (b *Board) isSquareAttacked(pos Position, color bool) bool {
	return b.attackedByColor(!color)[pos]
}

// Calculates which squares are attacked by color
func (b *Board) attackedByColor(white bool) map[Position]bool {
	threats := make(map[Position]bool)
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
func (b *Board) findKingPos(white bool) Position {
	for i := range 8 {
		for j := range 8 {
			p := (*b.grid)[i][j]
			if p != nil && p.isWhite() == white {
				if _, ok := p.(*king); ok {
					return Position{row: i, col: j}
				}
			}
		}
	}
	return Position{}
}

func (b *Board) IsOccupied(pos Position) bool {
	return b.grid[pos.row][pos.col] != nil
}

func (b *Board) clone() Board {
	newGrid := &[8][8]Movable{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			piece := (*b.grid)[i][j]
			if piece != nil {
				newGrid[i][j] = piece.clone()
			}
		}
	}
	return Board{grid: newGrid}
}

func (b *Board) clearSquare(pos Position) {
	b.grid[pos.row][pos.col] = nil
}

func (b *Board) isRowPathClear(from, to Position) bool {
	if from.row != to.row {
		return false
	}

	step := 1
	if from.col > to.col {
		step = -1
	}

	for col := from.col + step; col != to.col; col += step {
		if b.IsOccupied(Position{row: from.row, col: col}) {
			return false
		}
	}

	return true
}

func (b *Board) GetGrid() [8][8]Movable {
	return *b.grid
}

func (b *Board) String() string {
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
