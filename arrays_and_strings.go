package ctci

const AsciiLetterCount = 128

// Determine if a string has all unique chars.
func IsUnique(input string) bool {
	// Assume basic ASCII char set: 128 characters.
	var chars_seen [AsciiLetterCount]bool

	for _, c := range input {
		if chars_seen[int(c)] {
			return false
		}
		chars_seen[int(c)] = true
	}

	return true
}

// Check if one string is a permutation of another.
func IsPermutation(strA string, strB string) bool {
	return charMapOf(strA) == charMapOf(strB)
}

func charMapOf(str string) [AsciiLetterCount]int {
	var charMap [AsciiLetterCount]int

	for _, c := range str {
		charMap[int(c)] += 1
	}

	return charMap
}
