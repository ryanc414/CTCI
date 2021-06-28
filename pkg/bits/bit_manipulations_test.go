package bits

import (
	"testing"
)

// Test inserting one section of a 32 bit number into another.
func TestInsertion(t *testing.T) {
	var N int32 = 1 << 10
	var M int32 = 0b10011
	var i uint = 2
	var j uint = 6

	var expected int32 = 0b10001001100
	output := Insert(N, M, i, j)

	if output != expected {
		t.Error(output)
	}
}

// Test converting a float to a binary string.
func TestBinaryString(t *testing.T) {
	var input float32 = 0.72
	expected := ".1011100001010001111011"

	output, err := BinaryString(input)
	if err != nil {
		t.Error(err)
	}

	if output != expected {
		t.Error(output)
	}

	output, err = BinaryString(1.0)
	if err == nil {
		t.Error(output)
	}

	output, err = BinaryString(-0.1)
	if err == nil {
		t.Error(output)
	}
}

// Test the FlipBitToWin function.
func TestFlipBitToWin(t *testing.T) {
	res := FlipBitToWin(0)
	if res != 1 {
		t.Error(res)
	}

	res = FlipBitToWin(1775)
	if res != 8 {
		t.Error(res)
	}

	res = FlipBitToWin(-1)
	if res != 32 {
		t.Error(res)
	}
}

// Test getting the next and previous numbers with the same number of 1s.
func TestNextAndPrev(t *testing.T) {
	var x int32 = 0b1101110011
	var expectedNext int32 = 0b1101110101
	var expectedPrev int32 = 0b1101101110

	next, err := GetNext(x)
	if err != nil {
		t.Error(err)
	}
	if next != expectedNext {
		t.Error(next)
	}

	prev, err := GetPrev(x)
	if err != nil {
		t.Error(err)
	}
	if prev != expectedPrev {
		t.Error(prev)
	}

	prev, err = GetPrev(0)
	if err == nil {
		t.Error(prev)
	}
}

// Test finding the number of bit flips required to convert one integer into
// another.
func TestFlipsRequired(t *testing.T) {
	numFlips := FlipsRequired(29, 15)
	if numFlips != 2 {
		t.Error(numFlips)
	}
}

// Test the pairwise swap.
func TestSwapPairs(t *testing.T) {
	var input int32 = 0b101110100111
	var expectedOutput int32 = 0b011101011011

	output := SwapPairs(input)
	if output != expectedOutput {
		t.Error(output)
	}
}

// Test drawing some lines.
func TestDrawLine(t *testing.T) {
	screen := make([]byte, 32)

	err := DrawLine(screen, 16, 6, 13, 11)
	if err != nil {
		t.Error(err)
	}

	expectedScreen := make([]byte, 32)
	expectedScreen[22] = 0x03
	expectedScreen[23] = 0xfc

	if !compareScreens(screen, expectedScreen) {
		t.Error(screen)
	}

	err = DrawLine(screen, 16, 11, 14, 6)
	if err != nil {
		t.Error(err)
	}

	expectedScreen[13] = 0x1e

	if !compareScreens(screen, expectedScreen) {
		t.Error(screen)
	}
}

// Return true if two screens are equal in value, false otherwise.
func compareScreens(actual, expected []byte) bool {
	if len(actual) != len(expected) {
		return false
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			return false
		}
	}

	return true
}
