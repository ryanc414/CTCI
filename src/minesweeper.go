package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	game := InitGame(8)
	game.Play()
}

type Game struct {
	grid   Grid
	status GameStatus
}

type GameStatus int

const (
	InProgress = iota
	GameWon
	GameLost
)

type Grid [][]Cell

type Cell int

const (
	Bomb     = 0x1
	Explored = 0x2
	Flagged  = 0x4
)

// Initialise a new game object.
func InitGame(size int) *Game {
	return &Game{
		grid:   InitGrid(size),
		status: InProgress,
	}
}

// Play a new game.
func (game Game) Play() {
	game.Display()
}

// Display the game.
func (game Game) Display() {
	var builder strings.Builder

	for i := range game.grid {
		builder.WriteString(strconv.Itoa(i))
		builder.WriteRune(' ')
	}
	builder.WriteRune('\n')

	for row := range game.grid {
		builder.WriteString(strconv.Itoa(row))
		builder.WriteRune(' ')

		for col := range game.grid[row] {
			builder.WriteRune(game.grid[row][col].DisplayChar())
			builder.WriteRune(' ')
		}
		builder.WriteRune('\n')
	}

	fmt.Print(builder.String())
}

// Return a character to represent a cell in the grid.
func (grid Grid) CellChar(row, col int) rune {
	cell := grid[row][col]
	if cell&Explored != 0 {
		if cell&Bomb != 0 {
			return 'X'
		} else {
			return strconv.Itoa(grid.countNeighbourBombs(row, col))[0]
		}
	} else {
		if cell&Flagged != 0 {
			return 'F'
		} else {
			return '.'
		}
	}
}

func (grid Grid) countNeighbourBombs(row, col int) int {
	// TODO
	return 0
}

// Initialise a new grid.
func InitGrid(size int) Grid {
	grid := make(Grid, size)
	for i := range grid {
		grid[i] = make([]Cell, size)
		grid[i][rand.Intn(size)] |= Bomb
	}
}
