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
