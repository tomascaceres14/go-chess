package main

type Queen struct {
	*BasePiece
}

func NewQueen(pos Position, p *Player) *Queen {
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

	queen := &Queen{
		BasePiece: NewBasePiece(white, 9, pos, directions),
	}

	p.Pieces = append(p.Pieces, queen)

	return queen
}

func (q *Queen) AttackedSquares(b *Board) map[Position]bool {
	return q.AttackedSquaresDefault(b)
}

func (q *Queen) LegalMoves(b *Board) map[Position]bool {
	return q.LegalMovesDefault(b)
}

func (q *Queen) GetPosition() Position {
	return q.Pos
}

func (q *Queen) IsWhite() bool {
	return q.White
}

func (q *Queen) String() string {
	piece := "♕"

	if !q.White {
		piece = "♛"
	}

	return piece
}

func (q *Queen) Clone() Movable {
	return &Queen{BasePiece: q.BasePiece.CloneBase()}
}

func (q *Queen) GetType() PieceType {
	return QueenType
}
