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

func (p *BasePiece) SetPosition(pos Position) {
	p.Pos = pos
}
