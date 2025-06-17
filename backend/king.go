package main

type King struct {
	*BasePiece
	hasMoved bool
}

func NewKing(pos Position, p *Player) *King {
	white := p.White
	directions := []Direction{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	king := &King{
		BasePiece: NewBasePiece(white, 0, pos, directions),
		hasMoved:  false,
	}

	p.Pieces = append(p.Pieces, king)

	return king
}

func (k *King) PossibleMoves(b *Board) map[Position]bool {
	positions := map[Position]bool{}

	for _, v := range k.Directions {
		pos := Position{Row: k.Pos.Row + v.dx, Col: k.Pos.Col + v.dy}

		if !pos.InBounds() {
			continue
		}

		pieceAt, occupied := b.GetPiece(pos)

		if !occupied || pieceAt.IsWhite() != k.White {
			positions[pos] = true
			continue
		}
	}

	return positions
}

func (r *King) GetPosition() Position {
	return r.Pos
}

func (k *King) SetPosition(pos Position) {
	k.hasMoved = true
	k.Pos = pos
}

func (r *King) IsWhite() bool {
	return r.White
}

func (r *King) String() string {
	color := "w"

	if !r.White {
		color = "b"
	}

	return "K" + color
}
