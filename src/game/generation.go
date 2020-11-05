package game

import (
	"time"

	"github.com/aintsashqa/game-of-life/src/config"
)

var (
	lastGeneration time.Time
)

func init() {
	lastGeneration = time.Now()
}

func NextGeneration(board *Board) {
	if !time.Now().After(lastGeneration.Add(config.Load().GenerationTimeout)) {
		return
	}
	lastGeneration = time.Now()

	nextGeneration := (NewBoard().Tiles)

	for row := 0; row < board.Height; row++ {
		for column := 0; column < board.Width; column++ {
			count := countOfNearestTiles(column, row, board.Width, board.Height, board.Tiles)

			if board.Tiles[row][column] {
				if canTileLiveInCurrentGeneration(count) {
					nextGeneration[row][column] = true
				}
			} else {
				if canReviveTile(count) {
					nextGeneration[row][column] = true
				}
			}
		}
	}

	board.Tiles = nextGeneration
}

func countOfNearestTiles(x, y, width, height int, tiles [][]bool) int {
	count := 0

	nearestTiles := [8][2]int{
		// Top left
		{-1, -1},
		// Top center
		{0, -1},
		// Top right
		{1, -1},
		// Middle left
		{-1, 0},
		// Middle right
		{1, 0},
		// Botton left
		{-1, 1},
		// Botton center
		{0, 1},
		// Botton right
		{1, 1},
	}

	for _, tile := range nearestTiles {
		currentX := x + tile[0]
		currentY := y + tile[1]

		if currentX < 0 {
			currentX = width - 1
		} else if currentX >= width {
			currentX = 0
		}

		if currentY < 0 {
			currentY = height - 1
		} else if currentY >= height {
			currentY = 0
		}

		if tiles[currentY][currentX] {
			count++
		}
	}

	return count
}

func canTileLiveInCurrentGeneration(count int) bool {
	return count == 2 || count == 3
}

func canReviveTile(count int) bool {
	return count == 3
}
