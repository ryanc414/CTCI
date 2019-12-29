package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	math_rand "math/rand"
)

const NUM_BOTTLES = 1000

func main() {
	bottles := make([]bool, NUM_BOTTLES)
	seedRng()
	actualIndex := math_rand.Intn(NUM_BOTTLES)
	bottles[actualIndex] = true

	index := FindPoisonBottle(bottles)
	if index != actualIndex {
		panic(fmt.Sprintf("Got index %v, expected %v", index, actualIndex))
	}

	fmt.Printf("Found poison bottle at index %v\n", index)
}

// Seed the RNG so that different results are produced each time.
func seedRng() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("Cannot send RNG")
	}

	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

type TestStrip struct {
	result bool
}

// Find the poisoned bottle using binary search.
func FindPoisonBottle(bottles []bool) int {
	results := runTests(bottles)
	return getBottleIndex(results)
}

// Test each strip for poison.
func runTests(bottles []bool) []bool {
	results := make([]bool, 10)

	for i := range bottles {
		addDropsForBottle(results, bottles, i)
	}

	return results
}

// Add drops from a single bottle to the test strips.
func addDropsForBottle(results []bool,
	bottles []bool,
	bottleIndex int) {
	stripIndex := 0
	indexCopy := bottleIndex

	for indexCopy > 0 {
		if indexCopy&1 == 1 {
			results[stripIndex] = (results[stripIndex] || bottles[bottleIndex])
		}
		indexCopy >>= 1
		stripIndex++
	}
}

// Get the bottle number from a the results.
func getBottleIndex(results []bool) int {
	mask := 1
	bottleIndex := 0

	for i := range results {
		if results[i] {
			bottleIndex += mask
		}
		mask <<= 1
	}

	return bottleIndex
}
