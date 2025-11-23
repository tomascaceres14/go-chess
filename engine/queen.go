package gochess

type queen struct {
	*basePiece
}

func newQueen(pos Position, p *player) *queen {
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

func (q *queen) visibleSquares(b *Board) map[Position]bool {
	return q.visibleSquaresDefault(b)
}

func (q *queen) legalMoves(b *Board) map[Position]bool {
	return q.legalMovesDefault(b)
}

func (q *queen) move(to Position, game *game) Movable {
	return moveDefault(q, to, game)
}

func (q *queen) getPosition() Position {
	return q.pos
}

func (q *queen) IsWhite() bool {
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

func (q *queen) clone() Movable {
	return &queen{basePiece: q.basePiece.cloneBase()}
}

func (q *queen) GetType() PieceType {
	return QueenType
}
