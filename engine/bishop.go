package gochess

type bishop struct {
	*basePiece
}

func newBishop(pos Position, p *player) *bishop {
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

func (b *bishop) visibleSquares(board *Board) map[Position]bool {
	return b.visibleSquaresDefault(board)
}

func (b *bishop) legalMoves(board *Board) map[Position]bool {
	return b.legalMovesDefault(board)
}

func (b *bishop) move(to Position, game *game) Movable {
	return moveDefault(b, to, game)
}

func (b *bishop) getPosition() Position {
	return b.pos
}

func (b *bishop) IsWhite() bool {
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

func (b *bishop) clone() Movable {
	return &bishop{basePiece: b.basePiece.cloneBase()}
}

func (b *bishop) GetType() PieceType {
	return BishopType
}
