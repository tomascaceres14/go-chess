package gochess

type queen struct {
	*basePiece
}

func newQueen(pos position, p *player) *queen {
	white := p.isWhite
	directions := []direction{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	queen := &queen{
		basePiece: newBasePiece(white, 9, pos, directions),
	}

	p.pieces = append(p.pieces, queen)

	return queen
}

func (q *queen) visibleSquares(b *board) map[position]bool {
	return q.visibleSquaresDefault(b)
}

func (q *queen) legalMoves(b *board) map[position]bool {
	return q.legalMovesDefault(b)
}

func (q *queen) move(to position, game *game) movable {
	return q.moveDefault(to, game)
}

func (q *queen) getPosition() position {
	return q.pos
}

func (q *queen) isWhite() bool {
	return q.white
}

func (q *queen) String() string {
	piece := "Q"

	if !q.white {
		piece = "q"
	}

	return piece
}

func (b *queen) getAlgebraicString() string {
	return "Q"
}

func (q *queen) clone() movable {
	return &queen{basePiece: q.basePiece.cloneBase()}
}

func (q *queen) getType() pieceType {
	return queenType
}
