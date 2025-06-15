package main

type Knight struct {
	*BasePiece
}

var knightMoves = []struct{ dx, dy int }{
	{2, 1},
	{1, 2},
	{2, -1},
	{1, -2},
	{-2, 1},
	{-1, 2},
	{-2, -1},
	{-1, -2},
}

func NewKnight(white bool, pos Position) *Knight {
	return &Knight{
		BasePiece: NewBasePiece(white, 3, pos),
	}
}

func (k *Knight) PossibleMoves(b *Board) map[Position]bool {
	positions := map[Position]bool{}

	for _, v := range knightMoves {
		pos := Position{Row: k.Pos.Row + v.dx, Col: k.Pos.Col + v.dy}

		if !pos.InBounds() {
			continue
		}

		pieceAt, occupied := b.GetPiece(pos)

		if !occupied || pieceAt.IsWhite() != k.White {
			positions[pos] = true
			continue
		}

	}

	return positions
}

func (r *Knight) GetPosition() Position {
	return r.Pos
}

func (r *Knight) IsWhite() bool {
	return r.White
}

func (r *Knight) String() string {
	color := "w"

	if !r.White {
		color = "b"
	}

	return "N" + color
}
