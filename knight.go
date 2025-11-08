package gochess

type knight struct {
	*basePiece
}

func newKnight(pos position, p *player) *knight {
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

func (n *knight) visibleSquares(b *board) map[position]bool {
	positions := map[position]bool{}

	for _, v := range n.directions {
		pos := position{Row: n.pos.Row + v.dx, Col: n.pos.Col + v.dy}

		if !pos.inBounds() {
			continue
		}

		positions[pos] = true
	}

	return positions
}

func (n *knight) legalMoves(b *board) map[position]bool {
	threats := n.visibleSquares(b)
	moves := map[position]bool{}

	for k := range threats {
		piece, occupied := b.getPiece(k)
		if !occupied || piece.isWhite() != n.white {
			moves[k] = true
			continue
		}
	}

	return moves
}

func (n *knight) getPosition() position {
	return n.pos
}

func (n *knight) isWhite() bool {
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

func (n *knight) clone() movable {
	return &knight{basePiece: n.basePiece.cloneBase()}
}

func (n *knight) getType() pieceType {
	return knightType
}
