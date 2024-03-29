package recursion

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/ryanc414/ctci/pkg/objects"
	"github.com/ryanc414/ctci/pkg/stacks"
)

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

type RobotPath []objects.GridDirection

// The robot can only go right or down.
var RobotDirections = []objects.GridDirection{objects.Right, objects.Down}

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
		objects.GridCoords{Row: 0, Col: 0},
		make(RobotPath, 0, len(grid)),
	)
}

// Recursive implementation.
func (grid RobotGrid) findPathRecur(
	currPos objects.GridCoords, currPath RobotPath,
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

func (grid RobotGrid) validCoords(coords objects.GridCoords) bool {
	if coords.Row < 0 || coords.Row >= len(grid) {
		return false
	}

	if coords.Col < 0 || coords.Col >= len(grid) {
		return false
	}

	return true
}

// Finds a "magic" index in a sorted array where A[i] == i, if one exists.
func FindMagicIndex(sortedArr []int) (int, error) {
	return findMagicIndexRecur(sortedArr, 0)
}

func findMagicIndexRecur(sortedArr []int, offset int) (int, error) {
	// Base case: return error for empty slice.
	if len(sortedArr) == 0 {
		return 0, errors.New("No magic index found.")
	}

	middle := (len(sortedArr) - 1) / 2
	midVal := sortedArr[middle]

	switch {
	case midVal == middle+offset:
		return midVal, nil

	case midVal > middle+offset:
		return findMagicIndexRecur(sortedArr[:middle], offset)

	default:
		return findMagicIndexRecur(sortedArr[middle+1:], offset+middle+1)
	}
}

// Finds all subsets of a set.
func PowerSet(set string) []string {
	cache := make(map[string]bool)

	// Note: could optimize by pre-allocating the slice of strings up front.
	return powerSetRecur(set, cache)
}

func powerSetRecur(set string, cache map[string]bool) []string {
	// Base case: empty set.
	if len(set) == 0 {
		return nil
	}

	// Check if this set has already been found, if so we skip.
	if _, ok := cache[set]; ok {
		return nil
	}

	cache[set] = true

	sets := []string{set}
	for i := range set {
		var builder strings.Builder
		builder.WriteString(set[:i])
		builder.WriteString(set[i+1:])
		sets = append(sets, powerSetRecur(builder.String(), cache)...)
	}

	return sets
}

// Multiply two positive integers recursively, without using the * operator.
func MultiplyRecur(a, b int) int {
	if a < 0 || b < 0 {
		panic("Expect positive integers")
	}

	if a == 0 || b == 0 {
		return 0
	}

	return a + MultiplyRecur(a, b-1)
}

// Represent the Towers of Hanoi by an array of three stacks.
type TowersOfHanoi struct {
	stacks    [3]stacks.Stack
	numPieces int
}

// Initialise the towers of Hanoi with a given number of pieces on the left
// tower.
func InitTowersOfHanoi(numPieces int) TowersOfHanoi {
	var towers TowersOfHanoi

	// Initialise three stacks.
	for i := 0; i < 3; i++ {
		towers.stacks[i] = stacks.NewBasicStack()
	}

	// Place pieces on the left stack.
	for i := numPieces; i > 0; i-- {
		towers.stacks[0].Push(i)
	}

	towers.numPieces = numPieces
	return towers
}

// Solve the towers of hanoi, by moving all pieces from the left tower onto
// the right without placing a bigger piece on top of a smaller one.
func (towers TowersOfHanoi) Solve() {
	//towers.Display()
	towers.solveTowersRecur(0, 1, 2, towers.numPieces)
}

// Recursive implementation.
func (towers TowersOfHanoi) solveTowersRecur(
	tower0, tower1, tower2, numPieces int,
) {
	if numPieces == 1 {
		towers.movePiece(tower0, tower2)
		return
	}

	towers.solveTowersRecur(tower0, tower2, tower1, numPieces-1)
	towers.movePiece(tower0, tower2)
	towers.solveTowersRecur(tower1, tower0, tower2, numPieces-1)
}

// Move a piece from one tower to another.
func (towers TowersOfHanoi) movePiece(fromTower, toTower int) {
	piece, err := towers.stacks[fromTower].Pop()
	if err != nil {
		panic(err)
	}

	// Check that we are placing onto a larger piece, or the base.
	onPiece, err := towers.stacks[toTower].Peek()
	if err == nil && onPiece.(int) < piece.(int) {
		panic("Cannot move bigger piece onto smaller one.")
	}

	towers.stacks[toTower].Push(piece)
	//towers.Display()
}

// Format and print the towers to stdout.
func (towers TowersOfHanoi) Display() {
	maxTowerSize := 0
	for i := range towers.stacks {
		towerSize := len(towers.stacks[i].(*stacks.BasicStack).Data)
		if towerSize > maxTowerSize {
			maxTowerSize = towerSize
		}
	}

	var builder strings.Builder

	divLine := makeDivLine()

	for i := maxTowerSize - 1; i >= 0; i-- {
		for j := range towers.stacks {
			stackData := towers.stacks[j].(*stacks.BasicStack).Data
			if i < len(stackData) {
				builder.WriteString(renderPiece(stackData[i].(int)))
			} else {
				builder.WriteString(renderPiece(0))
			}
		}
		builder.WriteRune('\n')
	}
	builder.WriteString(divLine)
	builder.WriteRune('\n')
	fmt.Println(builder.String())
}

func makeDivLine() string {
	divLine := make([]byte, 6*2*3)
	for i := range divLine {
		divLine[i] = '_'
	}

	return string(divLine)
}

// Return a string representation of a tower piece.
func renderPiece(pieceNum int) string {
	maxSize := 5
	if pieceNum > maxSize {
		panic("Cannot render piece")
	}

	stringSize := 2*maxSize + 1
	pieceString := make([]byte, stringSize)
	for i := range pieceString {
		pieceString[i] = ' '
	}

	for j := 0; j < pieceNum; j++ {
		pieceString[maxSize+j] = 'X'
		pieceString[maxSize-j] = 'X'
	}

	return string(pieceString)
}

func Permutations(input string) []string {
	if len(input) == 0 {
		return nil
	}

	if len(input) == 1 {
		return []string{input}
	}

	head := input[0]
	tail := input[1:]
	tailPerms := Permutations(tail)
	var perms []string

	for i := range tailPerms {
		tailPerm := tailPerms[i]

		for j := 0; j <= len(tailPerm); j++ {
			var builder strings.Builder
			builder.WriteString(tailPerm[:j])
			builder.WriteByte(head)
			builder.WriteString(tailPerm[j:])
			perms = append(perms, builder.String())
		}
	}

	return perms
}

// Find all the unique permutations of a string, that may contain duplicate
// characters.
func UniquePermutations(input string) []string {
	if len(input) == 0 {
		return nil
	}

	charCounts := make(map[byte]int)
	for i := range input {
		charCounts[input[i]]++
	}

	return permutationsRecur(charCounts)
}

func permutationsRecur(charCounts map[byte]int) []string {
	var perms []string

	for k, v := range charCounts {
		if v == 0 {
			continue
		}

		charCounts[k]--
		tailPerms := permutationsRecur(charCounts)
		charCounts[k]++

		if tailPerms == nil {
			return []string{buildBaseString(k, v)}
		}

		for i := range tailPerms {
			perms = append(perms, appendTailString(k, tailPerms[i]))
		}

	}

	return perms
}

func buildBaseString(baseChar byte, numTimes int) string {
	bytes := make([]byte, numTimes)
	for i := range bytes {
		bytes[i] = baseChar
	}

	return string(bytes)
}

func appendTailString(char byte, tailString string) string {
	bytes := make([]byte, len(tailString)+1)
	bytes[0] = char
	for i := range tailString {
		bytes[i+1] = tailString[i]
	}

	return string(bytes)
}

func Parens(numPairs int) []string {
	if numPairs < 0 {
		panic("numPairs must be positive")
	}

	// Base cases.
	switch numPairs {
	case 0:
		return nil

	case 1:
		return []string{"()"}

	case 2:
		return []string{"(())", "()()"}
	}

	var parens []string

	tailParens := Parens(numPairs - 1)
	for i := range tailParens {
		parens = append(
			parens,
			concatStrings("(", tailParens[i], ")"),
			concatStrings("()", tailParens[i]),
		)
	}

	return parens
}

func concatStrings(strs ...string) string {
	var builder strings.Builder
	for i := range strs {
		builder.WriteString(strs[i])
	}
	return builder.String()
}

type Color struct {
	Red   int
	Green int
	Blue  int
}

func PaintFill(screen [][]Color, point objects.GridCoords, color Color) {
	origColour := screen[point.Row][point.Col]
	screen[point.Row][point.Col] = color

	for i := range objects.GridDirections {
		nextPoint := point.MoveDirection(objects.GridDirections[i])
		if validCoords(screen, nextPoint) &&
			screen[nextPoint.Row][nextPoint.Col] == origColour {
			PaintFill(screen, nextPoint, color)
		}
	}
}

func validCoords(screen [][]Color, point objects.GridCoords) bool {
	if point.Row < 0 || point.Row >= len(screen) {
		return false
	}

	if point.Col < 0 || point.Col >= len(screen[0]) {
		return false
	}

	return true
}

// Count the number of ways of making a value with coins of 25, 10, 5 and 1.
func NumCoinCombos(value int) int {
	coinVals := []int{25, 10, 5, 1}
	return numCoinCombosRecur(value, coinVals)
}

func numCoinCombosRecur(value int, coinVals []int) int {
	// Base case
	if value < 5 {
		return 1
	}

	numCombos := 0

	for i := range coinVals {
		if coinVals[i] <= value {
			numCombos += numCoinCombosRecur(value-coinVals[i], coinVals[i:])
		}
	}

	return numCombos
}

func PlaceEightQueens() [][]objects.GridCoords {
	return placeQueensRecur(nil)
}

func placeQueensRecur(placedQueens []objects.GridCoords) [][]objects.GridCoords {
	row := len(placedQueens)

	if row == 8 {
		return [][]objects.GridCoords{placedQueens}
	}

	var places [][]objects.GridCoords
	for col := 0; col < 8; col++ {
		pos := objects.GridCoords{Row: row, Col: col}
		if checkQueens(placedQueens, pos) {
			newPlacedQueens := append(placedQueens, pos)
			places = append(places, placeQueensRecur(newPlacedQueens)...)
		}
	}

	return places
}

func checkQueens(placedQueens []objects.GridCoords, pos objects.GridCoords) bool {
	for i := range placedQueens {
		queen := placedQueens[i]
		if queen.Col == pos.Col {
			return false
		}

		if Abs(queen.Row-pos.Row) == Abs(queen.Col-pos.Col) {
			return false
		}
	}

	return true
}

// Calculate the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

type Box struct {
	width  int
	height int
	depth  int
}

type byHeight []Box

func (boxes byHeight) Len() int {
	return len(boxes)
}

func (boxes byHeight) Less(i, j int) bool {
	return boxes[i].height > boxes[j].height
}

func (boxes byHeight) Swap(i, j int) {
	boxes[i], boxes[j] = boxes[j], boxes[i]
}

func TallestStack(boxes []Box) int {
	sort.Sort(byHeight(boxes))
	cache := make(map[int]int)
	return tallestStackRecur(boxes, -1, cache)
}

func tallestStackRecur(boxes []Box, lastStacked int, cache map[int]int) int {
	cachedVal, exist := cache[lastStacked]
	if exist {
		return cachedVal
	}

	first, end := possibleNextBoxes(boxes, lastStacked)
	maxHeight := 0

	for i := first; i < end; i++ {
		height := boxes[i].height + tallestStackRecur(boxes, i, cache)

		if height > maxHeight {
			maxHeight = height
		}
	}

	cache[lastStacked] = maxHeight
	return maxHeight
}

func possibleNextBoxes(boxes []Box, lastStacked int) (int, int) {
	first := firstPossibleNextBox(boxes, lastStacked)

	for i := first + 1; i < len(boxes); i++ {
		if dominatesBox(boxes[first], boxes[i]) {
			return first, i
		}
	}

	return first, len(boxes)
}

func firstPossibleNextBox(boxes []Box, lastStacked int) int {
	if lastStacked == -1 {
		return 0
	}

	for i := range boxes {
		if dominatesBox(boxes[lastStacked], boxes[i]) {
			return i
		}
	}

	return len(boxes)
}

func dominatesBox(boxA, boxB Box) bool {
	return boxA.width > boxB.width && boxA.height > boxB.height && boxA.depth > boxB.depth
}

// Count the number of ways of parenthesising a boolean expression that gives
// the desired result.
func BooleanEvaluation(expr string, targetVal bool) int {
	if len(expr) == 0 {
		return 0
	}

	if !validBoolExpr(expr) {
		panic("Invalid expression")
	}

	cache := make(map[boolExpr]int)

	// Pre-cache the base cases.
	cache[boolExpr{str: "1", target: true}] = 1
	cache[boolExpr{str: "1", target: false}] = 0
	cache[boolExpr{str: "0", target: true}] = 0
	cache[boolExpr{str: "0", target: false}] = 1

	return boolEvalRecur(boolExpr{str: expr, target: targetVal}, cache)
}

type boolExpr struct {
	str    string
	target bool
}

func boolEvalRecur(expr boolExpr, cache map[boolExpr]int) int {
	// First check the cache.
	cachedVal, exist := cache[expr]
	if exist {
		return cachedVal
	}

	// Iterate through the operators.
	totalNumWays := 0
	for i := 1; i < len(expr.str); i += 2 {
		numWays := boolNumWays(expr, i, cache)
		totalNumWays += numWays
	}

	// Cache the return value for this bool expression.
	cache[expr] = totalNumWays
	return totalNumWays
}

func boolNumWays(expr boolExpr, i int, cache map[boolExpr]int) int {
	switch expr.str[i] {
	case '|':
		if expr.target {
			// Require EITHER left or right sides to evaluate to true.
			numWays := boolEvalRecur(
				boolExpr{str: expr.str[:i], target: true}, cache,
			) * boolEvalRecur(
				boolExpr{str: expr.str[i+1:], target: true}, cache,
			)

			numWays += boolEvalRecur(
				boolExpr{str: expr.str[:i], target: true}, cache,
			) * boolEvalRecur(
				boolExpr{str: expr.str[i+1:], target: false}, cache,
			)

			numWays += boolEvalRecur(
				boolExpr{str: expr.str[:i], target: false}, cache,
			) * boolEvalRecur(
				boolExpr{str: expr.str[i+1:], target: true}, cache,
			)

			return numWays
		} else {
			// Require BOTH left or right sides to evaluate to false.
			return boolEvalRecur(
				boolExpr{str: expr.str[:i], target: false}, cache,
			) * boolEvalRecur(
				boolExpr{str: expr.str[i+1:], target: false}, cache,
			)
		}

	case '&':
		if expr.target {
			// Require BOTH left or right sides to evaluate to true.
			return boolEvalRecur(
				boolExpr{str: expr.str[:i], target: true}, cache,
			) * boolEvalRecur(
				boolExpr{str: expr.str[i+1:], target: true}, cache,
			)
		} else {
			// Require EITHER left or right sides to evaluate to false.
			numWays := boolEvalRecur(
				boolExpr{str: expr.str[:i], target: false}, cache,
			) * boolEvalRecur(
				boolExpr{str: expr.str[i+1:], target: false}, cache,
			)

			numWays += boolEvalRecur(
				boolExpr{str: expr.str[:i], target: false}, cache,
			) * boolEvalRecur(
				boolExpr{str: expr.str[i+1:], target: true}, cache,
			)

			numWays += boolEvalRecur(
				boolExpr{str: expr.str[:i], target: true}, cache,
			) * boolEvalRecur(
				boolExpr{str: expr.str[i+1:], target: false}, cache,
			)
			return numWays
		}

	case '^':
		if expr.target {
			// Require EITHER but not BOTH left or right sides to evaluate
			// to true.
			numWays := boolEvalRecur(
				boolExpr{str: expr.str[:i], target: true}, cache,
			) * boolEvalRecur(
				boolExpr{str: expr.str[i+1:], target: false}, cache,
			)

			numWays += boolEvalRecur(
				boolExpr{str: expr.str[:i], target: false}, cache,
			) * boolEvalRecur(
				boolExpr{str: expr.str[i+1:], target: true}, cache,
			)

			return numWays
		} else {
			// Require BOTH left or right to be true OR BOTH left or right
			// to be false.
			numWays := boolEvalRecur(
				boolExpr{str: expr.str[:i], target: true}, cache,
			) * boolEvalRecur(
				boolExpr{str: expr.str[i+1:], target: true}, cache,
			)

			numWays += boolEvalRecur(
				boolExpr{str: expr.str[:i], target: false}, cache,
			) * boolEvalRecur(
				boolExpr{str: expr.str[i+1:], target: false}, cache,
			)
			return numWays
		}

	default:
		panic("Unexpected token")
	}

}

func totalBoolPossibilities(boolExpr string) int {
	if len(boolExpr) == 1 {
		return 1
	}
	return (len(boolExpr) - 1) / 2
}

func validBoolExpr(boolExpr string) bool {
	if len(boolExpr)%2 != 1 {
		return false
	}

	for i := range boolExpr {
		if i%2 == 0 {
			switch boolExpr[i] {
			case '0', '1':
				continue

			default:
				return false
			}
		} else {
			switch boolExpr[i] {
			case '|', '&', '^':
				continue

			default:
				return false
			}
		}
	}

	return true
}
