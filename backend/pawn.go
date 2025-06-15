package main

type Pawn struct {
	*BasePiece
	direction int
	hasMoved  bool
}

func NewPawn(white bool, pos Position) *Pawn {

	dir := 1

	if !white {
		dir = -1
	}

	return &Pawn{
		BasePiece: NewBasePiece(white, 1, pos),
		direction: dir,
		hasMoved:  false,
	}
}

func (p *Pawn) PossibleMoves(b *Board) map[Position]bool {

	positions := map[Position]bool{}

	pos := Position{Row: p.Pos.Row + 1*p.direction, Col: p.Pos.Col}

	if b.IsOccupied(pos) {
		return positions
	}

	positions[pos] = true

	if !p.hasMoved {
		pos.Row += 1 * p.direction
		positions[pos] = true
	}

	return positions
}

func (p *Pawn) GetPosition() Position {
	return p.Pos
}

func (p *Pawn) SetPosition(pos Position) {
	p.hasMoved = true
	p.Pos = pos
}

func (p *Pawn) IsWhite() bool {
	return p.White
}

func (p *Pawn) String() string {
	color := "w"

	if !p.White {
		color = "b"
	}

	return "P" + color
}
