package main

type Queen struct {
	*BasePiece
}

func NewQueen(pos Position, p *Player) *Queen {
	white := p.White
	directions := []Direction{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	queen := &Queen{
		BasePiece: NewBasePiece(white, 9, pos, directions),
	}

	p.Pieces = append(p.Pieces, queen)

	return queen
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
