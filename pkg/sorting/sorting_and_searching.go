package sorting

import "sort"

func SortedMerge(arrA, arrB []int) []int {
	outArr := make([]int, len(arrA)+len(arrB))
	outIx := len(outArr) - 1
	aIx := len(arrA) - 1
	bIx := len(arrB) - 1

	for bIx >= 0 {
		if aIx >= 0 && arrA[aIx] > arrB[bIx] {
			outArr[outIx] = arrA[aIx]
			aIx--
		} else {
			outArr[outIx] = arrB[bIx]
			bIx--
		}
		outIx--
	}

	if aIx >= 0 {
		copy(outArr[:aIx+1], arrA[:aIx+1])
	}

	return outArr
}

type sortByAnagram struct {
	words       []string
	sortedWords map[string]string
}

func GroupAnagrams(words []string) {
	sort.Sort(newSortByAnagrams(words))
}

func newSortByAnagrams(words []string) *sortByAnagram {
	sortedWords := make(map[string]string)
	for i := range words {
		sortedWords[words[i]] = sortedString(words[i])
	}

	return &sortByAnagram{words: words, sortedWords: sortedWords}
}

func sortedString(str string) string {
	chars := []rune(str)
	sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
	return string(chars)
}

func (words sortByAnagram) Len() int {
	return len(words.words)
}

func (words sortByAnagram) Less(i, j int) bool {
	return words.sortedWords[words.words[i]] < words.sortedWords[words.words[j]]
}

func (words sortByAnagram) Swap(i, j int) {
	words.words[i], words.words[j] = words.words[j], words.words[i]
}
