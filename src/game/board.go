package game

import (
	"fmt"

	"github.com/aintsashqa/game-of-life/src/config"
	"github.com/aintsashqa/game-of-life/src/resources"
	"github.com/hajimehoshi/ebiten"
)

type Board struct {
	CurrentImage *ebiten.Image
	Tiles        [][]bool
	Width        int
	Height       int
}

func NewBoard() Board {
	board := Board{}

	board.CurrentImage = resources.LoadDefaultImage()
	board.Width = config.Load().BoardWidth
	board.Height = config.Load().BoardHeight

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
}

func (b *Board) AddTile(x, y int) {
	x = x / config.Load().DefaultImageWidth
	y = y / config.Load().DefaultImageHeight

	if x >= b.Width || y >= b.Height {
		fmt.Println("Invalid coordinates")
		return
	}

	b.Tiles[y][x] = true
}

func (b *Board) Render(screen *ebiten.Image) error {
	for row := 0; row < b.Height; row++ {
		for column := 0; column < b.Width; column++ {
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
