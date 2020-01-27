package ctci

import (
	"fmt"
	"sort"
	"strings"
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

// Test generating combinations of parentheses.
func TestParens(t *testing.T) {
	parens := Parens(3)
	expected := []string{
		"((()))",
		"(()())",
		"()(())",
		"()()()",
	}

	if !compareUnordered(parens, expected) {
		t.Error(parens)
	}
}

// Test the PaintFill function.
func TestPaintFill(t *testing.T) {
	screen := makeScreen()
	expected := makeScreen()

	drawRectangle(screen, Color{Red: 255, Green: 255, Blue: 255})
	drawRectangle(expected, Color{Red: 255, Green: 255, Blue: 255})

	fillManually(expected, Color{Red: 255, Green: 0, Blue: 0})
	PaintFill(
		screen,
		GridCoords{Row: 15, Col: 15},
		Color{Red: 255, Green: 0, Blue: 0},
	)

	if !compareColorScreens(screen, expected) {
		t.Error(screen)
	}
}

func makeScreen() [][]Color {
	screen := make([][]Color, 32)
	for i := range screen {
		screen[i] = make([]Color, 32)
	}

	return screen
}

func drawRectangle(screen [][]Color, color Color) {
	for i := 2; i < 31; i++ {
		screen[i][2] = color
		screen[i][30] = color
		screen[2][i] = color
		screen[30][i] = color
	}
}

func fillManually(screen [][]Color, color Color) {
	for row := 3; row < 30; row++ {
		for col := 3; col < 30; col++ {
			screen[row][col] = color
		}
	}
}

func compareColorScreens(actual, expected [][]Color) bool {
	for row := range actual {
		for col := range expected {
			if actual[row][col] != expected[row][col] {
				return false
			}
		}
	}

	return true
}

func displayScreen(screen [][]Color) {
	var builder strings.Builder

	for row := range screen {
		for col := range screen[row] {
			color := screen[row][col]
			if color.Red == 0 {
				builder.WriteByte(' ')
			} else if color.Green == 0 {
				builder.WriteByte('R')
			} else {
				builder.WriteByte('X')
			}
		}
		builder.WriteByte('\n')
	}

	fmt.Println(builder.String())
}

// Test finding the number of ways of making a value with coins.
func TestCoins(t *testing.T) {
	numCombos := NumCoinCombos(25)
	if numCombos != 13 {
		t.Error(numCombos)
	}
}

// Test finding the number of ways to place 8 Queens
func TestEightQueens(t *testing.T) {
	queenPositions := PlaceEightQueens()
	if len(queenPositions) != 92 {
		t.Error(queenPositions)
	}
}

// Test finding the tallest possible stack from a set of boxes.
func TestTallestStack(t *testing.T) {
	boxes := []Box{
		Box{height: 7, width: 8, depth: 2},
		Box{height: 8, width: 6, depth: 10},
		Box{height: 10, width: 1, depth: 1},
		Box{height: 7, width: 7, depth: 7},
		Box{height: 3, width: 1, depth: 3},
		Box{height: 15, width: 12, depth: 13},
	}

	res := TallestStack(boxes)
	if res != 26 {
		t.Error(res)
	}
}

// Test counting the number of possible boolean evaluation orders.
func TestBooleanEvaluation(t *testing.T) {
	//numWays := BooleanEvaluation("1^0|0|1", false)
	//if numWays != 2 {
	//	t.Error(numWays)
	//}

	numWays := BooleanEvaluation("0&0&0&1^1|0", true)
	if numWays != 10 {
		t.Error(numWays)
	}
}
