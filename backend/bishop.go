package main

type Bishop struct {
	*BasePiece
}

var bishopDirs = []struct{ dx, dy int }{
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func NewBishop(white bool, pos Position) *Bishop {
	return &Bishop{
		BasePiece: &BasePiece{
			White: white,
			Value: 3,
			Pos:   pos,
		},
	}
}

func (bp *Bishop) PossibleMoves(b *Board) []Position {
	positions := []Position{}

	for _, v := range bishopDirs {

		dir := Position{Row: bp.Pos.Row + v.dx, Col: bp.Pos.Col + v.dy}
		CheckRayRecursive(dir, v.dx, v.dy, b, bp.White, &positions)
	}

	return positions
}

func (b *Bishop) GetPosition() Position {
	return b.Pos
}

func (b *Bishop) IsWhite() bool {
	return b.White
}

func (b *Bishop) String() string {
	color := "w"

	if !b.White {
		color = "b"
	}

	return "B" + color
}
