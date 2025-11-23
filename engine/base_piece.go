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

type Movable interface {
	visibleSquares(b *Board) map[Position]bool
	legalMoves(b *Board) map[Position]bool
	getPosition() Position
	setPosition(pos Position)
	isWhite() bool
	getValue() int
	String() string
	getAlgebraicString() string
	move(to Position, game *game) Movable
	hasMoved() bool
	setMoved(moved bool)
	clone() Movable
	getType() pieceType
}

type basePiece struct {
	white      bool
	value      int
	pos        Position
	directions []direction
	moved      bool
}

type moveFunc func(to Position, game *game) Movable

func newBasePiece(white bool, value int, pos Position, directions []direction) *basePiece {
	return &basePiece{
		white:      white,
		value:      value,
		pos:        pos,
		directions: directions,
		moved:      false,
	}
}

func (bp *basePiece) visibleSquaresDefault(b *Board) map[Position]bool {
	positions := map[Position]bool{}

	for _, v := range bp.directions {
		dir := Position{row: bp.pos.row + v.dx, col: bp.pos.col + v.dy}
		castRay(dir, v.dx, v.dy, b, bp.white, positions)
	}

	return positions
}

func (bp *basePiece) visibleSquares(b *Board) map[Position]bool {
	return bp.visibleSquaresDefault(b)
}

func (bp *basePiece) legalMovesDefault(b *Board) map[Position]bool {
	threats := bp.visibleSquaresDefault(b)
	moves := map[Position]bool{}
	for k := range threats {
		piece, occupied := b.getPiece(k)
		if !occupied || piece.isWhite() != bp.white {
			moves[k] = true
			continue
		}
	}

	return moves
}

func (bp *basePiece) legalMoves(b *Board) map[Position]bool {
	return bp.legalMovesDefault(b)
}

func moveDefault(piece Movable, to Position, game *game) Movable {
	board := game.gameBoard

	capture := board.movePiece(piece, to)
	piece.setPosition(to)
	piece.setMoved(true)

	board.enPassantTarget = nil
	return capture
}

func (bp *basePiece) move(to Position, game *game) Movable {
	return moveDefault(bp, to, game)
}

func (bp *basePiece) getPosition() Position {
	return bp.pos
}

func (bp *basePiece) setPosition(pos Position) {
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

func (bp *basePiece) clone() Movable {
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
