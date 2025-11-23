package gochess

type knight struct {
	*basePiece
}

func newKnight(pos Position, p *player) *knight {
	white := p.isWhite
	directions := []direction{
		{2, 1},
		{1, 2},
		{2, -1},
		{1, -2},
		{-2, 1},
		{-1, 2},
		{-2, -1},
		{-1, -2},
	}
	knight := &knight{
		basePiece: newBasePiece(white, 3, pos, directions),
	}

	p.pieces = append(p.pieces, knight)

	return knight
}

func (n *knight) visibleSquares(b *Board) map[Position]bool {
	positions := map[Position]bool{}

	for _, v := range n.directions {
		pos := Position{row: n.pos.row + v.dx, col: n.pos.col + v.dy}

		if !pos.inBounds() {
			continue
		}

		positions[pos] = true
	}

	return positions
}

func (n *knight) legalMoves(b *Board) map[Position]bool {
	threats := n.visibleSquares(b)
	moves := map[Position]bool{}

	for k := range threats {
		piece, occupied := b.getPiece(k)
		if !occupied || piece.IsWhite() != n.white {
			moves[k] = true
			continue
		}
	}

	return moves
}

func (n *knight) move(to Position, game *game) Movable {
	return moveDefault(n, to, game)
}

func (n *knight) getPosition() Position {
	return n.pos
}

func (n *knight) IsWhite() bool {
	return n.white
}

func (n *knight) String() string {
	piece := "N"

	if !n.white {
		piece = "n"
	}

	return piece
}

func (b *knight) getAlgebraicString() string {
	return "N"
}

func (n *knight) clone() Movable {
	return &knight{basePiece: n.basePiece.cloneBase()}
}

func (n *knight) GetType() PieceType {
	return KnightType
}
