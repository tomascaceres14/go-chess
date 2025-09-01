package main

type Pawn struct {
	*BasePiece
	direction int
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
	}

	p.Pieces = append(p.Pieces, pawn)

	return pawn
}

func (p *Pawn) AttackedSquares(b *Board) map[Position]bool {

	positions := map[Position]bool{}
	front := Position{Row: p.Pos.Row + 1*p.direction, Col: p.Pos.Col}
	diag1 := Position{Row: p.Pos.Row + 1*p.direction, Col: p.Pos.Col + 1}
	diag2 := Position{Row: p.Pos.Row + 1*p.direction, Col: p.Pos.Col - 1}

	if diag1.InBounds() {
		positions[diag1] = true
	}

	if diag2.InBounds() {
		positions[diag2] = true
	}

	positions[front] = true

	if !p.hasMoved {
		front.Row += 1 * p.direction
		positions[front] = true
	}

	return positions
}

func (p *Pawn) LegalMoves(b *Board) map[Position]bool {

	positions := p.AttackedSquares(b)
	legalMoves := map[Position]bool{}

	for pos := range positions {

		piece, occupied := b.GetPiece(pos)

		// move front
		if pos.Col == p.Pos.Col {
			if !occupied {
				legalMoves[pos] = true
			}
			continue
		}

		// capture diagonal
		if occupied {
			if piece.IsWhite() != p.White {
				legalMoves[pos] = true
			}
			continue
		}

		// TODO: en passant
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
	piece := "♙"

	if !p.White {
		piece = "♟"
	}

	return piece
}

func (p *Pawn) Clone() Movable {
	return &Pawn{BasePiece: p.BasePiece.CloneBase()}
}

func (p *Pawn) GetType() PieceType {
	return PawnType
}
