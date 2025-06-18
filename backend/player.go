package main

type Player struct {
	Name    string
	White   bool
	Points  int
	Pieces  []Movable
	Threats map[Position]bool
	Checked bool
}

func NewPlayer(name string, isWhite bool) *Player {
	return &Player{
		Name:    name,
		White:   isWhite,
		Points:  0,
		Threats: map[Position]bool{},
		Checked: false,
	}
}

func (p *Player) CalculateThreats(b *Board) map[Position]bool {
	threats := make(map[Position]bool)

	pieces := p.Pieces
	for _, v := range pieces {
		for k := range v.PossibleThreats(b) {
			threats[k] = true
		}
	}

	return threats
}
