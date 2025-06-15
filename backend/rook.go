package main

type Rook struct {
	*BasePiece
	HasMoved bool
}

var rookDirs = []struct{ dx, dy int }{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

func NewRook(white bool, pos Position) *Rook {
	return &Rook{
		BasePiece: &BasePiece{
			White: white,
			Value: 5,
			Pos:   pos,
		},
		HasMoved: false,
	}
}

func (r *Rook) PossibleMoves(b *Board) []Position {
	positions := []Position{}

	for _, v := range rookDirs {
		dir := Position{Row: r.Pos.Row + v.dx, Col: r.Pos.Col + v.dy}
		CheckRayRecursive(dir, v.dx, v.dy, b, r.White, &positions)
	}

	return positions
}

func (r *Rook) GetPosition() Position {
	return r.Pos
}

func (r *Rook) IsWhite() bool {
	return r.White
}

func (r *Rook) String() string {
	color := "w"

	if !r.White {
		color = "b"
	}

	return "R" + color
}
