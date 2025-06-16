package main

type Rook struct {
	*BasePiece
	hasMoved bool
}

var rookDirs = []struct{ dx, dy int }{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

func NewRook(pos Position, p *Player) *Rook {
	white := p.White

	rook := &Rook{
		BasePiece: NewBasePiece(white, 5, pos),
		hasMoved:  false,
	}

	p.Pieces = append(p.Pieces, rook)
	return rook
}

func (r *Rook) PossibleMoves(b *Board) map[Position]bool {
	positions := map[Position]bool{}

	for _, v := range rookDirs {
		dir := Position{Row: r.Pos.Row + v.dx, Col: r.Pos.Col + v.dy}
		CastRay(dir, v.dx, v.dy, b, r.White, positions)
	}

	return positions
}

func (r *Rook) GetPosition() Position {
	return r.Pos
}

func (r *Rook) SetPosition(pos Position) {
	r.hasMoved = true
	r.Pos = pos
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
