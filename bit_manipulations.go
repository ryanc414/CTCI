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

// Get the next smallest integer with the same number of binary 1s.
func GetNext(x int32) (int32, error) {
	tmp := x
	c0 := 0
	c1 := 0

	for ((tmp & 1) == 0) && (tmp != 0) {
		c0++
		tmp >>= 1
	}

	for (tmp & 1) == 1 {
		c1++
		tmp >>= 1
	}

	// If x == 11..1100..00 then there is no bigger number with the same number
	// of 1s
	if c0+c1 == 31 || c0+c1 == 0 {
		return 0, errors.New("There is no next number.")
	}

	return x + (1 << c0) + (1 << (c1 - 1)) - 1, nil
}

// Get the previous largest integer with the same number of binary 1s.
func GetPrev(x int32) (int32, error) {
	tmp := x
	c0 := 0
	c1 := 0

	for (tmp & 1) == 1 {
		c1++
		tmp >>= 1
	}

	if tmp == 0 {
		return 0, errors.New("There is no previous number.")
	}

	for (tmp&1) == 0 && tmp != 0 {
		c0++
		tmp >>= 1
	}

	return x - (1 << c1) - (1 << (c0 - 1)) + 1, nil
}

// Calculate the number of bit flips required to turn integer x into integer y.
func FlipsRequired(x, y int32) int {
	numFlips := 0

	for x != y {
		if (x & 1) != (y & 1) {
			numFlips++
		}
		x >>= 1
		y >>= 1
	}

	return numFlips
}
