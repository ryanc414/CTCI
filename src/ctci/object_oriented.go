package ctci

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
)

// The CircularArray type gives an array-like interface but, unlike a regular
// array, it can be efficiently rotated. Can be iterated like:

// for i := circ.First(); i != EndOfArray; i = circ.Next(i) {
//   // do something with circ.Data[i]...
// }
type CircularArray struct {
	Data   []interface{}
	offset int
}

const EndOfArray = -1

func InitCircularArray(data []interface{}) CircularArray {
	return CircularArray{Data: data, offset: 0}
}

// Return the index of the first element in the array
func (circ CircularArray) First() int {
	if len(circ.Data) == 0 {
		return EndOfArray
	} else {
		return circ.offset
	}
}

// Return the next index during iteration, given the current one.
func (circ CircularArray) Next(curr int) int {
	// bounds checking
	if curr < 0 || curr >= len(circ.Data) {
		panic("Index out of range")
	}

	next := (curr + 1) % len(circ.Data)
	if next == circ.offset {
		return EndOfArray
	} else {
		return next
	}
}

// Return a new circular array rotated by the amount specified.
func (circ CircularArray) Rotate(numPlaces int) CircularArray {
	return CircularArray{
		Data:   circ.Data,
		offset: (circ.offset + numPlaces) % len(circ.Data),
	}
}

// Seed the RNG so that different results are produced each time.
func SeedRng() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("Cannot send RNG")
	}

	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

// Represents a direction of movement on a 2D grid.
type GridDirection int

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

var GridDirections = [...]GridDirection{
	Up, UpRight, Right, DownRight, Down, DownLeft, Left, UpLeft,
}

type GridCoords struct {
	Row int
	Col int
}

// Return the next consecutive position in the given direction. Note: bounds
// checking is not performed here, so it is up to the caller to check if the
// position is within the grid or not.
func (position GridCoords) MoveDirection(direction GridDirection) GridCoords {
	switch direction {
	case Up:
		return GridCoords{Row: position.Row - 1, Col: position.Col}

	case UpRight:
		return GridCoords{Row: position.Row - 1, Col: position.Col + 1}

	case Right:
		return GridCoords{Row: position.Row, Col: position.Col + 1}

	case DownRight:
		return GridCoords{Row: position.Row + 1, Col: position.Col + 1}

	case Down:
		return GridCoords{Row: position.Row + 1, Col: position.Col}

	case DownLeft:
		return GridCoords{Row: position.Row + 1, Col: position.Col - 1}

	case Left:
		return GridCoords{Row: position.Row, Col: position.Col - 1}

	case UpLeft:
		return GridCoords{Row: position.Row - 1, Col: position.Col - 1}

	default:
		panic("Unexpected direction.")
	}
}
