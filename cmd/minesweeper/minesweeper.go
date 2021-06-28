package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/ryanc414/ctci/pkg/objects"
)

// Seend the RNG and play a new game.
func main() {
	objects.SeedRng()
	game := InitGame(10, 6)
	game.Play()
}

// Represents all game state.
type Game struct {
	grid   Grid
	status GameStatus
}

// A game can either be in progres, or finished in a win or lose state.
type GameStatus int

const (
	InProgress = iota
	GameWon
	GameLost
)

// Every turn, the user may take one of two actions: explore or flag.
type GameAction int

const (
	Explore = iota
	Flag
)

type Grid [][]Cell

// Each cell in the grid may be in a combination of states. We use bitflags
// to store the states, we could have also stored a struct of bools.
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
		coords := game.getCoordsInput()
		game.applyAction(action, coords)
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
			builder.WriteRune(game.grid.CellChar(
				objects.GridCoords{Row: row, Col: col},
			))
			builder.WriteRune(' ')
		}
		builder.WriteRune('\n')
	}

	fmt.Print(builder.String())
}

// Get a valid row and column input from the user.
func (game Game) getCoordsInput() objects.GridCoords {
	coords := promptCoords()
	for !game.grid.validCoords(coords) ||
		game.grid[coords.Row][coords.Col]&Explored != 0 {
		fmt.Println("Invalid row/col, try again.")
		coords = promptCoords()
	}

	return coords
}

// Apply a game action to a specified cell.
func (game Game) applyAction(action GameAction, coords objects.GridCoords) {
	switch action {
	case Explore:
		game.grid[coords.Row][coords.Col] |= Explored
		if game.grid[coords.Row][coords.Col]&Bomb == 0 &&
			game.grid.countNeighbourBombs(coords) == 0 {
			game.exploreNeighbours(coords)
		}

	case Flag:
		if game.grid[coords.Row][coords.Col]&Flagged == 0 {
			game.grid[coords.Row][coords.Col] |= Flagged
		} else {
			game.grid[coords.Row][coords.Col] &= ^Flagged
		}

	default:
		panic("Invalid action")
	}
}

// Explore all neighbours, when it is known there are no bombs.
func (game Game) exploreNeighbours(coords objects.GridCoords) {
	for i := range objects.GridDirections {
		newCoords := coords.MoveDirection(objects.GridDirections[i])
		if game.grid.validCoords(newCoords) &&
			game.grid[newCoords.Row][newCoords.Col]&Explored == 0 {
			game.applyAction(Explore, newCoords)
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
func (grid Grid) validCoords(coords objects.GridCoords) bool {
	if coords.Row < 0 || coords.Row >= len(grid) {
		return false
	}

	if coords.Col < 0 || coords.Col >= len(grid[0]) {
		return false
	}

	return true
}

// Return a character to represent a cell in the grid.
func (grid Grid) CellChar(coords objects.GridCoords) rune {
	cell := grid[coords.Row][coords.Col]
	if cell&Explored != 0 {
		if cell&Bomb != 0 {
			return 'X'
		} else {
			neighbourBombs := grid.countNeighbourBombs(coords)
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
func (grid Grid) countNeighbourBombs(coords objects.GridCoords) int {
	bombCount := 0

	for i := range objects.GridDirections {
		newCoords := coords.MoveDirection(objects.GridDirections[i])
		if grid.validCoords(newCoords) &&
			grid[newCoords.Row][newCoords.Col]&Bomb != 0 {
			bombCount++
		}
	}

	return bombCount
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
func promptCoords() objects.GridCoords {
	fmt.Println("Enter row and column:")
	row := getIntInput("row:\n> ")
	col := getIntInput("col:\n> ")

	return objects.GridCoords{Row: row, Col: col}
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
