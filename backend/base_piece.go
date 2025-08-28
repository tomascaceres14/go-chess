package main

type Movable interface {
	AttackedSquares(b *Board) map[Position]bool
	LegalMoves(b *Board) map[Position]bool
	GetPosition() Position
	SetPosition(pos Position)
	IsWhite() bool
	GetValue() int
	String() string
	HasMoved() bool
	Clone() Movable
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

func (bp *BasePiece) AttackedSquares(b *Board) map[Position]bool {
	return bp.AttackedSquaresDefault(b)
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

func (bp *BasePiece) HasMoved() bool {
	return bp.hasMoved
}

func (bp *BasePiece) CloneBase() *BasePiece {
	if bp == nil {
		return nil
	}
	cp := *bp
	// si Directions es slice, copiarlo también para que no compartan backing array
	if bp.Directions != nil {
		cp.Directions = make([]Direction, len(bp.Directions))
		copy(cp.Directions, bp.Directions)
	}
	return &cp
}
