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

func FlipBitToWin(x int32) int {
	// As a special case, return 32 for -1 (all ones) since there are no zero
	// bits that can be flipped.
	if x == -1 {
		return 32
	}

	bits := intToBits(x)
	groups := bitsToGroups(bits[:])

	if len(groups) > 1 {
		return sumLargestPair(groups) + 1
	} else if len(groups) == 1 {
		return groups[0] + 1
	} else {
		// Must be all zeroes, so one bit may be flipped.
		return 1
	}
}

// Convert a 32-bit integer to an array of bools indicating the bits.
func intToBits(x int32) [32]bool {
	var bits [32]bool

	for i := 0; i < 32; i++ {
		if x == 0 {
			break
		}

		if x%2 == 1 {
			bits[31-i] = true
		}

		x >>= 1
	}

	return bits
}

// Convert a slice of bits to a slice containing the size of contiguous groups
// of 1s.
func bitsToGroups(bits []bool) []int {
	var groups []int
	currGroupSize := 0

	for i := range bits {
		if bits[i] {
			currGroupSize++
		} else if currGroupSize > 0 {
			groups = append(groups, currGroupSize)
			currGroupSize = 0
		}
	}

	if currGroupSize > 0 {
		groups = append(groups, currGroupSize)
		currGroupSize = 0
	}

	return groups
}

// Return the largest sum of an adjacent pair of values in a slice.
func sumLargestPair(vals []int) int {
	largestPair := 0

	for i := 0; i < len(vals)-1; i++ {
		pairSum := vals[i] + vals[i+1]
		if pairSum > largestPair {
			largestPair = pairSum
		}
	}

	return largestPair
}
