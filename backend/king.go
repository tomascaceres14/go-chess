package main

type King struct {
	*BasePiece
	hasMoved bool
}

func NewKing(white bool, pos Position) *King {
	return &King{
		BasePiece: NewBasePiece(white, 0, pos),
		hasMoved:  false,
	}
}

var kingDirs = []struct{ dx, dy int }{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func (k *King) PossibleMoves(b *Board) map[Position]bool {
	positions := map[Position]bool{}

	for _, v := range kingDirs {
		pos := Position{Row: k.Pos.Row + v.dx, Col: k.Pos.Col + v.dy}
		piece, occupied := b.GetPiece(pos)

		if !occupied || piece.IsWhite() != k.White {
			positions[pos] = true
			continue
		}
	}

	return positions
}

func (r *King) GetPosition() Position {
	return r.Pos
}

func (k *King) SetPosition(pos Position) {
	k.hasMoved = true
	k.Pos = pos
}

func (r *King) IsWhite() bool {
	return r.White
}

func (r *King) String() string {
	color := "w"

	if !r.White {
		color = "b"
	}

	return "K" + color
}
