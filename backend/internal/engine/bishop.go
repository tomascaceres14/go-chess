package engine

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

func (bp *Bishop) VisibleSquares(b *Board) map[Position]bool {
	return bp.VisibleSquaresDefault(b)
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
	piece := "♗"

	if !b.White {
		piece = "♝"
	}

	return piece
}

func (b *Bishop) GetAlgebraicString() string {
	return "B"
}

func (b *Bishop) Clone() Movable {
	return &Bishop{BasePiece: b.BasePiece.CloneBase()}
}

func (b *Bishop) GetType() PieceType {
	return BishopType
}
