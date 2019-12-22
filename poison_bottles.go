package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	math_rand "math/rand"
    "errors"
)

const NUM_BOTTLES = 1000

func main() {
    bottles := make([]bool, NUM_BOTTLES)
    seedRng()
    bottles[math_rand.Intn(NUM_BOTTLES)] = true

    index, numDays, err := FindPoisonBottle(bottles, 10, 0)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Found poison bottle at index %v in %v days\n", index, numDays)
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

// Find the poisoned bottle using binary search.
func FindPoisonBottle(bottles []bool,
                      numStrips,
                      numDays int) (int, int, error) {
    if len(bottles) == 1 {
        return 0, numDays, nil
    }

    if numStrips < 1 {
        return -1, numDays, errors.New("Ran out of strips")
    }

    halfway := len(bottles) / 2
    leftResult := TestPoison(bottles[:halfway])

    if leftResult {
        fmt.Println("Check left half")
        return FindPoisonBottle(bottles[:halfway], numStrips - 1, numDays + 7)
    } else {
        fmt.Println("Check right half")
        res, newNumDays, err := FindPoisonBottle(
            bottles[halfway:],
            numStrips - 1,
            numDays + 7,
        )
        return res + halfway, newNumDays, err
    }
}

// Test a single bottle for poison.
func TestPoison(bottles []bool) bool {
    for i := range bottles {
        if bottles[i] {
            return true
        }
    }

    return false
}
