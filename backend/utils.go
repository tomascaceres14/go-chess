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

func CheckRayRecursive(pos Position, dx, dy int, b *Board, white bool, positions *[]Position) {
	if !pos.InBounds() {
		return
	}

	if b.isOccupied(pos) {
		if b.GetPiece(pos).IsWhite() != white {
			*positions = append(*positions, pos)
		}
		return
	}

	*positions = append(*positions, pos)
	next := Position{Row: pos.Row + dx, Col: pos.Col + dy}
	CheckRayRecursive(next, dx, dy, b, white, positions)
}
