package ctci

import "testing"

// Test calculating the number of step ways recursively and iteratively.
func TestNumStepWays(t *testing.T) {
	res := NumStepWays(6)
	if res != 24 {
		t.Error(res)
	}

	res = NumStepWaysIter(6)
	if res != 24 {
		t.Error(res)
	}
}

// Test finding valid paths through a grid for a robot.
func TestRobotPaths(t *testing.T) {
	grid := InitRobotGrid(5)

	// First try with an empty grid (all paths valid). The first path found
	// should be the one that goes right 4 times, then down 4 times.
	expected := RobotPath{
		Right,
		Right,
		Right,
		Right,
		Down,
		Down,
		Down,
		Down,
	}

	path, err := grid.FindPath()
	if err != nil {
		t.Error(err)
	}
	if !compareRobotPaths(path, expected) {
		t.Error(path)
	}

	// Now try a slightly more complex path for our robot.
	grid = InitRobotGrid(7)

	grid[1][1] = false
	grid[1][2] = false
	grid[1][3] = false
	grid[2][0] = false
	grid[2][5] = false
	grid[3][5] = false
	grid[3][6] = false
	grid[4][6] = false
	grid[5][1] = false
	grid[6][5] = false

	expected = RobotPath{
		Right,
		Right,
		Right,
		Right,
		Down,
		Down,
		Down,
		Down,
		Right,
		Down,
		Right,
		Down,
	}

	path, err = grid.FindPath()
	if err != nil {
		t.Error(err)
	}
	if !compareRobotPaths(path, expected) {
		t.Error(path)
	}
}

func compareRobotPaths(actual, expected RobotPath) bool {
	if len(actual) != len(expected) {
		return false
	}

	for i := range actual {
		if actual[i] != expected[i] {
			return false
		}
	}

	return true
}

// Test finding "magic" indices in arrays.
func TestMagicIndex(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4}
	val, err := FindMagicIndex(arr)
	if err != nil {
		t.Error(err)
	}
	if val != 2 {
		t.Error(val)
	}

	arr = []int{-10, -5, 0, 3, 10}
	val, err = FindMagicIndex(arr)
	if err != nil {
		t.Error(err)
	}
	if val != 3 {
		t.Error(val)
	}

	arr = []int{20, 40, 60, 80, 100}
	val, err = FindMagicIndex(arr)
	if err == nil {
		t.Error(val)
	}
}
