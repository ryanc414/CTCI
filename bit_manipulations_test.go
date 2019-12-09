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
