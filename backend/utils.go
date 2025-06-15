package main

import (
	"fmt"
)

// func randomBool() bool {
// 	return rand.Intn(2) == 0
// }

const cols = "ABCDEFGH"

func PrintError(err error) {
	fmt.Printf("--- ERROR: %v\n", err)
}

func CastRay(pos Position, dx, dy int, b *Board, white bool, positions *[]Position) {

	if !pos.InBounds() {
		return
	}

	piece, ok := b.GetPiece(pos)
	if ok && piece.IsWhite() != white {
		*positions = append(*positions, pos)
		return
	}

	*positions = append(*positions, pos)
	next := Position{Row: pos.Row + dx, Col: pos.Col + dy}
	CastRay(next, dx, dy, b, white, positions)
}

func GetCol(col int) string {
	return string(cols[col])
}

func GetRow(row int) int {
	return row + 1
}
