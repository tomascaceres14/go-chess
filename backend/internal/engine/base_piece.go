package engine

type PieceType int

const (
	KingType PieceType = iota
	QueenType
	RookType
	BishopType
	KnightType
	PawnType
	BasePieceType
)

type Movable interface {
	VisibleSquares(b *Board) map[Position]bool
	LegalMoves(b *Board) map[Position]bool
	GetPosition() Position
	SetPosition(pos Position)
	IsWhite() bool
	GetValue() int
	String() string
	GetAlgebraicString() string
	HasMoved() bool
	SetMoved(moved bool)
	Clone() Movable
	GetType() PieceType
}

type BasePiece struct {
	White      bool
	Value      int
	Pos        Position
	Directions []Direction
	hasMoved   bool
}

func NewBasePiece(white bool, value int, pos Position, directions []Direction) *BasePiece {
	return &BasePiece{
		White:      white,
		Value:      value,
		Pos:        pos,
		Directions: directions,
		hasMoved:   false,
	}
}

func (bp *BasePiece) VisibleSquaresDefault(b *Board) map[Position]bool {
	positions := map[Position]bool{}

	for _, v := range bp.Directions {
		dir := Position{Row: bp.Pos.Row + v.dx, Col: bp.Pos.Col + v.dy}
		CastRay(dir, v.dx, v.dy, b, bp.White, positions)
	}

	return positions
}

func (bp *BasePiece) VisibleSquares(b *Board) map[Position]bool {
	return bp.VisibleSquaresDefault(b)
}

func (bp *BasePiece) LegalMovesDefault(b *Board) map[Position]bool {
	threats := bp.VisibleSquaresDefault(b)
	moves := map[Position]bool{}
	for k := range threats {
		piece, occupied := b.GetPiece(k)
		if !occupied || piece.IsWhite() != bp.White {
			moves[k] = true
			continue
		}
	}

	return moves
}

func (bp *BasePiece) LegalMoves(b *Board) map[Position]bool {
	return bp.LegalMovesDefault(b)
}

func (bp *BasePiece) GetPosition() Position {
	return bp.Pos
}

func (bp *BasePiece) SetPosition(pos Position) {
	bp.Pos = pos
	bp.hasMoved = true
}

func (bp *BasePiece) IsWhite() bool {
	return bp.White
}

func (bp *BasePiece) GetValue() int {
	return bp.Value
}

func (bp *BasePiece) String() string {
	color := "w"

	if !bp.White {
		color = "b"
	}

	return "BP" + color
}

func (b *BasePiece) GetAlgebraicString() string {
	return ""
}

func (bp *BasePiece) HasMoved() bool {
	return bp.hasMoved
}

func (bp *BasePiece) Clone() Movable {
	return bp.CloneBase()
}

func (bp *BasePiece) CloneBase() *BasePiece {
	if bp == nil {
		return nil
	}
	cp := *bp
	// Dont believe is necessary. Will check
	if bp.Directions != nil {
		cp.Directions = make([]Direction, len(bp.Directions))
		copy(cp.Directions, bp.Directions)
	}
	return &cp
}

func (bp *BasePiece) GetType() PieceType {
	return BasePieceType
}

func (bp *BasePiece) SetMoved(moved bool) {
	bp.hasMoved = moved
}
