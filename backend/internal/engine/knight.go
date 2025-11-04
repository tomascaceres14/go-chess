package engine

type Knight struct {
	*BasePiece
}

func NewKnight(pos Position, p *Player) *Knight {
	white := p.White
	directions := []Direction{
		{2, 1},
		{1, 2},
		{2, -1},
		{1, -2},
		{-2, 1},
		{-1, 2},
		{-2, -1},
		{-1, -2},
	}
	knight := &Knight{
		BasePiece: NewBasePiece(white, 3, pos, directions),
	}

	p.Pieces = append(p.Pieces, knight)

	return knight
}

func (n *Knight) VisibleSquares(b *Board) map[Position]bool {
	positions := map[Position]bool{}

	for _, v := range n.Directions {
		pos := Position{Row: n.Pos.Row + v.dx, Col: n.Pos.Col + v.dy}

		if !pos.InBounds() {
			continue
		}

		positions[pos] = true
	}

	return positions
}

func (n *Knight) LegalMoves(b *Board) map[Position]bool {
	threats := n.VisibleSquares(b)
	moves := map[Position]bool{}

	for k := range threats {
		piece, occupied := b.GetPiece(k)
		if !occupied || piece.IsWhite() != n.White {
			moves[k] = true
			continue
		}
	}

	return moves
}

func (n *Knight) GetPosition() Position {
	return n.Pos
}

func (n *Knight) IsWhite() bool {
	return n.White
}

func (n *Knight) String() string {
	piece := "N"

	if !n.White {
		piece = "n"
	}

	return piece
}

func (b *Knight) GetAlgebraicString() string {
	return "N"
}

func (n *Knight) Clone() Movable {
	return &Knight{BasePiece: n.BasePiece.CloneBase()}
}

func (n *Knight) GetType() PieceType {
	return KnightType
}
