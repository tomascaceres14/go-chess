package main

type Player struct {
	Name    string
	White   bool
	Points  int
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
