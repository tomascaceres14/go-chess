package main

type Position struct {
	Row, Col int
}

func (p Position) Equals(other Position) bool {
	return p.Row == other.Row && p.Col == other.Col
}

func ContainsPosition(list []Position, pos Position) bool {
	for _, p := range list {
		if p.Equals(pos) {
			return true
		}
	}

	return false
}
