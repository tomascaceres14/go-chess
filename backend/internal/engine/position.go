package engine

import (
	"fmt"
	"strconv"
	"strings"
)

type position struct {
	Row, Col int
}

type direction struct {
	dx, dy int
}

func pos(pos string) position {
	nilPos := position{-1, -1}

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

	return position{
		Row: row - 1,
		Col: int(col[0] - 'a'),
	}
}

func (p position) equals(other position) bool {
	return p.Row == other.Row && p.Col == other.Col
}

func (p position) inBounds() bool {
	return (0 <= p.Col && p.Col <= 7) && (0 <= p.Row && p.Row <= 7)
}

func (p position) getCol() string {
	return string(cols[p.Col])
}

func (p position) getRow() int {
	return p.Row + 1
}

func (p position) String() string {
	return fmt.Sprintf("%v%v", string(cols[p.Col]), p.Row+1)
}

func containsPosition(list []position, pos position) bool {
	for _, p := range list {
		if p.equals(pos) {
			return true
		}
	}

	return false
}
