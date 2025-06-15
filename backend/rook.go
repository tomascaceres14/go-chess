package main

import "fmt"

type Rook struct {
	*BasePiece
	HasMoved bool
}

func NewRook(white bool, pos Position) *Rook {
	return &Rook{
		BasePiece: &BasePiece{
			White: white,
			Value: 5,
			Pos:   pos,
		},
		HasMoved: false,
	}
}

func (r *Rook) PossibleMoves(b *Board) []Position {
	positions := []Position{}

	top := Position{Row: r.Pos.Row, Col: r.Pos.Col + 1}
	bottom := Position{Row: r.Pos.Row, Col: r.Pos.Col - 1}
	left := Position{Row: r.Pos.Row - 1, Col: r.Pos.Col}
	right := Position{Row: r.Pos.Row + 1, Col: r.Pos.Col}

	CheckRayRecursive(top, 0, 1, b, r.White, &positions)
	CheckRayRecursive(bottom, 0, -1, b, r.White, &positions)
	CheckRayRecursive(left, -1, 0, b, r.White, &positions)
	CheckRayRecursive(right, 1, 0, b, r.White, &positions)

	return positions
}

func CheckRayRecursive(pos Position, dx, dy int, b *Board, white bool, positions *[]Position) {
	fmt.Printf("checking %v dir x=%v y=%v\n", pos, dx, dy)

	if !pos.InBounds() {
		fmt.Println("oob:", pos)
		return
	}

	if b.isOccupied(pos) {
		fmt.Println("occupied:", pos)
		if b.GetPiece(pos).IsWhite() != white {
			fmt.Println("can eat enemy:", pos)
			*positions = append(*positions, pos)
		}
		return
	}

	fmt.Println("valid pos:", pos)
	*positions = append(*positions, pos)
	next := Position{Row: pos.Row + dx, Col: pos.Col + dy}
	CheckRayRecursive(next, dx, dy, b, white, positions)
}

func (r *Rook) GetPosition() Position {
	return r.Pos
}

func (r *Rook) IsWhite() bool {
	return r.White
}

func (r *Rook) String() string {
	color := "w"

	if !r.White {
		color = "b"
	}

	return "R" + color
}
