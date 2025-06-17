package main

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

func (p *Pawn) PossibleThreats(b *Board) map[Position]bool {

	positions := map[Position]bool{}

	front := Position{Row: p.Pos.Row + 1*p.direction, Col: p.Pos.Col}
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

	if !p.hasMoved {
		front.Row += 1 * p.direction
		positions[front] = true
	}

	return positions
}

func (p *Pawn) PossibleMoves(b *Board) map[Position]bool {

	positions := p.PossibleThreats(b)
	legalMoves := map[Position]bool{}

	front := Position{Row: p.Pos.Row + 1*p.direction, Col: p.Pos.Col}

	for k := range positions {

		if k == front {
			legalMoves[k] = true
			continue
		}

		piece, occ := b.GetPiece(k)

		if occ && piece.IsWhite() != p.White {
			legalMoves[k] = true
		}
	}

	return positions
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
