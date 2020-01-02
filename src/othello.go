package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	board := InitBoard()
	board.Display()
}

// Contains all state for a single game of Othello.
type Board struct {
	grid     [8][8]*Piece
	players  [2]Player
	currTurn int
	status   GameStatus
}

// A game may be either in progress, or finished with either black winning,
// white winning, or a tied result.
type GameStatus int

const (
	InProgress = iota
	BlackWin
	WhiteWin
	Draw
)

// Represents a piece placed on the board. A piece shows either white or black
// as its colour at any one time, but may be flipped to show the opposite
// colour.
type Piece struct {
	colour   Colour
	position Coords
}

type Colour int

const (
	Black = iota
	White
)

type Coords struct {
	x int
	y int
}

type Player interface {
	Name() string
	ChooseMove(board *Board) Coords
}

// Implements the Player interface. Prompts a human for input when making
// moves.
type HumanPlayer struct {
	name   string
	colour Colour
}

// Initialise a fresh board for a new game.
func InitBoard() *Board {
	return &Board{
		players: [2]Player{
			InitHumanPlayer(Black),
			InitHumanPlayer(White),
		},
		currTurn: 0,
		status:   InProgress,
	}
}

// Display the board. Currently just print a text representation to the
// terminal. Could provide generic interface to plug in other board renderers
// in future.
func (board *Board) Display() {
	var builder strings.Builder
	// TODO could grow builder up front to save allocations.

	for row := range board.grid {
		for col := range board.grid[row] {
			builder.WriteRune(board.grid[row][col].DisplaySymbol())
			builder.WriteRune(' ')
		}
		builder.WriteRune('\n')
	}

	fmt.Println(builder.String())
}

// Initialise a new player, prompting to enter their name.
func InitHumanPlayer(colour Colour) HumanPlayer {
	fmt.Printf(
		"Please enter name for the %v player\n> ", colour.DisplayName(),
	)
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return HumanPlayer{
		name:   name,
		colour: colour,
	}
}

func (player HumanPlayer) Name() string {
	return player.name
}

func (player HumanPlayer) ChooseMove(board *Board) Coords {
	// TODO
	return Coords{x: 0, y: 0}
}

// Return a character to represent a single piece for display.
func (piece *Piece) DisplaySymbol() rune {
	if piece == nil {
		return '.'
	}

	switch piece.colour {
	case Black:
		return 'X'

	case White:
		return 'O'

	default:
		panic("Unexpected piece colour")
	}
}

// Return the name of a colour for display.
func (colour Colour) DisplayName() string {
	switch colour {
	case Black:
		return "black"

	case White:
		return "white"

	default:
		panic("Unexpected colour")
	}
}
