package main

type Queen struct {
	*BasePiece
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

func NewQueen(pos Position, p *Player) *Queen {
	white := p.White

	queen := &Queen{
		BasePiece: NewBasePiece(white, 9, pos),
	}

	p.Pieces = append(p.Pieces, queen)

	return queen
}

func (r *Queen) PossibleMoves(b *Board) map[Position]bool {
	positions := map[Position]bool{}

	for _, v := range queenDirs {
		pos := Position{Row: r.Pos.Row + v.dx, Col: r.Pos.Col + v.dy}
		CastRay(pos, v.dx, v.dy, b, r.White, positions)
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

	return "Q" + color
}
