package main

type Movable interface {
	PossibleMoves(g *Board) []Position
	GetPosition() Position
	String() string
}

type BasePiece struct {
	white bool
	Value int
	Pos   Position
}

