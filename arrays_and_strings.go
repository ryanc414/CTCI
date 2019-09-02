package ctci

const AsciiLetterCount = 128
const NumLetters = 26

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

// Return a map of character counts in a string, assuming only standard
// ASCII characters.
func charMapOf(str string) [AsciiLetterCount]int {
	var charMap [AsciiLetterCount]int

	for _, c := range str {
		charMap[int(c)] += 1
	}

	return charMap
}

// Convert spaces to %20 inside an array of bytes in-place.
func Urlify(str []byte) []byte {
	reqdSize := (countChar(str, ' ') * 2) + len(str)
	origLen := len(str)
	retSlice := str[:reqdSize]

	newIx := reqdSize - 1
	for oldIx := origLen - 1; oldIx >= 0; oldIx-- {
		if str[oldIx] == ' ' {
			retSlice[newIx] = '0'
			retSlice[newIx-1] = '2'
			retSlice[newIx-2] = '%'
			newIx -= 3
		} else {
			retSlice[newIx] = str[oldIx]
			newIx--
		}
	}

	return retSlice
}

// Count occurences of a character in a string (given as a byte slice)
func countChar(str []byte, char byte) int {
	count := 0
	for _, c := range str {
		if c == char {
			count++
		}
	}
	return count
}

// Check if a word is a permutation of a palindrome
func IsPalindromePerm(word string) bool {
	letterFreqs := countLetterFreqs(word)
	return maxOneOdd(letterFreqs)
}

// Count the frequency of letters in a word. Casing is ignored (A and a are
// equivalent) and non-letter characters e.g. punctuation is ignored (not
// included in any count)
func countLetterFreqs(word string) []int {
	var letterFreqs [NumLetters]int

	for _, c := range word {
		if (int(c) >= int('a')) && (int(c) <= int('z')) {
			letterFreqs[int(c)-int('a')]++
		} else if (int(c) >= int('A')) && (int(c) <= int('Z')) {
			letterFreqs[int(c)-int('A')]++
		}
		// Other characters are ignored.
	}

	return letterFreqs[:]
}

// Check if a slice contains at most one odd value.
func maxOneOdd(letterFreqs []int) bool {
	oddFound := false

	for _, freq := range letterFreqs {
		if freq%2 == 1 {
			if oddFound {
				return false
			}
			oddFound = true
		}
	}

	return true
}
