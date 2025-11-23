package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	gochess "github.com/tomascaceres14/go-chess/engine"
)

const (
	tileSize  = 130
	boardSize = 8
)

type Sprite struct {
	Img        *ebiten.Image
	posX, posY float64
}

type Piece struct {
	*Sprite
	pieceType gochess.PieceType
}

type Player struct {
	Pieces []*Piece
}

func (g Game) GetPiece(pieceType gochess.PieceType, color bool) *Piece {

	player := g.WPlayer

	if !color {
		player = g.BPlayer
	}

	for _, v := range player.Pieces {
		if v.pieceType == pieceType {
			return v
		}

	}

	return nil
}

type Game struct {
	WPlayer            *Player
	BPlayer            *Player
	PaddingX, PaddingY float32
	WSquare, BSquare   *ebiten.Image
	Board              gochess.Grid
}

// return sprite of img centered in indicated x and y positions
func CenteredPieceSprite(img *ebiten.Image, x, y float64) *Sprite {
	imgLength := img.Bounds().Dx()
	imgWidth := img.Bounds().Dy()
	return &Sprite{
		Img:  img,
		posX: x + float64((tileSize-imgLength)/2),
		posY: y + float64((tileSize-imgWidth)/2),
	}
}

func CenteredPieceSprite2(img *ebiten.Image) *Sprite {
	imgLength := img.Bounds().Dx()
	imgWidth := img.Bounds().Dy()
	return &Sprite{
		Img:  img,
		posX: float64((tileSize - imgLength) / 2),
		posY: float64((tileSize - imgWidth) / 2),
	}
}

func NewPiece(img *ebiten.Image, imgX, imgY float64) *Piece {
	return &Piece{
		Sprite: CenteredPieceSprite(img, imgX, imgY),
	}
}

func NewPiece2(img *ebiten.Image) *Piece {
	return &Piece{
		Sprite: CenteredPieceSprite2(img),
	}
}

// cuts img and extracts piece sub image
func PieceSubImg(img *ebiten.Image, x0, y0, x1, y1 int) *ebiten.Image {
	return img.SubImage(
		image.Rect(x0, y0, x1, y1),
	).(*ebiten.Image)
}

func NewPawn(img *ebiten.Image, x, y float64) *Piece {

	X0 := 278
	Y0 := 136
	X1 := 356
	Y1 := 256

	subImg := PieceSubImg(img, X0, Y0, X1, Y1)

	p := NewPiece2(subImg)
	p.pieceType = gochess.PawnType

	return p
}

func NewBishop(isWhite bool, img *ebiten.Image, x, y float64) *Piece {

	X0 := 127
	Y0 := 2
	X1 := 226
	Y1 := 117

	subImg := PieceSubImg(img, X0, Y0, X1, Y1)

	p := NewPiece(subImg, x, y)
	p.pieceType = gochess.BishopType

	return p

}

func NewKnight(isWhite bool, img *ebiten.Image, x, y float64) *Piece {

	X0 := 129
	Y0 := 140
	X1 := 218
	Y1 := 256

	subImg := PieceSubImg(img, X0, Y0, X1, Y1)

	p := NewPiece(subImg, x, y)
	p.pieceType = gochess.KnightType

	return p

}

func NewRook(isWhite bool, img *ebiten.Image, x, y float64) *Piece {

	x0 := 0
	y0 := 0
	x1 := 82
	y1 := 117

	subImg := PieceSubImg(img, x0, y0, x1, y1)
	p := NewPiece(subImg, x, y)
	p.pieceType = gochess.RookType

	return p
}

func NewQueen(isWhite bool, img *ebiten.Image, x, y float64) *Piece {
	X0 := 3
	Y0 := 136
	X1 := 79
	Y1 := 256

	subImg := PieceSubImg(img, X0, Y0, X1, Y1)
	p := NewPiece(subImg, x, y)
	p.pieceType = gochess.QueenType

	return p
}

func NewKing(isWhite bool, img *ebiten.Image, x, y float64) *Piece {
	X0 := 279
	Y0 := 0
	X1 := 356
	Y1 := 117

	subImg := PieceSubImg(img, X0, Y0, X1, Y1)

	p := NewPiece(subImg, x, y)
	p.pieceType = gochess.KingType

	return p
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{240, 240, 240, 255})
	opts := &ebiten.DrawImageOptions{}
	x := 0
	for i := len(g.Board) - 1; i >= 0; i-- {
		row := g.Board[i]

		for j, square := range row {

			posX := float64(g.PaddingX + float32(j*tileSize))
			posY := float64(g.PaddingY + float32(x*tileSize))

			opts.GeoM.Translate(posX, posY)

			if (i+j)%2 == 0 {
				//draw black
				screen.DrawImage(g.BSquare, opts)
			} else {
				//draw white
				screen.DrawImage(g.WSquare, opts)
			}

			if square != nil {

				piece := g.GetPiece(square.GetType(), square.IsWhite())
				screen.DrawImage(
					piece.Sprite.Img,
					opts,
				)
			}
			opts.GeoM.Reset()

		}
		x += 1
	}

	ebitenutil.DebugPrint(screen, "Hello, World!")
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func main() {
	ebiten.SetWindowTitle("Mates")
	ebiten.SetWindowSize(1920, 1080)

	g := &Game{}

	winX, winY := ebiten.WindowSize()
	g.PaddingX = float32(winX)/2 - (tileSize*boardSize)/2
	g.PaddingY = float32(winY)/2 - (tileSize*boardSize)/2

	whitePieces, _, err := ebitenutil.NewImageFromFile(`assets/images/whitepieces.png`)
	if err != nil {
		log.Fatal(err)
	}

	blackPieces, _, err := ebitenutil.NewImageFromFile(`assets/images/blackpieces.png`)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: pls make func to generate automatically positions based on algebraic notation.
	// This burns my eyes.
	wPlayer := &Player{
		Pieces: []*Piece{
			NewRook(true, whitePieces, float64(g.PaddingX), float64(g.PaddingY)+tileSize*7),
			NewKnight(true, whitePieces, float64(g.PaddingX)+tileSize, float64(g.PaddingY)+tileSize*7),
			NewBishop(true, whitePieces, float64(g.PaddingX)+tileSize*2, float64(g.PaddingY)+tileSize*7),
			NewQueen(true, whitePieces, float64(g.PaddingX)+tileSize*3, float64(g.PaddingY)+tileSize*7),
			NewKing(true, whitePieces, float64(g.PaddingX)+tileSize*4, float64(g.PaddingY)+tileSize*7),
			NewBishop(true, whitePieces, float64(g.PaddingX)+tileSize*5, float64(g.PaddingY)+tileSize*7),
			NewKnight(true, whitePieces, float64(g.PaddingX)+tileSize*6, float64(g.PaddingY)+tileSize*7),
			NewRook(true, whitePieces, float64(g.PaddingX)+tileSize*7, float64(g.PaddingY)+tileSize*7),
			NewPawn(whitePieces, float64(g.PaddingX), float64(g.PaddingY)+tileSize*6),
			NewPawn(whitePieces, float64(g.PaddingX)+tileSize, float64(g.PaddingY)+tileSize*6),
			NewPawn(whitePieces, float64(g.PaddingX)+tileSize*2, float64(g.PaddingY)+tileSize*6),
			NewPawn(whitePieces, float64(g.PaddingX)+tileSize*3, float64(g.PaddingY)+tileSize*6),
			NewPawn(whitePieces, float64(g.PaddingX)+tileSize*4, float64(g.PaddingY)+tileSize*6),
			NewPawn(whitePieces, float64(g.PaddingX)+tileSize*5, float64(g.PaddingY)+tileSize*6),
			NewPawn(whitePieces, float64(g.PaddingX)+tileSize*6, float64(g.PaddingY)+tileSize*6),
			NewPawn(whitePieces, float64(g.PaddingX)+tileSize*7, float64(g.PaddingY)+tileSize*6),
		},
	}

	bPlayer := &Player{
		Pieces: []*Piece{
			NewRook(false, blackPieces, float64(g.PaddingX), float64(g.PaddingY)),
			NewKnight(false, blackPieces, float64(g.PaddingX)+tileSize, float64(g.PaddingY)),
			NewBishop(false, blackPieces, float64(g.PaddingX)+tileSize*2, float64(g.PaddingY)),
			NewQueen(false, blackPieces, float64(g.PaddingX)+tileSize*3, float64(g.PaddingY)),
			NewKing(false, blackPieces, float64(g.PaddingX)+tileSize*4, float64(g.PaddingY)),
			NewBishop(false, blackPieces, float64(g.PaddingX)+tileSize*5, float64(g.PaddingY)),
			NewKnight(false, blackPieces, float64(g.PaddingX)+tileSize*6, float64(g.PaddingY)),
			NewRook(false, blackPieces, float64(g.PaddingX)+tileSize*7, float64(g.PaddingY)),
			NewPawn(blackPieces, float64(g.PaddingX), float64(g.PaddingY)+tileSize),
			NewPawn(blackPieces, float64(g.PaddingX)+tileSize, float64(g.PaddingY)+tileSize),
			NewPawn(blackPieces, float64(g.PaddingX)+tileSize*2, float64(g.PaddingY)+tileSize),
			NewPawn(blackPieces, float64(g.PaddingX)+tileSize*3, float64(g.PaddingY)+tileSize),
			NewPawn(blackPieces, float64(g.PaddingX)+tileSize*4, float64(g.PaddingY)+tileSize),
			NewPawn(blackPieces, float64(g.PaddingX)+tileSize*5, float64(g.PaddingY)+tileSize),
			NewPawn(blackPieces, float64(g.PaddingX)+tileSize*6, float64(g.PaddingY)+tileSize),
			NewPawn(blackPieces, float64(g.PaddingX)+tileSize*7, float64(g.PaddingY)+tileSize),
		},
	}

	g.WPlayer = wPlayer
	g.BPlayer = bPlayer

	wSquare := ebiten.NewImage(tileSize, tileSize)
	wSquare.Fill(color.RGBA{100, 215, 255, 255})
	g.WSquare = wSquare

	bSquare := ebiten.NewImage(tileSize, tileSize)
	bSquare.Fill(color.RGBA{155, 0, 180, 255})
	g.BSquare = bSquare

	engine := gochess.NewChessEngine()
	id, err := engine.NewGame("white", "black")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(id)

	board := engine.GetBoard()

	g.Board = board.GetGrid()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
