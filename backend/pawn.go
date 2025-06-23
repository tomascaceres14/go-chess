package main

import "fmt"

type Pawn struct {
	*BasePiece
	direction int
	hasMoved  bool
}

func NewPawn(pos Position, p *Player) *Pawn {
	white := p.White

	dir := 1
	if !white {
		dir = -1
	}

	pawn := &Pawn{
		BasePiece: NewBasePiece(white, 1, pos, nil),
		direction: dir,
		hasMoved:  false,
	}

	p.Pieces = append(p.Pieces, pawn)

	return pawn
}

func (p *Pawn) AttackedSquares(b *Board) map[Position]bool {

	positions := map[Position]bool{}
	front := Position{Row: p.Pos.Row + 1*p.direction, Col: p.Pos.Col}
	front2 := Position{Row: p.Pos.Row + 1*p.direction, Col: p.Pos.Col}
	diag1 := Position{Row: p.Pos.Row + 1*p.direction, Col: p.Pos.Col + 1}
	diag2 := Position{Row: p.Pos.Row + 1*p.direction, Col: p.Pos.Col - 1}

	if b.IsOccupied(front) {
		return positions
	}

	if diag1.InBounds() {
		if piece, occupied := b.GetPiece(diag1); !occupied || piece.IsWhite() != p.White {
			positions[diag1] = true
		}
	}

	if diag2.InBounds() {
		if piece, occupied := b.GetPiece(diag2); !occupied || piece.IsWhite() != p.White {
			positions[diag2] = true
		}
	}

	positions[front] = true

	if !p.hasMoved && !b.IsOccupied(front2) {
		positions[front2] = true
	}

	return positions
}

func (p *Pawn) LegalMoves(b *Board) map[Position]bool {

	positions := p.AttackedSquares(b)
	fmt.Println("ATTACKING POSITIONS", positions)
	legalMoves := map[Position]bool{}

	for k := range positions {

		if k.Col != p.Pos.Col {
			continue
		}

		piece, occ := b.GetPiece(k)

		if !occ || piece.IsWhite() != p.White {
			legalMoves[k] = true
		}
	}

	return legalMoves
}

func (p *Pawn) GetPosition() Position {
	return p.Pos
}

func (p *Pawn) SetPosition(pos Position) {
	p.hasMoved = true
	p.Pos = pos
}

func (p *Pawn) IsWhite() bool {
	return p.White
}

func (p *Pawn) String() string {
	color := "w"

	if !p.White {
		color = "b"
	}

	return "P" + color
}
