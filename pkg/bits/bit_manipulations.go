package bits

import (
	"errors"
	"fmt"
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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
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

// Swap odd and even position bits in an integer.
func SwapPairs(x int32) int32 {
	evenComb := generateComb()
	oddComb := (evenComb << 1) | 1

	oddBits := x & oddComb
	evenBits := x & evenComb

	evenBits >>= 1
	oddBits <<= 1

	return evenBits | oddBits
}

// Generate an alternating comb of 1s and 0s
func generateComb() int32 {
	var comb int32 = 0

	for i := 0; i < 31; i++ {
		if i%2 == 0 {
			comb |= 1
		}
		comb <<= 1
	}

	return comb
}

// Draw a horizontal line on a "screen" represented by an array of bytes,
// where each bit in the array represents a single monochrome pixel.
func DrawLine(screen []byte, width, x1, x2, y int) error {
	// Error checking inputs.
	if x1 < 0 || x1 >= width || x1 > x2 {
		return errors.New("Invalid x1 value")
	}

	if x2 < 0 || x2 > width {
		return errors.New("Invalid x2 value")
	}

	if width%8 != 0 {
		return errors.New("Invalid width")
	}

	height := len(screen) / (width / 8)

	if y < 0 || y >= height {
		return errors.New("Invalid y")
	}

	rowStart := y * (width / 8)
	startByte := rowStart + (x1 / 8)
	endByte := rowStart + (x2 / 8)

	if startByte == endByte {
		screen[startByte] |= getStartEndByte(x1, x2)
	} else {
		screen[startByte] |= getStartByte(x1)

		for i := startByte + 1; i < endByte; i++ {
			screen[i] = 0xff
		}

		screen[endByte] |= getEndByte(x2)
	}

	return nil
}

// Return the byte for a line that lies entirely within a single byte
// boundary.
func getStartEndByte(x1, x2 int) byte {
	return getStartByte(x1) & getEndByte(x2)
}

func getStartByte(x1 int) byte {
	if x1%8 == 0 {
		return 0xff
	} else {
		return (0x01 << (8 - (x1 % 8))) - 1
	}
}

func getEndByte(x2 int) byte {
	if x2%8 == 7 {
		return 0xff
	} else {
		return ^getStartByte(x2 + 1)
	}
}

// Print out a "screen" of pixels.
func PrintScreen(screen []byte, width int) error {
	if width%8 != 0 {
		return errors.New("Invalid width")
	}

	var sb strings.Builder
	numRows := len(screen) / (width / 8)
	numCols := width / 8

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			writeScreenByte(&sb, screen[(row*numCols)+col])
		}
		sb.WriteRune('\n')
	}

	fmt.Print(sb.String())

	return nil
}

func writeScreenByte(sb *strings.Builder, screenByte byte) {
	var mask byte = 0x80

	for i := 0; i < 8; i++ {
		if mask&screenByte == 0 {
			sb.WriteRune('.')
		} else {
			sb.WriteRune('-')
		}
		mask >>= 1
	}
}
