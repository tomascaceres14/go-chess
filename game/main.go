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
	PieceType gochess.PieceType
	Color     bool
}

func (g Game) GetPiece(pieceType gochess.PieceType, color bool) *Piece {

	for _, v := range g.Pieces {
		if v.PieceType == pieceType && color == v.Color {
			return v
		}

	}

	return nil
}

type Game struct {
	Pieces             []*Piece
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

func NewPiece2(color bool, pieceType gochess.PieceType, img *ebiten.Image) *Piece {
	return &Piece{
		Sprite:    CenteredPieceSprite2(img),
		Color:     color,
		PieceType: pieceType,
	}
}

// cuts img and extracts piece sub image
func PieceSubImg(img *ebiten.Image, x0, y0, x1, y1 int) *ebiten.Image {
	return img.SubImage(
		image.Rect(x0, y0, x1, y1),
	).(*ebiten.Image)
}

func NewPawn(color bool, img *ebiten.Image) *Piece {

	X0 := 278
	Y0 := 136
	X1 := 356
	Y1 := 256

	subImg := PieceSubImg(img, X0, Y0, X1, Y1)

	p := NewPiece2(color, gochess.PawnType, subImg)

	return p
}

func NewBishop(color bool, img *ebiten.Image) *Piece {

	X0 := 127
	Y0 := 2
	X1 := 226
	Y1 := 117

	subImg := PieceSubImg(img, X0, Y0, X1, Y1)

	p := NewPiece2(color, gochess.BishopType, subImg)

	return p

}

func NewKnight(color bool, img *ebiten.Image) *Piece {

	X0 := 129
	Y0 := 140
	X1 := 218
	Y1 := 256

	subImg := PieceSubImg(img, X0, Y0, X1, Y1)

	p := NewPiece2(color, gochess.KnightType, subImg)

	return p

}

func NewRook(color bool, img *ebiten.Image) *Piece {

	x0 := 0
	y0 := 0
	x1 := 82
	y1 := 117

	subImg := PieceSubImg(img, x0, y0, x1, y1)
	p := NewPiece2(color, gochess.RookType, subImg)

	return p
}

func NewQueen(color bool, img *ebiten.Image) *Piece {
	X0 := 3
	Y0 := 136
	X1 := 79
	Y1 := 256

	subImg := PieceSubImg(img, X0, Y0, X1, Y1)
	p := NewPiece2(color, gochess.QueenType, subImg)

	return p
}

func NewKing(color bool, img *ebiten.Image) *Piece {
	X0 := 279
	Y0 := 0
	X1 := 356
	Y1 := 117

	subImg := PieceSubImg(img, X0, Y0, X1, Y1)

	p := NewPiece2(color, gochess.KingType, subImg)

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

	g.Pieces = []*Piece{
		NewRook(true, whitePieces),
		NewKnight(true, whitePieces),
		NewBishop(true, whitePieces),
		NewQueen(true, whitePieces),
		NewKing(true, whitePieces),
		NewPawn(true, whitePieces),
		NewRook(false, blackPieces),
		NewKnight(false, blackPieces),
		NewBishop(false, blackPieces),
		NewQueen(false, blackPieces),
		NewKing(false, blackPieces),
		NewPawn(false, blackPieces),
	}

	wSquare := ebiten.NewImage(tileSize, tileSize)
	wSquare.Fill(color.RGBA{100, 215, 255, 255})
	g.WSquare = wSquare

	bSquare := ebiten.NewImage(tileSize, tileSize)
	bSquare.Fill(color.RGBA{155, 0, 180, 255})
	g.BSquare = bSquare

	engine := gochess.NewChessEngine()
	id, err := engine.NewGameFENString("white", "black", "r1bqk2r/pppp1ppp/2n2n2/2b1p3/2B1P3/3P1N2/PPP2PPP/RNBQ1RK1 b kq - 2 5")
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
