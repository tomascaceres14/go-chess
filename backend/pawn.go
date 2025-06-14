package main

type Pawn struct {
	*BasePiece
	direction int
	HasMoved  bool
}

func NewPawn(white bool, pos Position) *Pawn {

	dir := 1

	if !white {
		dir = -1
	}

	return &Pawn{
		BasePiece: &BasePiece{
			white: white,
			Value: 1,
			Pos:   pos,
		},
		direction: dir,
		HasMoved:  false,
	}
}

func (p *Pawn) PossibleMoves(b *Board) []Position {

	positions := []Position{}

	forward := p.Pos.Row + 1*p.direction

	if b.grid[forward][p.Pos.Col] != nil {
		return positions
	}

	positions = append(positions, Position{Row: forward, Col: p.Pos.Col})

	if !p.HasMoved {
		positions = append(positions, Position{Row: forward + 1*p.direction, Col: p.Pos.Col})
	}

	return positions
}

func (p *Pawn) Move(pos Position) error {
	return nil
}

func (p *Pawn) GetPosition() Position {
	return p.Pos
}

func (p *Pawn) String() string {
	color := "w"

	if !p.white {
		color = "b"
	}

	return "P" + color
}
