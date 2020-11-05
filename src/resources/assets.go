package resources

import (
	"bytes"
	"image/color"
	"image/png"

	"github.com/aintsashqa/game-of-life/src/config"
	"github.com/hajimehoshi/ebiten"
)

func LoadDefaultImage() *ebiten.Image {
	currentBytes, err := Asset(config.Load().DefaultImageFilename)
	if err != nil {
		panic(err)
	}

	currentImage, err := png.Decode(bytes.NewReader(currentBytes))
	if err != nil {
		panic(err)
	}

	current, err := ebiten.NewImageFromImage(currentImage, ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}

	return current
}

func LoadDefaultBorder() *ebiten.Image {
	current, err := ebiten.NewImage(1, 1, ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}

	current.Fill(color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	})

	return current
}
