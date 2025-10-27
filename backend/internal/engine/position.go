package engine

import (
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	Row, Col int
}

type Direction struct {
	dx, dy int
}

func Pos(pos string) Position {
	nilPos := Position{-1, -1}

	if len(pos) != 2 {
		return nilPos
	}

	col := strings.ToLower(string(pos[0]))
	row, err := strconv.Atoi(string(pos[1]))
	if err != nil {
		return nilPos
	}

	if (row < 1 || row > 8) || (col[0] < 'a' || col[0] > 'h') {
		return nilPos
	}

	return Position{
		Row: row - 1,
		Col: int(col[0] - 'a'),
	}
}

func (p Position) Equals(other Position) bool {
	return p.Row == other.Row && p.Col == other.Col
}

func (p Position) InBounds() bool {
	return (0 <= p.Col && p.Col <= 7) && (0 <= p.Row && p.Row <= 7)
}

func (p Position) GetCol() string {
	return string(cols[p.Col])
}

func (p Position) String() string {
	return fmt.Sprintf("%v%v", string(cols[p.Col]), p.Row+1)
}

func ContainsPosition(list []Position, pos Position) bool {
	for _, p := range list {
		if p.Equals(pos) {
			return true
		}
	}

	return false
}
