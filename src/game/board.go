package game

import (
	"fmt"
	"image"

	"github.com/aintsashqa/game-of-life/src/config"
	"github.com/aintsashqa/game-of-life/src/resources"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Board struct {
	CurrentImage    *ebiten.Image
	BorderImage     *ebiten.Image
	RenderRect      image.Rectangle
	Tiles           [][]bool
	Width           int
	Height          int
	Scale           float64
	GenerationCount int
}

func NewBoard(width, height int) Board {
	board := Board{}

	board.CurrentImage = resources.LoadDefaultImage()
	board.BorderImage = resources.LoadDefaultBorder()
	board.RenderRect = image.Rect(0, 0, config.Load().BoardWidth, config.Load().BoardHeight)
	board.Width = width
	board.Height = height
	board.Scale = config.Load().BoardScale
	board.GenerationCount = 0

	board.fill()

	return board
}

func (b *Board) fill() {
	var tiles [][]bool

	for row := 0; row < b.Height; row++ {
		var currentRow []bool
		for column := 0; column < b.Width; column++ {
			currentRow = append(currentRow, false)
		}

		tiles = append(tiles, currentRow)
	}

	b.Tiles = tiles
}

func (b *Board) Render(screen *ebiten.Image) error {
	for row := 0; row < b.RenderRect.Dy(); row++ {

		borderOption := &ebiten.DrawImageOptions{}
		borderScaleX := float64(config.Load().DefaultImageWidth * b.RenderRect.Dx())
		borderPositionY := float64(row * config.Load().DefaultImageHeight)
		borderOption.GeoM.Scale(borderScaleX, 1)
		borderOption.GeoM.Translate(0, borderPositionY)
		screen.DrawImage(b.BorderImage, borderOption)

		for column := 0; column < b.RenderRect.Dx(); column++ {

			borderOption := &ebiten.DrawImageOptions{}
			borderScaleY := float64(config.Load().DefaultImageHeight * b.RenderRect.Dy())
			borderPositionX := float64(column * config.Load().DefaultImageWidth)
			borderOption.GeoM.Scale(1, borderScaleY)
			borderOption.GeoM.Translate(borderPositionX, 0)
			screen.DrawImage(b.BorderImage, borderOption)

			if !b.Tiles[row+b.RenderRect.Min.Y][column+b.RenderRect.Min.X] {
				continue
			}

			options := &ebiten.DrawImageOptions{}

			x := column*config.Load().DefaultImageWidth + column*config.Load().DefaultImageMargin
			y := row*config.Load().DefaultImageHeight + row*config.Load().DefaultImageMargin
			options.GeoM.Scale(b.Scale, b.Scale)
			options.GeoM.Translate(float64(x), float64(y))

			screen.DrawImage(b.CurrentImage, options)
		}
	}

	return nil
}

func (b *Board) Update() {
	nextRenderRect := b.RenderRect

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		x = x / config.Load().DefaultImageWidth
		y = y / config.Load().DefaultImageHeight

		if x >= b.Width || y >= b.Height {
			fmt.Println("Invalid coordinates")
			return
		}

		b.Tiles[y][x] = !b.Tiles[y][x]
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyR) {
		b.fill()
		b.GenerationCount = 0
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		// TODO: Add +Scale
		fmt.Println("Pressed key Q")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		// TODO: Add -Scale
		fmt.Println("Pressed key E")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		temp := nextRenderRect.Add(image.Pt(1, 0))
		if temp.Max.X <= b.Width {
			nextRenderRect = temp
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		temp := nextRenderRect.Add(image.Pt(-1, 0))
		if temp.Min.X >= 0 {
			nextRenderRect = temp
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		temp := nextRenderRect.Add(image.Pt(0, 1))
		if temp.Max.Y <= b.Height {
			nextRenderRect = temp
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		temp := nextRenderRect.Add(image.Pt(0, -1))
		if temp.Min.Y >= 0 {
			nextRenderRect = temp
		}
	}

	b.RenderRect = nextRenderRect
}
