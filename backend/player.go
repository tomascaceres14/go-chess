package main

type Player struct {
	Name   string
	White  bool
	Points int
}

func NewPlayer(name string, isWhite bool) *Player {
	return &Player{
		Name:   name,
		White:  isWhite,
		Points: 0,
	}
}
