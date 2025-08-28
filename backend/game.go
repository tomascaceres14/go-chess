package main

import (
	"errors"
	"fmt"
)

type Game struct {
	board          *Board
	PWhite, PBlack *Player
	Captures       []Movable
	WhiteTurn      bool
	MoveHistory    []Move
}

// Generates a new board with classic chess configuration
func NewGame(whites, blacks *Player) *Game {

	println("generating new board...")

	board := [8][8]Movable{}

	game := &Game{
		board:       &Board{grid: &board},
		PWhite:      whites,
		PBlack:      blacks,
		Captures:    []Movable{},
		WhiteTurn:   true,
		MoveHistory: []Move{},
	}

	// Pawns
	for i := range 8 {
		game.board.InsertPiece(NewPawn(Pos(GetCol(i)+"7"), blacks)) // black
		game.board.InsertPiece(NewPawn(Pos(GetCol(i)+"2"), whites)) // white
	}

	// Rooks
	game.board.InsertPiece(NewRook(Pos("a8"), blacks)) // black
	game.board.InsertPiece(NewRook(Pos("h8"), blacks)) // black
	game.board.InsertPiece(NewRook(Pos("a1"), whites)) // white
	game.board.InsertPiece(NewRook(Pos("h1"), whites)) // white

	// Knights
	game.board.InsertPiece(NewKnight(Pos("b8"), blacks)) // black
	game.board.InsertPiece(NewKnight(Pos("g8"), blacks)) // black
	game.board.InsertPiece(NewKnight(Pos("b1"), whites)) // white
	game.board.InsertPiece(NewKnight(Pos("g1"), whites)) // white

	// Bishops
	game.board.InsertPiece(NewBishop(Pos("c8"), blacks)) // black
	game.board.InsertPiece(NewBishop(Pos("f8"), blacks)) // black
	game.board.InsertPiece(NewBishop(Pos("c1"), whites)) // white
	game.board.InsertPiece(NewBishop(Pos("f1"), whites)) // white

	// Queens
	game.board.InsertPiece(NewQueen(Pos("d8"), blacks)) // black
	game.board.InsertPiece(NewQueen(Pos("d1"), whites)) // white

	// Kings
	bKing := NewKing(Pos("e8"), blacks)
	blacks.King = bKing
	game.board.InsertPiece(bKing)

	wKing := NewKing(Pos("e1"), whites)
	whites.King = wKing
	game.board.InsertPiece(wKing)

	return game
}

// Obtains piece at given position if player is owner of piece
func (g *Game) GetPiece(pos Position, player *Player) (Movable, error) {

	piece, ok := g.board.GetPiece(pos)
	if !ok {
		return nil, fmt.Errorf("Position %v is empty.", pos)
	}

	// validar que pieza pertenezca a player
	if piece.IsWhite() != player.White {
		return nil, fmt.Errorf("Not your piece, %s.", player.Name)
	}

	return piece, nil
}

// Moves piece in position `from` to position `to` if player is owner of piece
func (g *Game) MovePiece(from, to Position, player *Player) error {

	if !g.WhiteTurn == player.White {
		return fmt.Errorf("Not your turn, %s.", player.Name)
	}

	// Check if positions are in bounds
	if !from.InBounds() || !to.InBounds() {
		return errors.New("Position out of bounds.")
	}

	// Obtains opponent
	opponent := g.GetPlayerOpponent(player)

	// Obtains piece to move
	piece, err := g.GetPiece(from, player)
	if err != nil {
		return err
	}

	// Check if piece can move to desired position or if is trying to move in-place
	legalMoves := piece.LegalMoves(g.board)
	if !legalMoves[to] || piece.GetPosition() == to {
		return fmt.Errorf("%s can't move from %s to %s.", piece.String(), from, to)
	}

	if !IsMoveSafeToKing(piece, to, g) {
		return fmt.Errorf("%s to %s leaves king checked.", piece, to)
	}

	capture := g.board.MovePiece(piece, to)

	if capture != nil {
		opponent.Pieces = DeletePiece(opponent.Pieces, capture)
		player.Points += capture.GetValue()
	}

	attackedSquares := player.AttackedSquares(g.board)

	// Update threats map of opponent and flag as checked or not
	opponent.Threats = attackedSquares
	opponent.Checked = attackedSquares[opponent.King.Pos]

	// Check winning / draw conditions
	if !opponent.HasLegalMoves(g) {
		if opponent.Checked {
			fmt.Println("CHECKMATE!!!", player.Name, "WINS")
		} else {
			fmt.Println("Stalemate pal :(")
		}
	} else if len(player.Pieces) == 1 && len(opponent.Pieces) == 1 {
		fmt.Println("Stalemate pal :(")
	}

	move := Move{
		PieceCopy: piece.Clone(),
		From:      from,
		To:        to,
		Capture:   capture,
	}

	g.MoveHistory = append(g.MoveHistory, move)

	g.WhiteTurn = !g.WhiteTurn

	fmt.Println(g.MoveHistory)
	return nil
}

// Returns pointer to player based on color
func (g *Game) GetPlayer(white bool) *Player {
	player := g.PWhite

	if !white {
		player = g.PBlack
	}

	return player
}

// Returns pointer to player based on color
func (g *Game) GetPlayerCopy(white bool) Player {
	player := g.PWhite

	if !white {
		player = g.PBlack
	}

	return *player
}

// Returns pointer to player based on player
func (g *Game) GetPlayerOpponent(p *Player) *Player {
	opponent := g.PBlack

	if !p.White {
		opponent = g.PWhite
	}

	return opponent
}

// Returns pointer to player based on player
func (g *Game) GetPlayerOpponentCopy(white bool) Player {
	opponent := g.PBlack

	if !white {
		opponent = g.PWhite
	}

	return *opponent
}
