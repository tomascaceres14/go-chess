package main

import (
	"errors"
	"fmt"
)

type Player struct {
	name   string
	white  bool
	points int
}

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

type Movable interface {
	PossibleMoves(g *Board) []Position
	GetPosition() Position
	String() string
}

type BasePiece struct {
	white bool
	Value int
	Pos   Position
}

type Pawn struct {
	*BasePiece
	direction int
	HasMoved  bool
}

func NewPawn(white bool, pos Position) *Pawn {

	dir := 1

	if !white {
		dir = -1
	}

	return &Pawn{
		BasePiece: &BasePiece{
			white: white,
			Value: 1,
			Pos:   pos,
		},
		direction: dir,
		HasMoved:  false,
	}
}

func (p *Pawn) PossibleMoves(b *Board) []Position {

	positions := []Position{}

	forward := p.Pos.Row + 1*p.direction

	if b.grid[forward][p.Pos.Col] != nil {
		return positions
	}

	positions = append(positions, Position{Row: forward, Col: p.Pos.Col})

	if !p.HasMoved {
		positions = append(positions, Position{Row: forward + 1*p.direction, Col: p.Pos.Col})
	}

	return positions
}

func (p *Pawn) Move(pos Position) error {
	return nil
}

func (p *Pawn) GetPosition() Position {
	return p.Pos
}

func (p *Pawn) String() string {
	color := "w"

	if !p.white {
		color = "b"
	}

	return "P" + color
}

type Board struct {
	grid [8][8]Movable
}

func (b *Board) MovePiece(piece Movable, pos Position) error {
	// asegurar que pos este dentro del tablero
	if pos.Col < 0 || pos.Col > 7 || pos.Row < 0 || pos.Row > 7 {
		return errors.New("Pos out of bounds")
	}

	// verificar si pieza puede moverse a pos
	if !ContainsPosition(piece.PossibleMoves(b), pos) {
		return errors.New("Piece cant move there")
	}

	currPos := piece.GetPosition()

	if pos == currPos {
		return errors.New("Cannot move to the same position")
	}

	b.grid[pos.Row][pos.Col] = piece
	b.grid[currPos.Row][currPos.Col] = nil

	return nil
}

func (b *Board) GetPiece(pos Position) Movable {
	return b.grid[pos.Row][pos.Col]
}

func (b *Board) String() string {
	output := ""

	for row := 7; row >= 0; row-- { // Mostrar del 8 al 1
		output += fmt.Sprintf("%d ", row+1)
		for col := 0; col < 8; col++ {
			piece := b.grid[row][col]
			if piece != nil {
				output += fmt.Sprintf("%-3s", piece.String())
			} else {
				output += fmt.Sprintf("%-3s", "--")
			}
		}
		output += "\n"
	}

	output += "   A  B  C  D  E  F  G  H\n"

	return output
}

type Game struct {
	board *Board
}

func NewGame() *Game {

	println("generating new board...")

	board := [8][8]Movable{}

	for i := range 8 {
		blackPos := Position{Row: 6, Col: i}
		whitePos := Position{Row: 1, Col: i}
		board[6][i] = NewPawn(false, blackPos) // black
		board[1][i] = NewPawn(true, whitePos)  // white
	}

	return &Game{board: &Board{grid: board}}
}

func main() {
	g := NewGame()

	fmt.Println(g.board)

	pawn := g.board.GetPiece(Position{Row: 1, Col: 2})
	fmt.Println(pawn)

	if err := g.board.MovePiece(pawn, Position{Row: 4, Col: 2}); err != nil {
		fmt.Printf("--- ERROR: %v\n", err)
	}

	fmt.Println(g.board)
}
