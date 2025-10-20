package engine

type Rook struct {
	*BasePiece
}

func NewRook(pos Position, p *Player) *Rook {
	white := p.White
	directions := []Direction{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}

	rook := &Rook{
		BasePiece: NewBasePiece(white, 5, pos, directions),
	}

	p.Pieces = append(p.Pieces, rook)
	return rook
}

func (r *Rook) VisibleSquares(b *Board) map[Position]bool {
	return r.VisibleSquaresDefault(b)
}

func (r *Rook) LegalMoves(b *Board) map[Position]bool {
	return r.LegalMovesDefault(b)
}

func (r *Rook) GetPosition() Position {
	return r.Pos
}

func (r *Rook) SetPosition(pos Position) {
	r.hasMoved = true
	r.Pos = pos
}

func (r *Rook) IsWhite() bool {
	return r.White
}

func (r *Rook) String() string {
	piece := "♖"

	if !r.White {
		piece = "♜"
	}

	return piece
}

func (b *Rook) GetAlgebraicString() string {
	return "R"
}

func (r *Rook) Clone() Movable {
	return &Rook{BasePiece: r.BasePiece.CloneBase()}
}

func (r *Rook) GetType() PieceType {
	return RookType
}
