package main

type Movable interface {
	PossibleMoves(g *Board) []Position
	GetPosition() Position
	IsWhite() bool
	String() string
}

type BasePiece struct {
	White bool
	Value int
	Pos   Position
}
