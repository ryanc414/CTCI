package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	math_rand "math/rand"
)

const NUM_TRIALS = 1000000

func main() {
	seedRng()

	results := make([]int, NUM_TRIALS)
	for i := 0; i < NUM_TRIALS; i++ {
		results[i] = simulateChildren()
	}

	avgNumBoys := findAvg(results)
	fmt.Printf(
		"On average, each couple had %v boys for every 1 girl.\n",
		avgNumBoys,
	)
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

// Simulate how many boys are had for one girl.
func simulateChildren() int {
	numBoys := 0
	for math_rand.Float64() < 0.5 {
		numBoys++
	}

	return numBoys
}

// Find the average of a slice of ints.
func findAvg(values []int) float64 {
	var sum float64 = 0.0

	for i := range values {
		sum += float64(values[i])
	}

	return sum / float64(len(values))
}
