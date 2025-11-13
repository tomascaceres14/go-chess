package gochess

type king struct {
	*basePiece
	longCastlingOpt, shortCastlingOpt bool
	castleDir                         int
	moveFunc                          moveFunc
}

type castleMove struct {
	rookFrom, rookTo position
	castleDir        int
}

// Mapping initial available castle square for king (key) and rook from/to movement option (value)
var castlingPositions = map[position]castleMove{
	// whites
	pos("g1"): {rookFrom: pos("h1"), rookTo: pos("f1"), castleDir: 0},
	pos("c1"): {rookFrom: pos("a1"), rookTo: pos("d1"), castleDir: 1},

	// blacks
	pos("g8"): {rookFrom: pos("h8"), rookTo: pos("f8"), castleDir: 0},
	pos("c8"): {rookFrom: pos("a8"), rookTo: pos("d8"), castleDir: 1},
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
		castleDir:        -1,
	}

	king.moveFunc = king.moveWithCastling

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

	// Initial rook pos hardcoded. Should be obtained at game setup
	shortCastlePos := position{Row: k.pos.Row, Col: 7}
	longCastlePos := position{Row: k.pos.Row, Col: 0}

	shortRook, shortOcc := b.getPiece(shortCastlePos)
	longRook, longOcc := b.getPiece(longCastlePos)

	canShortCastle := shortOcc && !shortRook.hasMoved() && b.isRowPathClear(k.pos, shortCastlePos)
	canLongCastle := longOcc && !longRook.hasMoved() && b.isRowPathClear(k.pos, longCastlePos)

	// King's castling position is hardcoded. Should make calculation based on initial pos and distance to rook for Chess960
	legalMoves[position{Row: k.pos.Row, Col: 6}] = canShortCastle
	legalMoves[position{Row: k.pos.Row, Col: 2}] = canLongCastle

	return legalMoves
}

func (k *king) moveWithCastling(to position, game *game) movable {
	prevPos := k.pos
	board := game.gameBoard

	capture := board.movePiece(k, to)

	k.setPosition(to)
	k.setMoved(true)

	diff := prevPos.Col - to.Col
	k.moveFunc = k.moveDefault
	k.longCastlingOpt = false
	k.shortCastlingOpt = false

	// detect if king is castling
	if diff != 2 && diff != -2 {
		return capture
	}

	castleMove, ok := castlingPositions[to]
	if !ok {
		return nil
	}

	rook, _ := game.getPlayerPiece(castleMove.rookFrom, k.isWhite())
	board.movePiece(rook, castleMove.rookTo)
	k.castleDir = castleMove.castleDir

	return capture
}

func (k *king) move(to position, game *game) movable {
	return k.moveFunc(to, game)
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

func castKing(piece movable) (*king, bool) {
	king, ok := piece.(*king)
	return king, ok
}
