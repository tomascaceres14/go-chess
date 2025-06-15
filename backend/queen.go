package main

type Queen struct {
	*BasePiece
	HasMoved bool
}

var queenDirs = []struct{ dx, dy int }{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func NewQueen(white bool, pos Position) *Queen {
	return &Queen{
		BasePiece: &BasePiece{
			White: white,
			Value: 5,
			Pos:   pos,
		},
		HasMoved: false,
	}
}

func (r *Queen) PossibleMoves(b *Board) []Position {
	positions := []Position{}

	for _, v := range queenDirs {
		dir := Position{Row: r.Pos.Row + v.dx, Col: r.Pos.Col + v.dy}
		CastRay(dir, v.dx, v.dy, b, r.White, &positions)
	}

	return positions
}

func (r *Queen) GetPosition() Position {
	return r.Pos
}

func (r *Queen) IsWhite() bool {
	return r.White
}

func (r *Queen) String() string {
	color := "w"

	if !r.White {
		color = "b"
	}

	return "R" + color
}
