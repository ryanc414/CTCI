package ctci

import (
	"sort"
	"testing"
)

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

// Test finding all subsets of a set.
func TestPowerSet(t *testing.T) {
	// First test the empty set.
	subsets := PowerSet("")
	if subsets != nil {
		t.Error(subsets)
	}

	// Now test a set of four values.
	set := "1234"
	expectedSubsets := []string{
		"1234",
		"234",
		"134",
		"124",
		"123",
		"12",
		"13",
		"14",
		"23",
		"24",
		"34",
		"1",
		"2",
		"3",
		"4",
	}

	subsets = PowerSet(set)
	if !compareSubsets(subsets, expectedSubsets) {
		t.Error(subsets)
	}
}

// Compare a slice of subsets with what is expected.
func compareSubsets(subsets, expectedSubsets []string) bool {
	if len(subsets) != len(expectedSubsets) {
		return false
	}

	// Sort both slices before comparing, since the order is not important.
	sort.Strings(subsets)
	sort.Strings(expectedSubsets)

	for i := range subsets {
		if subsets[i] != expectedSubsets[i] {
			return false
		}
	}

	return true
}

// Test solving the towers of Hanoi.
func TestTowersOfHanoi(t *testing.T) {
	numPieces := 5
	tower := InitTowersOfHanoi(numPieces)
	tower.Solve()

	// Check that the first two towers are both empty.
	for i := 0; i < 2; i++ {
		if !tower.stacks[i].IsEmpty() {
			t.Error(i)
		}
	}

	// Check that all pieces are in the correct order on the last tower.
	for i := 1; i < numPieces+1; i++ {
		val, err := tower.stacks[2].Pop()
		if err != nil {
			t.Error(err)
			break
		}

		if val.(int) != i {
			t.Error(val)
		}
	}
}

// Test calculating permutations of a string with unique chars.
func TestPermutations(t *testing.T) {
	perms := Permutations("abc")
	expectedPerms := []string{
		"abc",
		"bac",
		"bca",
		"acb",
		"cab",
		"cba",
	}

	if len(perms) != len(expectedPerms) {
		t.Error(perms)
	} else {
		for i := range perms {
			if perms[i] != expectedPerms[i] {
				t.Error(perms)
				break
			}
		}
	}
}

// Test calculating permutations of a string with non-unique chars. Each
// permutation should be unique.
func TestUniquePermutations(t *testing.T) {
	perms := UniquePermutations("aabc")
	expectedPerms := []string{
		"aabc",
		"aacb",
		"abac",
		"abca",
		"acab",
		"acba",
		"baac",
		"baca",
		"bcaa",
		"caab",
		"caba",
		"cbaa",
	}

	if !compareUnordered(perms, expectedPerms) {
		t.Error(perms)
	}
}

// Compare two slices of strings without caring about order.
func compareUnordered(actual, expected []string) bool {
	if len(actual) != len(expected) {
		return false
	}

	sort.Strings(actual)
	sort.Strings(expected)

	for i := range actual {
		if actual[i] != expected[i] {
			return false
		}
	}

	return true
}
