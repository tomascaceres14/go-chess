package gochess

type bishop struct {
	*basePiece
}

func newBishop(pos position, p *player) *bishop {
	white := p.isWhite
	directions := []direction{
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	bishop := &bishop{
		basePiece: newBasePiece(white, 3, pos, directions),
	}

	p.pieces = append(p.pieces, bishop)

	return bishop
}

func (bp *bishop) visibleSquares(b *board) map[position]bool {
	return bp.visibleSquaresDefault(b)
}

func (bp *bishop) legalMoves(b *board) map[position]bool {
	return bp.legalMovesDefault(b)
}

func (b *bishop) getPosition() position {
	return b.pos
}

func (b *bishop) isWhite() bool {
	return b.white
}

func (b *bishop) String() string {
	piece := "B"

	if !b.white {
		piece = "b"
	}

	return piece
}

func (b *bishop) getAlgebraicString() string {
	return "B"
}

func (b *bishop) clone() movable {
	return &bishop{basePiece: b.basePiece.cloneBase()}
}

func (b *bishop) getType() pieceType {
	return bishopType
}
