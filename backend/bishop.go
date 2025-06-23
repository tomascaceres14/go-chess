package main

type Bishop struct {
	*BasePiece
}

func NewBishop(pos Position, p *Player) *Bishop {
	white := p.White
	directions := []Direction{
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	bishop := &Bishop{
		BasePiece: NewBasePiece(white, 3, pos, directions),
	}

	p.Pieces = append(p.Pieces, bishop)

	return bishop
}

func (bp *Bishop) AttackedSquares(b *Board) map[Position]bool {
	return bp.AttackedSquaresDefault(b)
}

func (bp *Bishop) LegalMoves(b *Board) map[Position]bool {
	return bp.LegalMovesDefault(b)
}

func (b *Bishop) GetPosition() Position {
	return b.Pos
}

func (b *Bishop) IsWhite() bool {
	return b.White
}

func (b *Bishop) String() string {
	color := "w"

	if !b.White {
		color = "b"
	}

	return "B" + color
}
