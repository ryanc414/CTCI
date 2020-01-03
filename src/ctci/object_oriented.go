package ctci

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
