package ctci

import "testing"

// Test the IsUnique functin.
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
