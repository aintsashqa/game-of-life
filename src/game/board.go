package game

import (
	"fmt"

	"github.com/aintsashqa/game-of-life/src/config"
	"github.com/aintsashqa/game-of-life/src/resources"
	"github.com/hajimehoshi/ebiten"
)

type Board struct {
	CurrentImage    *ebiten.Image
	BorderImage     *ebiten.Image
	Tiles           [][]bool
	Width           int
	Height          int
	GenerationCount int
}

func NewBoard() Board {
	board := Board{}

	board.CurrentImage = resources.LoadDefaultImage()
	board.BorderImage = resources.LoadDefaultBorder()
	board.Width = config.Load().BoardWidth
	board.Height = config.Load().BoardHeight
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

func (b *Board) Reset() {
	b.fill()
	b.GenerationCount = 0
}

func (b *Board) SwitchTile(x, y int) {
	x = x / config.Load().DefaultImageWidth
	y = y / config.Load().DefaultImageHeight

	if x >= b.Width || y >= b.Height {
		fmt.Println("Invalid coordinates")
		return
	}

	b.Tiles[y][x] = !b.Tiles[y][x]
}

func (b *Board) Render(screen *ebiten.Image) error {
	for row := 0; row < b.Height; row++ {

		borderOption := &ebiten.DrawImageOptions{}
		borderScaleX := float64(config.Load().WindowWidth)
		borderPositionY := float64(row * config.Load().DefaultImageHeight)
		borderOption.GeoM.Scale(borderScaleX, 1)
		borderOption.GeoM.Translate(0, borderPositionY)
		screen.DrawImage(b.BorderImage, borderOption)

		for column := 0; column < b.Width; column++ {

			borderOption := &ebiten.DrawImageOptions{}
			borderScaleY := float64(config.Load().WindowHeight)
			borderPositionX := float64(column * config.Load().DefaultImageWidth)
			borderOption.GeoM.Scale(1, borderScaleY)
			borderOption.GeoM.Translate(borderPositionX, 0)
			screen.DrawImage(b.BorderImage, borderOption)

			if !b.Tiles[row][column] {
				continue
			}

			options := &ebiten.DrawImageOptions{}

			x := column*config.Load().DefaultImageWidth + column*config.Load().DefaultImageMargin
			y := row*config.Load().DefaultImageHeight + row*config.Load().DefaultImageMargin
			options.GeoM.Translate(float64(x), float64(y))

			screen.DrawImage(b.CurrentImage, options)
		}
	}

	return nil
}
