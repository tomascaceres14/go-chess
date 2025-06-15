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
		BasePiece: &BasePiece{
			White: white,
			Value: 1,
			Pos:   pos,
		},
		direction: dir,
		hasMoved:  false,
	}
}

func (p *Pawn) PossibleMoves(b *Board) []Position {

	positions := []Position{}

	pos := Position{Row: p.Pos.Row + 1*p.direction, Col: p.Pos.Col}

	if b.GetPiece(pos) != nil {
		return positions
	}

	positions = append(positions, pos)

	if !p.hasMoved {
		pos.Row += 1 * p.direction
		positions = append(positions, pos)
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
