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
