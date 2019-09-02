package ctci

import "testing"

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
