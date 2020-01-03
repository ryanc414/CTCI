package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Play a single game of Othello.
func main() {
	board := InitBoard()
	board.PlayGame()
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

// Represents a direction of movement on the board.
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
	board := &Board{
		players: [2]Player{
			InitHumanPlayer(Black),
			InitHumanPlayer(White),
		},
		currTurn: 0,
		status:   InProgress,
	}

	// Place the initial pieces: two white and two black in the centre of the
	// board.
	board.grid[3][3] = &Piece{
		colour: White, position: Coords{x: 3, y: 3},
	}
	board.grid[4][4] = &Piece{
		colour: White, position: Coords{x: 4, y: 4},
	}
	board.grid[3][4] = &Piece{
		colour: Black, position: Coords{x: 3, y: 4},
	}
	board.grid[4][3] = &Piece{
		colour: Black, position: Coords{x: 4, y: 3},
	}

	return board
}

// Display the board. Currently just print a text representation to the
// terminal. Could provide generic interface to plug in other board renderers
// in future.
func (board *Board) Display() {
	var builder strings.Builder
	// TODO could grow builder up front to save allocations.

	builder.WriteString("  ")
	for col := range board.grid[0] {
		builder.WriteString(strconv.Itoa(col))
		builder.WriteRune(' ')
	}
	builder.WriteRune('\n')

	for row := range board.grid {
		builder.WriteString(strconv.Itoa(row))
		builder.WriteRune(' ')

		for col := range board.grid[row] {
			builder.WriteRune(board.grid[row][col].DisplaySymbol())
			builder.WriteRune(' ')
		}

		builder.WriteRune('\n')
	}

	fmt.Println(builder.String())
}

// Start a new game.
func (board *Board) PlayGame() {
	for board.status == InProgress {
		board.Display()
		currPlayer := board.players[board.currTurn]
		currColour := board.getCurrColour()
		nextMove := currPlayer.ChooseMove(board)
		board.placePiece(nextMove, currColour)

		if board.NoMovesPossible(currColour) {
			board.status = board.GetGameResult()
		} else {
			board.currTurn = (board.currTurn + 1) % 2
		}
	}

	board.Display()
	board.printStatus()
}

// Validate a move.
func (board *Board) ValidMove(move Coords, colour Colour) bool {
	// Bounds checking.
	if !board.checkBounds(move) {
		return false
	}

	// Check if space is occupied by another piece.
	if board.grid[move.y][move.x] != nil {
		return false
	}

	// Check if placing a piece here will form a terminated run of the opposite
	// colour in any direction.
	for i := range Directions {
		if board.runExists(move, colour, Directions[i]) {
			return true
		}
	}

	// No run exists, so not a valid move.
	return false
}

// Get the colour of the current player.
func (board *Board) getCurrColour() Colour {
	switch board.currTurn {
	case 0:
		return Black

	case 1:
		return White

	default:
		panic("Unexpected turn value")
	}
}

// Place a new piece on the board.
func (board *Board) placePiece(nextMove Coords, colour Colour) {
	// The move should have already been validated, but check again for sanity.
	if !board.ValidMove(nextMove, colour) {
		panic("Invalid move")
	}

	piece := &Piece{colour: colour, position: nextMove}
	board.grid[nextMove.y][nextMove.x] = piece
	board.flipAllRuns(piece)
}

// Flip all pieces of opposite colour that form terminated runs adjacent to
// this new piece.
func (board *Board) flipAllRuns(piece *Piece) {
	for i := range Directions {
		if board.runExists(piece.position, piece.colour, Directions[i]) {
			board.flipRun(piece.position, piece.colour, Directions[i])
		}
	}
}

// Flip all consecutive pieces of the opposite colour in a given direction.
func (board *Board) flipRun(
	position Coords, colour Colour, direction Direction,
) {
	position = position.MoveDirection(direction)
	piece := board.grid[position.y][position.x]

	for piece.colour != colour {
		piece.colour = colour
		position = position.MoveDirection(direction)
		piece = board.grid[position.y][position.x]
	}
}

// Check if a given position is within the grid boundaries.
func (board *Board) checkBounds(position Coords) bool {
	if position.y < 0 || position.y >= len(board.grid) {
		return false
	}

	if position.x < 0 || position.x >= len(board.grid[0]) {
		return false
	}

	return true
}

// Check if one or more pieces of opposite colour exist in a given direction,
// terminated by a piece of our colour.
func (board *Board) runExists(
	move Coords, colour Colour, direction Direction,
) bool {
	// First check if our immediate neighbour is a piece of the opposite
	// colour. If not, there is no run so return false.
	position := move.MoveDirection(direction)
	if !board.checkBounds(position) {
		return false
	}

	piece := board.grid[position.y][position.x]
	if piece == nil || piece.colour == colour {
		return false
	}

	// Now, iterate in that direction until we reach either a blank space or
	// a piece of our colour.
	for piece != nil && piece.colour != colour {
		position = position.MoveDirection(direction)
		if !board.checkBounds(position) {
			return false
		}
		piece = board.grid[position.y][position.x]
	}

	// If piece is non-nil, it must be of the opposite colour so we have
	// found a run.
	return piece != nil
}

// Return true if no more moves are possible for the current colour.
func (board *Board) NoMovesPossible(colour Colour) bool {
	for col := range board.grid {
		for row := range board.grid[col] {
			if board.ValidMove(Coords{x: col, y: row}, colour) {
				return false
			}
		}
	}

	return true
}

// Get the end result of a game - either a win for black or white, or a draw.
func (board *Board) GetGameResult() GameStatus {
	numBlack := 0
	numWhite := 0

	for row := range board.grid {
		for col := range board.grid[row] {
			if board.grid[row][col] != nil {
				switch board.grid[row][col].colour {
				case Black:
					numBlack++

				case White:
					numWhite++
				}
			}
		}
	}

	if numBlack == numWhite {
		return Draw
	} else if numBlack > numWhite {
		return BlackWin
	} else {
		return WhiteWin
	}
}

// At the end of a game, print the final outcome.
func (board *Board) printStatus() {
	switch board.status {
	case BlackWin:
		fmt.Printf("Black player %v wins!\n", board.players[0].Name())

	case WhiteWin:
		fmt.Printf("White player %v wins!\n", board.players[1].Name())

	case Draw:
		fmt.Println("It's a draw!")

	default:
		panic("Unexpected game status")
	}
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
		name:   strings.TrimSuffix(name, "\n"),
		colour: colour,
	}
}

func (player HumanPlayer) Name() string {
	return player.name
}

// Return a valid next move for this player.
func (player HumanPlayer) ChooseMove(board *Board) Coords {
	fmt.Printf(
		"%v (%v): your turn, please enter coords of next move.\n",
		player.name,
		player.colour.DisplayName(),
	)

	nextMove := player.getCoordsInput()
	for !board.ValidMove(nextMove, player.colour) {
		fmt.Println("Invalid move - please try again.")
		nextMove = player.getCoordsInput()
	}

	return nextMove
}

// Prompt user to enter the coordinates of a move.
func (player HumanPlayer) getCoordsInput() Coords {
	row := player.getIntInput("row: ")
	col := player.getIntInput("col: ")

	return Coords{x: col, y: row}
}

// Get an integer input from a human player.
func (player HumanPlayer) getIntInput(prompt string) int {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)
	inputStr, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	intVal, err := strconv.Atoi(strings.TrimSuffix(inputStr, "\n"))

	for err != nil {
		fmt.Println(err)
		fmt.Println("Invalid number, try again.")
		fmt.Print(prompt)
		inputStr, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		intVal, err = strconv.Atoi(strings.TrimSuffix(inputStr, "\n"))
	}

	return intVal
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

// Return the next consecutive position in the given direction. Note: bounds
// checking is not performed here, so it is up to the caller to check if the
// position is within the grid or not.
func (position Coords) MoveDirection(direction Direction) Coords {
	switch direction {
	case Up:
		return Coords{x: position.x, y: position.y - 1}

	case UpRight:
		return Coords{x: position.x + 1, y: position.y - 1}

	case Right:
		return Coords{x: position.x + 1, y: position.y}

	case DownRight:
		return Coords{x: position.x + 1, y: position.y + 1}

	case Down:
		return Coords{x: position.x, y: position.y + 1}

	case DownLeft:
		return Coords{x: position.x - 1, y: position.y + 1}

	case Left:
		return Coords{x: position.x - 1, y: position.y}

	case UpLeft:
		return Coords{x: position.x - 1, y: position.y - 1}

	default:
		panic("Unexpected direction.")
	}
}
