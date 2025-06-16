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

func NewBishop(pos Position, p *Player) *Bishop {
	white := p.White

	bishop := &Bishop{
		BasePiece: NewBasePiece(white, 3, pos),
	}

	p.Pieces = append(p.Pieces, bishop)

	return bishop
}

func (bp *Bishop) PossibleMoves(b *Board) map[Position]bool {
	positions := map[Position]bool{}

	for _, v := range bishopDirs {

		dir := Position{Row: bp.Pos.Row + v.dx, Col: bp.Pos.Col + v.dy}
		CastRay(dir, v.dx, v.dy, b, bp.White, positions)
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
