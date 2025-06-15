package main

type Movable interface {
	PossibleMoves(b *Board) []Position
	GetPosition() Position
	SetPosition(pos Position)
	IsWhite() bool
	String() string
}

type BasePiece struct {
	White    bool
	Value    int
	Pos      Position
	Captured bool
}

func NewBasePiece(white bool, value int, pos Position) *BasePiece {
	return &BasePiece{
		White:    white,
		Value:    value,
		Pos:      pos,
		Captured: false,
	}
}

// Updates piece position
func (bp *BasePiece) SetPosition(pos Position) {
	bp.Pos = pos
}
