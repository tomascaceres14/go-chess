package main

type Player struct {
	Name    string
	White   bool
	Points  int
	Pieces  []Movable
	Threats *map[Position]bool
}

func NewPlayer(name string, isWhite bool) *Player {
	return &Player{
		Name:    name,
		White:   isWhite,
		Points:  0,
		Threats: &map[Position]bool{},
	}
}

// func (p *Player) UpdateThreats(b *Board) {
// 	threats := make(map[Position]bool)
// 	for i := range b.grid {
// 		for j := range b.grid[i] {
// 			piece, ok := b.GetPiece(Position{Row: i, Col: j})
// 			if ok && piece.IsWhite() != p.White {

// 			}
// 		}
// 	}
// }
