package ctci

import "testing"

func TestSortedMerge(t *testing.T) {
	// Start with two easy cases.
	arrA := []int{1, 2, 3}
	arrB := []int{4, 5, 6}

	res := SortedMerge(arrA, arrB)
	expected := []int{1, 2, 3, 4, 5, 6}

	if !equalIntSlices(res, expected) {
		t.Error(res)
	}

	// Now try the other way
	arrA = []int{4, 5, 6}
	arrB = []int{1, 2, 3}
	res = SortedMerge(arrA, arrB)

	if !equalIntSlices(res, expected) {
		t.Error(res)
	}

	// Now try a more complicated example.
	arrA = []int{1, 2, 5, 7, 8, 10, 12, 15}
	arrB = []int{3, 4, 6, 9, 11, 13, 14}
	expected = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	res = SortedMerge(arrA, arrB)

	if !equalIntSlices(res, expected) {
		t.Error(res)
	}
}

func TestGroupAnagrams(t *testing.T) {
	words := []string{
		"abcde",
		"aaaaa",
		"b",
		"bcdea",
		"cde",
		"edcba",
	}

	GroupAnagrams(words)
	expectedGrouped := []string{
		"aaaaa",
		"abcde",
		"bcdea",
		"edcba",
		"b",
		"cde",
	}

	for i := range words {
		if words[i] != expectedGrouped[i] {
			t.Error(words)
			break
		}
	}
}

func equalIntSlices(arrA, arrB []int) bool {
	if len(arrA) != len(arrB) {
		return false
	}

	for i := range arrA {
		if arrA[i] != arrB[i] {
			return false
		}
	}

	return true
}
