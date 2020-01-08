package ctci

// Calculate the number of ways of navigating a fixed number of steps in
// jumps of 1, 2 and 3 only.
func NumStepWays(numSteps int) int {
	if numSteps < 1 {
		panic("Invalid number of steps.")
	}

	memo := make([]int, numSteps+1)
	memo[1] = 1
	memo[2] = 2
	memo[3] = 4

	return numStepWays(numSteps, memo)
}

func numStepWays(numSteps int, memo []int) int {
	if memo[numSteps] == 0 {
		memo[numSteps] = numStepWays(numSteps-1, memo) +
			numStepWays(numSteps-2, memo) +
			numStepWays(numSteps-3, memo)
	}

	return memo[numSteps]
}

// Calculate the number of step ways iteratively instead of recursively.
func NumStepWaysIter(numSteps int) int {
	if numSteps < 1 {
		panic("Invalid number of steps")
	}

	if numSteps < 3 {
		switch numSteps {
		case 1:
			return 1

		case 2:
			return 2
		}
	}

	x, y, z := 1, 2, 4
	for i := 3; i < numSteps; i++ {
		x, y, z = y, z, x+y+z
	}

	return z
}
