package engine

import (
	"errors"
	"fmt"
	"time"
)

type gameStatus int

const (
	aborted gameStatus = iota
	playing
	whiteWins
	blackWins
	draw
)

type game struct {
	id             string
	gameBoard      *board
	pWhite, pBlack *player
	WhiteTurn      bool
	moveHistory    []move
	status         gameStatus
}

// Mapping initial available castle square for king (key) and rook from/to movement option (value)
var castlingPositions = map[position]struct {
	rookFrom      position
	rookTo        position
	isShortCastle bool
}{
	// whites
	pos("g1"): {rookFrom: pos("h1"), rookTo: pos("f1"), isShortCastle: true},
	pos("c1"): {rookFrom: pos("a1"), rookTo: pos("d1"), isShortCastle: false},

	// blacks
	pos("g8"): {rookFrom: pos("h8"), rookTo: pos("f8"), isShortCastle: true},
	pos("c8"): {rookFrom: pos("a8"), rookTo: pos("d8"), isShortCastle: false},
}

// Generates a new board with classic chess configuration
func NewGame(whiteName, blackName string) *game {

	fmt.Println("generating new board...")

	gameBoard := [8][8]movable{}

	pWhite := newPlayer(whiteName, true)
	pBlack := newPlayer(blackName, false)

	// Generate ID for match
	now := time.Now()
	timestamp := now.Format("20060201150405")
	id := whiteName + "_" + blackName + "_" + timestamp

	game := &game{
		id:          id,
		gameBoard:   &board{grid: &gameBoard},
		pWhite:      pWhite,
		pBlack:      pBlack,
		WhiteTurn:   true,
		moveHistory: []move{},
		status:      playing,
	}

	// Pawns
	for i := range 8 {
		game.gameBoard.insertPiece(newPawn(pos(getColLetter(i)+"7"), pBlack)) // black
		game.gameBoard.insertPiece(newPawn(pos(getColLetter(i)+"2"), pWhite)) // white
	}

	// Rooks
	game.gameBoard.insertPiece(newRook(pos("a8"), pBlack)) // black
	game.gameBoard.insertPiece(newRook(pos("h8"), pBlack)) // black
	game.gameBoard.insertPiece(newRook(pos("a1"), pWhite)) // white
	game.gameBoard.insertPiece(newRook(pos("h1"), pWhite)) // white

	// Knights
	game.gameBoard.insertPiece(newKnight(pos("b8"), pBlack)) // black
	game.gameBoard.insertPiece(newKnight(pos("g8"), pBlack)) // black
	game.gameBoard.insertPiece(newKnight(pos("b1"), pWhite)) // white
	game.gameBoard.insertPiece(newKnight(pos("g1"), pWhite)) // white

	// Bishops
	game.gameBoard.insertPiece(newBishop(pos("c8"), pBlack)) // black
	game.gameBoard.insertPiece(newBishop(pos("f8"), pBlack)) // black
	game.gameBoard.insertPiece(newBishop(pos("c1"), pWhite)) // white
	game.gameBoard.insertPiece(newBishop(pos("f1"), pWhite)) // white

	// Queens
	game.gameBoard.insertPiece(newQueen(pos("d8"), pBlack)) // black
	game.gameBoard.insertPiece(newQueen(pos("d1"), pWhite)) // white

	// Kings
	bKing := newKing(pos("e8"), pBlack)
	pBlack.king = bKing
	game.gameBoard.insertPiece(bKing)

	wKing := newKing(pos("e1"), pWhite)
	pWhite.king = wKing
	game.gameBoard.insertPiece(wKing)

	fmt.Println(game.gameBoard)

	return game
}

// Obtains piece at given position if player is owner of piece
func (g *game) getPlayerPiece(pos position, pColor bool) (movable, error) {

	piece, ok := g.gameBoard.getPiece(pos)
	if !ok {
		return nil, fmt.Errorf("Position %v is empty.", pos)
	}

	// Validate piece belonging to player
	if piece.isWhite() != pColor {
		return nil, fmt.Errorf("Not your piece, %s.", colorToString(pColor))
	}

	return piece, nil
}

// Moves piece in position `from` to position `to` if player is owner of piece
func (game *game) movePiece(from, to position, pColor bool) error {

	// Check turn to play
	if game.WhiteTurn != pColor {
		return fmt.Errorf("Not your turn, %s.", colorToString(pColor))
	}

	// Check if positions are in bounds
	if !from.inBounds() || !to.inBounds() {
		return errors.New("Position out of bounds.")
	}

	// Obtain player and opponent
	player := game.GetPlayer(pColor)
	opponent := game.GetPlayerOpponent(player.isWhite)

	// Obtains piece to move
	piece, err := game.getPlayerPiece(from, player.isWhite)
	if err != nil {
		return err
	}

	// Check if piece can move to desired position or if is trying to move in-place
	legalMoves := piece.legalMoves(game.gameBoard)
	if !legalMoves[to] || piece.getPosition() == to {
		return fmt.Errorf("%s can't move from %s to %s.", piece.String(), from, to)
	}

	// Check if the move leaves king vulnerable
	if !isMoveSafeToKing(piece, to, game) {
		return fmt.Errorf("%s to %s leaves king checked.", piece, to)
	}

	// make move and return captured piece, if any
	capture := game.gameBoard.movePiece(piece, to)

	// Special moves: castling, promoting, etc.
	switch piece.getType() {
	case kingType:
		if rookMove, ok := castlingPositions[to]; ok {
			rook, _ := game.getPlayerPiece(rookMove.rookFrom, player.isWhite)
			game.gameBoard.movePiece(rook, rookMove.rookTo)

			king, _ := piece.(*king)
			if rookMove.isShortCastle {
				king.shortCastlingOpt = false
			} else {
				king.longCastlingOpt = false
			}
		}

	case pawnType:

		// Promotion
		if to.Row == 0 || to.Row == 7 {
			queen := newQueen(to, player)
			game.gameBoard.insertPiece(queen)
			queen.setMoved(true)
			break
		}

		pawn, _ := castPawn(piece)

		// Jumping
		if diff := from.Row - to.Row; diff == 2 || diff == -2 {
			pawn.jumped = true
			break
		}

		// En passant
		if to.Col != from.Col {
			capturedPos := position{Row: from.Row, Col: to.Col}
			capture, _ = game.gameBoard.getPiece(capturedPos)
			game.gameBoard.clearSquare(capturedPos)
			break
		}

	}

	// if piece was captured
	if capture != nil {
		opponent.pieces = deletePiece(opponent.pieces, capture)
		player.points += capture.getValue()
	}

	attackedSquares := player.attackedSquares(game.gameBoard)

	// Update threats map of opponent and flag as checked or not
	opponent.isChecked = attackedSquares[opponent.king.pos]

	// Check winning / draw conditions
	if !opponent.hasLegalMoves(game) {
		if opponent.isChecked {
			fmt.Println("CHECKMATE!!!", player.name, "WINS")
		} else {
			fmt.Println("Stalemate pal :(")
		}
	} else if len(player.pieces) == 1 && len(opponent.pieces) == 1 {
		fmt.Println("Stalemate pal :(")
	}

	// Update moves history
	move := newMove(piece.clone(), capture, from, to, opponent.isChecked)
	game.moveHistory = append(game.moveHistory, move)

	// Switch turns
	game.WhiteTurn = !game.WhiteTurn
	fmt.Println(game.gameBoard)

	return nil
}

// Returns pointer to player based on color
func (g *game) GetPlayer(pColor bool) *player {
	player := g.pWhite

	if !pColor {
		player = g.pBlack
	}

	return player
}

// Returns pointer to player based on color
func (g *game) GetPlayerCopy(white bool) player {
	player := g.pWhite

	if !white {
		player = g.pBlack
	}

	return *player
}

// Returns pointer to player based on player
func (g *game) GetPlayerOpponent(pColor bool) *player {
	opponent := g.pBlack

	if !pColor {
		opponent = g.pWhite
	}

	return opponent
}

// Returns pointer to player based on player
func (g *game) GetPlayerOpponentCopy(white bool) player {
	opponent := g.pBlack

	if !white {
		opponent = g.pWhite
	}

	return *opponent
}
