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
