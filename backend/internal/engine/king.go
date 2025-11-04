package engine

type king struct {
	*basePiece
	longCastlingOpt, shortCastlingOpt bool
}

func newKing(pos position, p *player) *king {
	white := p.isWhite
	directions := []direction{
		// left
		{-1, 0},
		// right
		{1, 0},
		// up
		{0, 1},
		// down
		{0, -1},
		// top right
		{1, 1},
		// top left
		{-1, 1},
		// bottom right
		{1, -1},
		// bottom left
		{-1, -1},
	}

	king := &king{
		basePiece:        newBasePiece(white, 0, pos, directions),
		longCastlingOpt:  true,
		shortCastlingOpt: true,
	}

	p.pieces = append(p.pieces, king)

	return king
}

func (k *king) visibleSquares(b *board) map[position]bool {
	positions := map[position]bool{}

	for _, v := range k.directions {
		pos := position{Row: k.pos.Row + v.dx, Col: k.pos.Col + v.dy}

		if !pos.inBounds() {
			continue
		}

		pieceAt, occupied := b.getPiece(pos)

		if !occupied || pieceAt.isWhite() != k.white {
			positions[pos] = true
			continue
		}
	}

	return positions
}

func (k *king) legalMoves(b *board) map[position]bool {
	legalMoves := k.visibleSquares(b)

	if k.moved {
		return legalMoves
	}

	// Initial rook pos hardocded. Should be obtained at game setup
	shortCastlePos := pos("a8")
	longCastlePos := pos("a1")

	shortRook, shortOcc := b.getPiece(shortCastlePos)
	longRook, longOcc := b.getPiece(longCastlePos)

	canShortCastle := shortOcc && !shortRook.hasMoved() && b.IsRowPathClear(k.pos, shortCastlePos)
	canLongCastle := longOcc && !longRook.hasMoved() && b.IsRowPathClear(k.pos, longCastlePos)

	// King's castling position is hardcoded. Should make calculation based on initial pos and distance to rook for Chess960
	legalMoves[pos("a7")] = canShortCastle
	legalMoves[pos("a3")] = canLongCastle

	return legalMoves
}

func (k *king) getPosition() position {
	return k.pos
}

func (k *king) setPosition(pos position) {
	k.moved = true
	k.pos = pos
}

func (k *king) isWhite() bool {
	return k.white
}

func (k *king) String() string {
	piece := "K"

	if !k.white {
		piece = "k"
	}

	return piece
}

func (b *king) getAlgebraicString() string {
	return "K"
}

func (k *king) clone() movable {
	return &king{basePiece: k.basePiece.cloneBase()}
}

func (k *king) getType() pieceType {
	return kingType
}
