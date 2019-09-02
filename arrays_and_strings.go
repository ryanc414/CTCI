package ctci

// Determine if a string has all unique chars.
func IsUnique(input string) bool {
	// Assume basic ASCII char set: 128 characters.
	var chars_seen [128]bool

	for _, c := range input {
		if chars_seen[int(c)] {
			return false
		}
		chars_seen[int(c)] = true
	}

	return true
}
