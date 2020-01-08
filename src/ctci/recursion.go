package ctci

import "errors"

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

// The robot's grid is made up of spaces it can and cannot visit.
type RobotGrid [][]bool

type RobotPath []GridDirection

// The robot can only go right or down.
var RobotDirections = []GridDirection{Right, Down}

// Return a new grid for a robot of given size. Initially all squares are set
// to true.
func InitRobotGrid(size int) RobotGrid {
	grid := make(RobotGrid, size)
	for i := range grid {
		grid[i] = make([]bool, size)
		for j := range grid[i] {
			grid[i][j] = true
		}
	}

	return grid
}

// Finds the first valid path for the robot from top left to bottom right.
func (grid RobotGrid) FindPath() (RobotPath, error) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil, errors.New("Empty grid")
	}

	if !grid[0][0] {
		return nil, errors.New("Starting space is invalid.")
	}

	return grid.findPathRecur(
		GridCoords{Row: 0, Col: 0},
		make(RobotPath, 0, len(grid)),
	)
}

// Recursive implementation.
func (grid RobotGrid) findPathRecur(
	currPos GridCoords, currPath RobotPath,
) (RobotPath, error) {
	// Base case: if we are at the end, return the current path.
	if currPos.Row == len(grid)-1 && currPos.Col == len(grid[0])-1 {
		return currPath, nil
	}

	// Try branching out right or down if possible.
	for i := range RobotDirections {
		newPos := currPos.MoveDirection(RobotDirections[i])
		if grid.validCoords(newPos) && grid[newPos.Row][newPos.Col] {
			newPath := append(currPath, RobotDirections[i])
			fullPath, err := grid.findPathRecur(newPos, newPath)
			if err == nil {
				return fullPath, nil
			}
		}
	}

	// No valid path found, so return error.
	return nil, errors.New("No valid path found.")
}

func (grid RobotGrid) validCoords(coords GridCoords) bool {
	if coords.Row < 0 || coords.Row >= len(grid) {
		return false
	}

	if coords.Col < 0 || coords.Col >= len(grid) {
		return false
	}

	return true
}
