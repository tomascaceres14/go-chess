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

func (k *King) AttackedSquares(b *Board) map[Position]bool {
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

func (k *King) LegalMoves(b *Board) map[Position]bool {
	return k.LegalMovesDefault(b)
}

func (k *King) GetPosition() Position {
	return k.Pos
}

func (k *King) SetPosition(pos Position) {
	k.hasMoved = true
	k.Pos = pos
}

func (k *King) IsWhite() bool {
	return k.White
}

func (k *King) String() string {
	color := "w"

	if !k.White {
		color = "b"
	}

	return "K" + color
}
