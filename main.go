package main

import (
	"fmt"
	"image/color"

	"github.com/aintsashqa/game-of-life/src/config"
	"github.com/aintsashqa/game-of-life/src/game"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var currentBoard game.Board
var isGenerate bool

func init() {
	config.Load()
}

func main() {
	isGenerate = false
	currentBoard = game.NewBoard(16, 12)
	// currentBoard.Tiles[0][1] = true
	// currentBoard.Tiles[1][2] = true
	// currentBoard.Tiles[2][0] = true
	// currentBoard.Tiles[2][1] = true
	// currentBoard.Tiles[2][2] = true

	if err := ebiten.Run(render,
		config.Load().WindowWidth,
		config.Load().WindowHeight,
		config.Load().WindowScale,
		config.Load().WindowTitle,
	); err != nil {
		panic(err)
	}
}

func render(screen *ebiten.Image) error {
	if err := screen.Clear(); err != nil {
		return err
	}

	if err := screen.Fill(color.Gray{Y: 64}); err != nil {
		return err
	}

	// if err := ebitenutil.DebugPrint(screen, fmt.Sprintf("fps: %f", ebiten.CurrentFPS())); err != nil {
	// 	return err
	// }

	if err := ebitenutil.DebugPrint(screen, fmt.Sprintf("Generations: %d", currentBoard.GenerationCount)); err != nil {
		return err
	}

	if err := currentBoard.Render(screen); err != nil {
		return err
	}

	if isGenerate {
		game.NextGeneration(&currentBoard)
	}

	currentBoard.Update()

	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		isGenerate = !isGenerate
	}

	return nil
}
