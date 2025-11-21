package gochess

type rook struct {
	*basePiece
}

func newRook(pos position, p *player) *rook {
	white := p.isWhite
	directions := []direction{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}

	rook := &rook{
		basePiece: newBasePiece(white, 5, pos, directions),
	}

	p.pieces = append(p.pieces, rook)
	return rook
}

func (r *rook) visibleSquares(b *board) map[position]bool {
	return r.visibleSquaresDefault(b)
}

func (r *rook) legalMoves(b *board) map[position]bool {
	return r.legalMovesDefault(b)
}

func (r *rook) move(to position, game *game) movable {
	from := r.pos
	board := game.gameBoard

	capture := board.movePiece(r, to)
	r.setPosition(to)

	if r.moved {
		return capture
	}

	r.setMoved(true)

	king := game.GetPlayer(r.white).getKing()
	if from.col == 0 {
		king.longCastlingOpt = false
	}

	if from.col == 7 {
		king.shortCastlingOpt = false
	}

	return capture
}

func (r *rook) getPosition() position {
	return r.pos
}

func (r *rook) setPosition(pos position) {
	r.pos = pos
}

func (r *rook) isWhite() bool {
	return r.white
}

func (r *rook) String() string {
	piece := "R"

	if !r.white {
		piece = "r"
	}

	return piece
}

func (b *rook) getAlgebraicString() string {
	return "R"
}

func (r *rook) clone() movable {
	return &rook{basePiece: r.basePiece.cloneBase()}
}

func (r *rook) getType() pieceType {
	return rookType
}
