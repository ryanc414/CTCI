package ctci

import "testing"
import "bytes"

// Test the IsUnique function.
func TestIsUnique(t *testing.T) {
	if !IsUnique("abcdefgh") {
		t.Error()
	}
	if IsUnique("abcdefga") {
		t.Error()
	}
	if !IsUnique("") {
		t.Error()
	}
}

// Test the IsPermutation function.
func TestIsPermutation(t *testing.T) {
	if !IsPermutation("abcdefgh", "hgfedcba") {
		t.Error()
	}
	if IsPermutation("abcdefgh", "abcdefga") {
		t.Error()
	}
	if !IsPermutation("", "") {
		t.Error()
	}
}

// Test the Urlify function.
func TestUrlify(t *testing.T) {
	input := "Mr John Smith"
	expected := []byte("Mr%20John%20Smith")

	str := make([]byte, len(input), len(input)+4)
	copy(str, input)

	newSlice := Urlify(str)

	if !bytes.Equal(newSlice, expected) {
		t.Errorf("Expected %v, got %v", expected, str)
	}
}

// Test the Palindrome Permutations function.
func TestIsPalindromePerm(t *testing.T) {
	if !IsPalindromePerm("Tact Coa") {
		t.Error()
	}
	if IsPalindromePerm("Tact Coat") {
		t.Error()
	}
	if !IsPalindromePerm("") {
		t.Error()
	}
}

// Test the IsOneAway function.
func TestIsOneAway(t *testing.T) {
	if !IsOneAway("pale", "ple") {
		t.Error()
	}
	if !IsOneAway("pales", "pale") {
		t.Error()
	}
	if !IsOneAway("pale", "bale") {
		t.Error()
	}
	if IsOneAway("pale", "bake") {
		t.Error()
	}
	if !IsOneAway("", "") {
		t.Error()
	}
}

// Test the CompressStr function.
func TestCompressStr(t *testing.T) {
	if CompressStr("aabcccccaaa") != "a2b1c5a3" {
		t.Error()
	}
	if CompressStr("abcdefgh") != "abcdefgh" {
		t.Error()
	}
	if CompressStr("") != "" {
		t.Error()
	}
}

// Test the RotateMatrix function.
func TestRotateMatrix(t *testing.T) {
	matrix := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}

	expected := [][]int{
		[]int{13, 9, 5, 1},
		[]int{14, 10, 6, 2},
		[]int{15, 11, 7, 3},
		[]int{16, 12, 8, 4},
	}

	RotateMatrix(matrix)
	t.Log(matrix)

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != expected[i][j] {
				t.Error()
			}
		}
	}

	var emptyMatrix [][]int
	RotateMatrix(emptyMatrix)
	if len(emptyMatrix) != 0 {
		t.Error()
	}
}
