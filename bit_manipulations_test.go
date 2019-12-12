package ctci

import "testing"

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

// Test the SwapBits function.
func TestSwapBits(t *testing.T) {
	var x int32 = 0b1101110011
	var expectedSmaller int32 = 0b1101101011
	var expectedLarger int32 = 0b1101110101

	smaller, larger, err := SwapBits(x)
	if err != nil {
		t.Error(err)
	}

	if smaller != expectedSmaller {
		t.Error(smaller)
	}

	if larger != expectedLarger {
		t.Error(larger)
	}

	smaller, larger, err = SwapBits(0)
	if err == nil {
		t.Error()
	}

	smaller, larger, err = SwapBits(-1)
	if err == nil {
		t.Error()
	}
}
