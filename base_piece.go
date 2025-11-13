package gochess

type pieceType int

const (
	kingType pieceType = iota
	queenType
	rookType
	bishopType
	knightType
	pawnType
	basePieceType
)

type movable interface {
	visibleSquares(b *board) map[position]bool
	legalMoves(b *board) map[position]bool
	getPosition() position
	setPosition(pos position)
	isWhite() bool
	getValue() int
	String() string
	getAlgebraicString() string
	move(to position, game *game) movable
	hasMoved() bool
	setMoved(moved bool)
	clone() movable
	getType() pieceType
}

type basePiece struct {
	white      bool
	value      int
	pos        position
	directions []direction
	moved      bool
}

func newBasePiece(white bool, value int, pos position, directions []direction) *basePiece {
	return &basePiece{
		white:      white,
		value:      value,
		pos:        pos,
		directions: directions,
		moved:      false,
	}
}

func (bp *basePiece) visibleSquaresDefault(b *board) map[position]bool {
	positions := map[position]bool{}

	for _, v := range bp.directions {
		dir := position{Row: bp.pos.Row + v.dx, Col: bp.pos.Col + v.dy}
		castRay(dir, v.dx, v.dy, b, bp.white, positions)
	}

	return positions
}

func (bp *basePiece) visibleSquares(b *board) map[position]bool {
	return bp.visibleSquaresDefault(b)
}

func (bp *basePiece) legalMovesDefault(b *board) map[position]bool {
	threats := bp.visibleSquaresDefault(b)
	moves := map[position]bool{}
	for k := range threats {
		piece, occupied := b.getPiece(k)
		if !occupied || piece.isWhite() != bp.white {
			moves[k] = true
			continue
		}
	}

	return moves
}

func (bp *basePiece) legalMoves(b *board) map[position]bool {
	return bp.legalMovesDefault(b)
}

func (bp *basePiece) moveDefault(to position, game *game) movable {
	board := game.gameBoard
	capture := board.movePiece(bp, to)
	bp.setPosition(to)
	bp.setMoved(true)
	return capture
}

func (bp *basePiece) move(to position, game *game) movable {
	return bp.moveDefault(to, game)
}

func (bp *basePiece) getPosition() position {
	return bp.pos
}

func (bp *basePiece) setPosition(pos position) {
	bp.pos = pos
	bp.moved = true
}

func (bp *basePiece) isWhite() bool {
	return bp.white
}

func (bp *basePiece) getValue() int {
	return bp.value
}

func (bp *basePiece) String() string {
	color := "w"

	if !bp.white {
		color = "b"
	}

	return "BP" + color
}

func (b *basePiece) getAlgebraicString() string {
	return ""
}

func (bp *basePiece) hasMoved() bool {
	return bp.moved
}

func (bp *basePiece) clone() movable {
	return bp.cloneBase()
}

func (bp *basePiece) cloneBase() *basePiece {
	if bp == nil {
		return nil
	}
	cp := *bp
	// Dont believe is necessary. Will check
	if bp.directions != nil {
		cp.directions = make([]direction, len(bp.directions))
		copy(cp.directions, bp.directions)
	}
	return &cp
}

func (bp *basePiece) getType() pieceType {
	return basePieceType
}

func (bp *basePiece) setMoved(moved bool) {
	bp.moved = moved
}
