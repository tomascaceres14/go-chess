package main

type Move struct {
	PieceCopy Movable
	From, To  Position
	Capture   Movable
}
