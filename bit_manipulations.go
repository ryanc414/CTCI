package ctci

import (
	"errors"
	"strings"
)

// Inserts M into N from indices i to j.
func Insert(N, M int32, i, j uint) int32 {
	mask := ((int32(1) << (j - i)) - 1) << i
	result := N & (^mask)
	result |= (M << i)

	return result
}

// Return a binary representation of a float between 0 and 1 as a string.
func BinaryString(x float32) (string, error) {
	if x < 0.0 || x >= 1.0 {
		return "ERROR", errors.New("x not in range 0 to 1")
	}

	var sb strings.Builder
	sb.WriteRune('.')
	var nextVal float32 = 1.0

	for i := 0; i < 32; i++ {
		nextVal /= 2.0
		if x >= nextVal {
			sb.WriteRune('1')
			x -= nextVal
			if x == 0.0 {
				return sb.String(), nil
			}
		} else {
			sb.WriteRune('0')
		}
	}

	return sb.String(), errors.New("Could not represent in 32 bits")
}

// Return the length of the longest sequence of 1s that can be made by flipping
// at most one bit.
func FlipBitToWin(x int32) int {
	// As a special case, return 32 for -1 (all ones) since there are no zero
	// bits that can be flipped.
	if x == -1 {
		return 32
	}

	currentLength := 0
	prevLength := 0

	// Can always flip at least one bit.
	maxLength := 1

	for x != 0 {
		// Check if current bit is 1.
		if x&1 == 1 {
			currentLength++
		} else {
			if x&2 == 0 {
				prevLength = 0
			} else {
				prevLength = currentLength
			}
			currentLength = 0
		}

		maxLength = max(prevLength+currentLength+1, maxLength)
		x >>= 1
	}

	return maxLength
}

// Return the next-smallest and next-largest integers that can be made with
// the same number of 0s and 1s.
func SwapBits(x int32) (int32, int32, error) {
	var mask int32 = 1
	prevBit := (x & mask) != 0
	mask <<= 1
	swapped := make([]int32, 0, 2)

	for mask != 0 && len(swapped) < 2 {
		currBit := (x & mask) != 0
		if currBit != prevBit {
			swapped = append(swapped, swapMask(x, mask, currBit))
		}
		mask <<= 1
		prevBit = currBit
	}

	if len(swapped) != 2 {
		return 0, 0, errors.New("Could not find all pairs to swap")
	}

	if swapped[0] > swapped[1] {
		return swapped[1], swapped[0], nil
	} else {
		return swapped[0], swapped[1], nil
	}
}

func swapMask(x, mask int32, currBit bool) int32 {
	if currBit {
		swapped := x &^ mask
		mask >>= 1
		swapped |= mask
		return swapped
	} else {
		swapped := x | mask
		mask >>= 1
		swapped &^= mask
		return swapped
	}
}
