package main

type Movable interface {
	PossibleMoves(b *Board) []Position
	GetPosition() Position
	SetPosition(pos Position)
	IsWhite() bool
	String() string
}

type BasePiece struct {
	White bool
	Value int
	Pos   Position
}

func (bp *BasePiece) SetPosition(pos Position) {
	bp.Pos = pos
}
