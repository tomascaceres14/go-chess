package main

type Movable interface {
	AttackedSquares(b *Board) map[Position]bool
	LegalMoves(b *Board) map[Position]bool
	GetPosition() Position
	SetPosition(pos Position)
	IsWhite() bool
	String() string
}

type BasePiece struct {
	White      bool
	Value      int
	Pos        Position
	Directions []Direction
}

func NewBasePiece(white bool, value int, pos Position, directions []Direction) *BasePiece {
	return &BasePiece{
		White:      white,
		Value:      value,
		Pos:        pos,
		Directions: directions,
	}
}

func (bp *BasePiece) AttackedSquaresDefault(b *Board) map[Position]bool {
	positions := map[Position]bool{}

	for _, v := range bp.Directions {
		dir := Position{Row: bp.Pos.Row + v.dx, Col: bp.Pos.Col + v.dy}
		CastRay(dir, v.dx, v.dy, b, bp.White, positions)
	}

	return positions
}

func (bp *BasePiece) LegalMovesDefault(b *Board) map[Position]bool {
	threats := bp.AttackedSquaresDefault(b)
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

// Updates piece position
func (bp *BasePiece) SetPosition(pos Position) {
	bp.Pos = pos
}
