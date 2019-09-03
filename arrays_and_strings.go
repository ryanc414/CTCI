package ctci

import (
	"fmt"
	"strings"
)

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

// Check if two strings are at most one edit away. An edit may be inserting,
// removing or swapping a single char.
func IsOneAway(strA string, strB string) bool {
	lenDiff := len(strA) - len(strB)

	switch lenDiff {
	case 0:
		return oneSwapAway(strA, strB)

	case 1:
		return oneInsertAway(strA, strB)

	case -1:
		return oneInsertAway(strB, strA)

	default:
		return false
	}
}

// Check if two strings of the same length are at most one swap away.
func oneSwapAway(strA string, strB string) bool {
	swapFound := false

	for i := range strA {
		if strA[i] != strB[i] {
			if swapFound {
				return false
			}
			swapFound = true
		}
	}

	return true
}

// Check if two strings are at most one insert away.
func oneInsertAway(strA string, strB string) bool {
	insertFound := false
	aIndex := 0

	for bIndex := range strB {
		if strA[aIndex] != strB[bIndex] {
			if insertFound {
				return false
			}
			insertFound = true
			aIndex += 2
		} else {
			aIndex++
		}
	}

	return true
}

// Perform basic string compression using counts of repeated chars.
func CompressStr(input string) string {
	var b strings.Builder
	var currChar rune
	currCount := 0

	for _, c := range input {
		if c == currChar {
			currCount++
		} else {
			if currCount > 0 {
				fmt.Fprintf(&b, "%s%d", string(currChar), currCount)
			}
			currChar = c
			currCount = 1
		}
	}

	if currCount > 0 {
		fmt.Fprintf(&b, "%s%d", string(currChar), currCount)
	}

	compressedStr := b.String()
	if len(compressedStr) > len(input) {
		return input
	} else {
		return compressedStr
	}
}

// Rotates an input matrix in-place by 90 degrees clockwise.
func RotateMatrix(matrix [][]int) {
	size := len(matrix)
	numShells := int(size / 2)

	for i := 0; i < numShells; i++ {
		rotateSubshell(matrix, i)
	}
}

// Rotate a specific subshell within a matrix 90 degrees clockwise.
func rotateSubshell(matrix [][]int, shellIx int) {
	last := len(matrix) - 1 - shellIx

	for i := shellIx; i < last; i++ {
		offset := i - shellIx
		top := matrix[shellIx][i]

		// left -> top
		matrix[shellIx][i] = matrix[last-offset][shellIx]

		// bottom -> left
		matrix[last-offset][shellIx] = matrix[last][last-offset]

		// right -> bottom
		matrix[last][last-offset] = matrix[i][last]

		// top -> right
		matrix[i][last] = top
	}
}

// For any element in an MxN matrix that is 0, set its entire row and column
// to 0.
func ZeroMatrix(matrix [][]int) {
	var zeroRows []int
	var zeroCols []int

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				zeroRows = append(zeroRows, i)
				zeroCols = append(zeroCols, j)
			}
		}
	}

	for _, rowIx := range zeroRows {
		for i := range matrix[rowIx] {
			matrix[rowIx][i] = 0
		}
	}

	for _, colIx := range zeroCols {
		for i := range matrix {
			matrix[i][colIx] = 0
		}
	}
}
