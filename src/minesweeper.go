package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	game := InitGame(8, 4)
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

type GameAction int

const (
	Explore = iota
	Flag
)

type Grid [][]Cell

// Represents a direction of movement on the grid.
type Direction int

const (
	Up = iota
	UpRight
	Right
	DownRight
	Down
	DownLeft
	Left
	UpLeft
)

var Directions = [...]Direction{
	Up, UpRight, Right, DownRight, Down, DownLeft, Left, UpLeft,
}

type Cell int

const (
	Bomb     = 0x1
	Explored = 0x2
	Flagged  = 0x4
)

// Initialise a new game object.
func InitGame(gridSize, numBombs int) *Game {
	return &Game{
		grid:   InitGrid(gridSize, numBombs),
		status: InProgress,
	}
}

// Play a new game.
func (game Game) Play() {
	for game.status == InProgress {
		game.display()
		action := promptAction()
		row, col := game.getRowColInput()
		game.applyAction(action, row, col)
		game.status = game.getStatus()
	}

	game.display()
	game.printEndStatus()
}

// Display the game.
func (game Game) display() {
	var builder strings.Builder

	builder.WriteString("  ")
	for i := range game.grid {
		builder.WriteString(strconv.Itoa(i))
		builder.WriteRune(' ')
	}
	builder.WriteRune('\n')

	for row := range game.grid {
		builder.WriteString(strconv.Itoa(row))
		builder.WriteRune(' ')

		for col := range game.grid[row] {
			builder.WriteRune(game.grid.CellChar(row, col))
			builder.WriteRune(' ')
		}
		builder.WriteRune('\n')
	}

	fmt.Print(builder.String())
}

// Get a valid row and column input from the user.
func (game Game) getRowColInput() (int, int) {
	row, col := promptRowCol()
	for !game.grid.validRowCol(row, col) ||
		game.grid[row][col]&Explored != 0 {
		fmt.Println("Invalid row/col, try again.")
		row, col = promptRowCol()
	}

	return row, col
}

// Apply a game action to a specified cell.
func (game Game) applyAction(action GameAction, row, col int) {
	switch action {
	case Explore:
		game.grid[row][col] |= Explored
		if game.grid[row][col]&Bomb != 0 {
			game.status = GameLost
		} else {
			if game.grid.countNeighbourBombs(row, col) == 0 {
				game.exploreNeighbours(row, col)
			}
		}

	case Flag:
		if game.grid[row][col]&Flagged == 0 {
			game.grid[row][col] |= Flagged
		} else {
			game.grid[row][col] &= ^Flagged
		}

	default:
		panic("Invalid action")
	}
}

// Explore all neighbours, when it is known there are no bombs.
func (game Game) exploreNeighbours(row, col int) {
	for i := range Directions {
		newRow, newCol := game.grid.Move(row, col, Directions[i])
		if game.grid.validRowCol(newRow, newCol) &&
			game.grid[newRow][newCol]&Explored == 0 {
			game.applyAction(Explore, newRow, newCol)
		}
	}
}

// Print the outcome at the end of the game.
func (game Game) printEndStatus() {
	switch game.status {
	case GameWon:
		fmt.Println("You win - congrats!")

	case GameLost:
		fmt.Println("BOOM: you lose.")

	default:
		panic("Unexpected game status")
	}
}

// Return the current game status based on the grid state.
func (game Game) getStatus() GameStatus {
	numUnexplored := 0

	for row := range game.grid {
		for col := range game.grid {
			cell := game.grid[row][col]

			if cell&Explored != 0 && cell&Bomb != 0 {
				return GameLost
			}

			if cell&(Explored|Flagged) == 0 {
				numUnexplored++
			}
		}
	}

	if numUnexplored > 0 {
		return InProgress
	} else {
		return GameWon
	}
}

// Initialise a new grid.
func InitGrid(size, numBombs int) Grid {
	grid := make(Grid, size)
	for i := range grid {
		grid[i] = make([]Cell, size)
	}

	bombsPlaced := 0
	for bombsPlaced != numBombs {
		row := rand.Intn(size)
		col := rand.Intn(size)
		if grid[row][col]&Bomb == 0 {
			grid[row][col] |= Bomb
			bombsPlaced++
		}
	}

	return grid
}

// Check if row and column indices are valid.
func (grid Grid) validRowCol(row, col int) bool {
	if row < 0 || row >= len(grid) {
		return false
	}

	if col < 0 || col >= len(grid[0]) {
		return false
	}

	return true
}

// Return a character to represent a cell in the grid.
func (grid Grid) CellChar(row, col int) rune {
	cell := grid[row][col]
	if cell&Explored != 0 {
		if cell&Bomb != 0 {
			return 'X'
		} else {
			neighbourBombs := grid.countNeighbourBombs(row, col)
			if neighbourBombs == 0 {
				return ' '
			} else {
				return []rune(strconv.Itoa(neighbourBombs))[0]
			}
		}
	} else {
		if cell&Flagged != 0 {
			return 'F'
		} else {
			return '.'
		}
	}
}

// Count the number of neighbouring bombs.
func (grid Grid) countNeighbourBombs(row, col int) int {
	bombCount := 0

	for i := range Directions {
		newRow, newCol := grid.Move(row, col, Directions[i])
		if grid.validRowCol(newRow, newCol) && grid[newRow][newCol]&Bomb != 0 {
			bombCount++
		}
	}

	return bombCount
}

// Return the row/col indices of the adjacent cell in a given direction.
func (grid Grid) Move(row, col int, direction Direction) (int, int) {
	switch direction {
	case Up:
		return row - 1, col

	case UpRight:
		return row - 1, col + 1

	case Right:
		return row, col + 1

	case DownRight:
		return row + 1, col + 1

	case Down:
		return row + 1, col

	case DownLeft:
		return row + 1, col - 1

	case Left:
		return row, col - 1

	case UpLeft:
		return row - 1, col - 1

	default:
		panic("Unexpected direction.")
	}
}

// Convert an input string to an Action value.
func toAction(actionInput string) (GameAction, error) {
	normalised := strings.ToUpper(strings.TrimSuffix(actionInput, "\n"))
	if len(normalised) < 1 {
		return -1, errors.New("No action specified")
	}

	switch normalised[0] {
	case 'E':
		return Explore, nil

	case 'F':
		return Flag, nil

	default:
		return -1, errors.New("Invalid action")
	}
}

// Prompt user to select an action.
func promptAction() GameAction {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select an action ((E)xplore or (F)lag)\n> ")
	actionInput, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	action, actionErr := toAction(actionInput)

	for actionErr != nil {
		fmt.Print(
			"Invalid action, please try again.\n" +
				"Valid actions are (E)xplore or (F)lag\n> ",
		)
		actionInput, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		action, actionErr = toAction(actionInput)
	}

	return action
}

// Prompt for user to enter a row and column
func promptRowCol() (int, int) {
	fmt.Println("Enter row and column:")
	row := getIntInput("row:\n> ")
	col := getIntInput("col:\n> ")

	return row, col
}

// Get an integer input from user.
func getIntInput(prompt string) int {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	rawInput, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	intInput, convertErr := strconv.Atoi(strings.TrimSuffix(rawInput, "\n"))

	for convertErr != nil {
		fmt.Println("Invalid integer, please try again.")
		fmt.Print(prompt)
		rawInput, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		intInput, convertErr = strconv.Atoi(strings.TrimSuffix(rawInput, "\n"))
	}

	return intInput
}
