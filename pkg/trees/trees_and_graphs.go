package trees

import (
	"container/list"
	"errors"
	"math/rand"

	"github.com/ryanc414/ctci/pkg/stacks"
)

type GraphNode struct {
	name     string
	adjacent []*GraphNode
	visited  bool
}

type Graph struct {
	nodes []GraphNode
}

func (graph Graph) RouteExists(nodeS, nodeE *GraphNode) bool {
	queue := stacks.NewBasicQueue()
	for i := range graph.nodes {
		graph.nodes[i].visited = false
	}

	nodeS.visited = true
	queue.Add(nodeS)

	for !queue.IsEmpty() {
		nextInQueue, err := queue.Remove()
		currNode := nextInQueue.(*GraphNode)
		if err != nil {
			panic(err)
		}

		if currNode == nodeE {
			return true
		}
		for i := range currNode.adjacent {
			adjNode := currNode.adjacent[i]
			if !adjNode.visited {
				adjNode.visited = true
				queue.Add(adjNode)
			}
		}
	}

	return false
}

type PathNode struct {
	person       *GraphNode
	previousNode *PathNode
}

type BFSData struct {
	toVisit []*PathNode
	visited map[string]*PathNode
}

// Find the shortest route between two nodes.
func (graph Graph) FindShortestPath(nodeA, nodeB *GraphNode) ([]*GraphNode, error) {
	sourceData := newBFSData(nodeA)
	destData := newBFSData(nodeB)

	for !sourceData.isFinished() && !destData.isFinished() {
		collision := searchLevel(graph, sourceData, destData)
		if collision != nil {
			return mergePaths(sourceData, destData, collision), nil
		}

		collision = searchLevel(graph, destData, sourceData)
		if collision != nil {
			return mergePaths(sourceData, destData, collision), nil
		}
	}

	return nil, errors.New("No path found")
}

func searchLevel(graph Graph, primary, secondary *BFSData) *GraphNode {
	var nextToVisit []*PathNode

	for i := range primary.toVisit {
		pathNode := primary.toVisit[i]
		personName := pathNode.person.name

		_, visited := secondary.visited[personName]
		if visited {
			return pathNode.person
		}

		person := pathNode.person
		friends := person.adjacent

		for j := range friends {
			friend := friends[j]
			_, visited = primary.visited[friend.name]
			if !visited {
				next := &PathNode{person: friend, previousNode: pathNode}
				primary.visited[friend.name] = next
				nextToVisit = append(nextToVisit, next)
			}
		}
	}

	primary.toVisit = nextToVisit
	return nil
}

func mergePaths(bfs1, bfs2 *BFSData, connection *GraphNode) []*GraphNode {
	end1 := bfs1.visited[connection.name]
	end2 := bfs2.visited[connection.name]

	pathOne := end1.collapse(false)
	pathTwo := end2.collapse(true)[1:]

	pathOne = append(pathOne, pathTwo...)
	return pathOne
}

func (pathNode *PathNode) collapse(startsWithRoot bool) []*GraphNode {
	var path []*GraphNode
	node := pathNode
	for node != nil {
		if startsWithRoot {
			path = append(path, node.person)
		} else {
			path = append([]*GraphNode{node.person}, path...)
		}
		node = node.previousNode
	}
	return path
}

func (data *BFSData) isFinished() bool {
	return len(data.toVisit) == 0
}

func newBFSData(node *GraphNode) *BFSData {
	sourcePath := &PathNode{person: node, previousNode: nil}
	visited := make(map[string]*PathNode)
	visited[node.name] = sourcePath

	return &BFSData{
		toVisit: []*PathNode{sourcePath},
		visited: visited,
	}
}

// A node in a binary search tree.
type BSTNode struct {
	value  int
	left   *BSTNode
	right  *BSTNode
	parent *BSTNode
}

// Generates a binary search tree of minimal height from a sorted array of
// unique integer elements.
func GenerateBST(sortedArr []int) *BSTNode {
	return generateBSTRecur(sortedArr, nil)
}

func generateBSTRecur(sortedArr []int, parent *BSTNode) *BSTNode {
	if len(sortedArr) == 0 {
		return nil
	}

	midpoint := (len(sortedArr) / 2)

	node := &BSTNode{
		value:  sortedArr[midpoint],
		parent: parent,
	}

	node.left = generateBSTRecur(sortedArr[:midpoint], node)
	node.right = generateBSTRecur(sortedArr[midpoint+1:], node)

	return node
}

// Generate linked lists containing nodes at each depth in a binary tree.
func (root *BSTNode) ListOfDepths() []*list.List {
	if root == nil {
		return nil
	}

	leftDepths := root.left.ListOfDepths()
	rightDepths := root.right.ListOfDepths()
	mergedDepths := mergeDepthLists(leftDepths, rightDepths)

	currDepth := list.New()
	currDepth.PushBack(root)

	return append([]*list.List{currDepth}, mergedDepths...)
}

// Merge two depth lists from left and right subtrees.
func mergeDepthLists(left, right []*list.List) []*list.List {
	merged := make([]*list.List, max(len(left), len(right)))

	for i := 0; i < len(left) || i < len(right); i++ {
		merged[i] = list.New()

		if i < len(left) && left[i] != nil {
			merged[i].PushBackList(left[i])
		}

		if i < len(right) && right[i] != nil {
			merged[i].PushBackList(right[i])
		}
	}

	return merged
}

// Return maximum of two integers.
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Check if a binary tree is balanced. Balanced means that the height of two
// subtrees for any given node do not differ by more than one.
func (tree BSTNode) CheckBalanced() bool {
	balanced, _ := checkBalancedRecur(&tree)
	return balanced
}

// Recursive step for checking if a binary tree is balanced. Check that both
// left and right subtrees are balanced and that their respective heights
// differ at most by one. Return if this node is balanced and its height.
func checkBalancedRecur(node *BSTNode) (bool, int) {
	// Base case: an empty node is balanced and has 0 height.
	if node == nil {
		return true, 0
	}

	// Recurse down into the left and right subtrees.
	leftBalanced, leftHeight := checkBalancedRecur(node.left)
	rightBalanced, rightHeight := checkBalancedRecur(node.right)

	// Now check if we are balanced and calculate the new height.
	currHeight := max(leftHeight, rightHeight) + 1
	heightDiff := abs(leftHeight - rightHeight)
	balanced := leftBalanced && rightBalanced && heightDiff <= 1

	return balanced, currHeight
}

// Validate if a binary tree is also a binary search tree. A BST node is valid
// if its value is greater than any value in the left sub-tree and greater
// than any value in the right sub-tree, and both left and right subtrees are
// valid BSTs.
func (tree BSTNode) ValidateBST() bool {
	valid, _, _ := validateBSTRecur(&tree)
	return valid
}

// Recursive step. For each node, as well as returning whether it is a valid
// BST, we also return the min and max values contained in the subtree to be
// able to validate the parent node. Note that the min and max values are only
// correct for a valid BST - for an invalid BST we short-circuit early and
// don't bother checking the rest of the tree.
func validateBSTRecur(node *BSTNode) (bool, int, int) {
	if node.left == nil && node.right == nil {
		// A single node with no children is valid.
		return true, node.value, node.value
	} else if node.left == nil {
		// Recurse down into the right subtree.
		rightIsValid, rightMin, rightMax := validateBSTRecur(node.right)
		isValid := rightIsValid && rightMin >= node.value
		return isValid, node.value, rightMax
	} else if node.right == nil {
		// Recurse down into the left subtree.
		leftIsValid, leftMin, leftMax := validateBSTRecur(node.left)
		isValid := leftIsValid && leftMax <= node.value
		return isValid, leftMin, node.value
	} else {
		// Recurse down into the left subtree first.
		leftIsValid, leftMin, leftMax := validateBSTRecur(node.left)

		// Short-circuit if the left side of the tree is not valid, to avoid
		// unnecessarily checking the right side.
		if !leftIsValid || leftMax > node.value {
			return false, leftMin, node.value
		}

		// Now check the right subtree.
		rightIsValid, rightMin, rightMax := validateBSTRecur(node.right)
		isValid := rightIsValid && rightMin > node.value
		return isValid, leftMin, rightMax
	}
}

// Return the next in-order successor to a node in a BST. If there is no
// successor (node is the last in the tree) then nil is returned instead.
func (node *BSTNode) Successor() *BSTNode {
	if node.right != nil {
		return node.right.minNode()
	} else {
		return node.parentSuccessor()
	}
}

// Return the minimal node in a BST.
func (node *BSTNode) minNode() *BSTNode {
	if node.left != nil {
		return node.left.minNode()
	} else {
		return node
	}
}

// Return the first parent that succeeds the current node, or nil if there is
// no parent successor.
func (node *BSTNode) parentSuccessor() *BSTNode {
	if node.parent == nil || node.parent.value > node.value {
		return node.parent
	} else {
		return node.parent.parentSuccessor()
	}
}

type NodeState int

const (
	BLANK = iota
	PARTIAL
	COMPLETE
)

type ProjectsGraphNode struct {
	name     string
	adjacent []*ProjectsGraphNode
	state    NodeState
}

type ProjectsGraph struct {
	nodes    []ProjectsGraphNode
	nodesMap map[string]*ProjectsGraphNode
}

// Find a valid build order for projects with dependencies.
func FindBuildOrder(projects []string,
	dependencies [][]string) ([]string, error) {
	graph := buildDepGraph(projects, dependencies)
	var buildOrder list.List

	for i := range graph.nodes {
		order, err := nodeBuildOrder(&graph.nodes[i])
		if err != nil {
			return nil, err
		}
		buildOrder.PushFrontList(&order)
	}

	return listToSlice(buildOrder), nil
}

// Build a graph of dependencies from lists of projects and pairs of
// dependencies.
func buildDepGraph(projects []string, dependencies [][]string) ProjectsGraph {
	depGraph := ProjectsGraph{
		nodes:    make([]ProjectsGraphNode, len(projects)),
		nodesMap: make(map[string]*ProjectsGraphNode),
	}

	for i := range projects {
		depGraph.nodes[i].name = projects[i]
		depGraph.nodesMap[projects[i]] = &depGraph.nodes[i]
	}

	for i := range dependencies {
		depGraph.nodesMap[dependencies[i][0]].adjacent = append(
			depGraph.nodesMap[dependencies[i][0]].adjacent,
			depGraph.nodesMap[dependencies[i][1]],
		)
	}

	return depGraph
}

func nodeBuildOrder(node *ProjectsGraphNode) (list.List, error) {
	var depList list.List
	if node.state == PARTIAL {
		return depList, errors.New("Circular dependency")
	} else if node.state == COMPLETE {
		return depList, nil
	} else if node.state != BLANK {
		return depList, errors.New("Unexpected node state")
	}

	node.state = PARTIAL

	for i := range node.adjacent {
		order, err := nodeBuildOrder(node.adjacent[i])
		if err != nil {
			return depList, err
		}
		depList.PushFrontList(&order)
	}

	node.state = COMPLETE
	depList.PushFront(node)

	return depList, nil
}

func listToSlice(buildOrder list.List) []string {
	buildSlice := make([]string, buildOrder.Len())
	i := 0
	for el := buildOrder.Front(); el != nil; el = el.Next() {
		buildSlice[i] = el.Value.(*ProjectsGraphNode).name
		i++
	}
	return buildSlice
}

// Node in a binary tree (not necessarily a BST)
type BinTreeNode struct {
	name  string
	left  *BinTreeNode
	right *BinTreeNode
}

// Path directions may be either LEFT or RIGHT.
type Direction int

const (
	LEFT = iota
	RIGHT
)

func FindCommonAncestor(root, nodeA, nodeB *BinTreeNode) (*BinTreeNode, error) {
	pathA, found := findNodePath(root, nodeA)
	if !found {
		return nil, errors.New("NodeA not found in tree")
	}

	pathB, found := findNodePath(root, nodeB)
	if !found {
		return nil, errors.New("NodeB not found in tree")
	}

	i := 0
	ancestor := root
	minPathLen := min(len(pathA), len(pathB))

	for i < minPathLen && pathA[i] == pathB[i] {
		if pathA[i] == LEFT {
			ancestor = ancestor.left
		} else {
			ancestor = ancestor.right
		}
		i++
	}

	return ancestor, nil
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Find the path to a given node in a binary tree.
func findNodePath(root, node *BinTreeNode) ([]Direction, bool) {
	if root == nil {
		return nil, false
	}

	if root == node {
		return nil, true
	}

	leftPath, leftFound := findNodePath(root.left, node)
	if leftFound {
		path := append([]Direction{LEFT}, leftPath...)
		return path, true
	}

	rightPath, rightFound := findNodePath(root.right, node)
	if rightFound {
		path := append([]Direction{RIGHT}, rightPath...)
		return path, true
	}

	return nil, false
}

// Find the possible sequences of values that could have created a given BST.
func FindBSTSequences(root *BSTNode) [][]int {
	if root == nil {
		return nil
	}
	return findBSTSeqsRecur(root, nil, nil, nil)
}

func findBSTSeqsRecur(currNode *BSTNode,
	possibleNext []*BSTNode,
	foundSeqs [][]int,
	currSeq []int) [][]int {

	currSeq = append(currSeq, currNode.value)

	if currNode.left != nil {
		possibleNext = append(possibleNext, currNode.left)
	}
	if currNode.right != nil {
		possibleNext = append(possibleNext, currNode.right)
	}

	// Check if we have reached the end of a sequence.
	if len(possibleNext) == 0 {
		foundSeqs = append(foundSeqs, make([]int, len(currSeq)))
		copy(foundSeqs[len(foundSeqs)-1], currSeq)
	} else {
		for i := range possibleNext {
			newPossNext := getNewPossNext(possibleNext, i)
			foundSeqs = findBSTSeqsRecur(
				possibleNext[i],
				newPossNext,
				foundSeqs,
				currSeq,
			)
		}
	}

	return foundSeqs
}

func getNewPossNext(currPossNext []*BSTNode, i int) []*BSTNode {
	newPossNext := make([]*BSTNode, len(currPossNext)-1)
	copy(newPossNext, currPossNext[:i])
	copy(newPossNext[i:], currPossNext[i+1:])

	return newPossNext
}

// Check if T2 is a subtree of T1.
func CheckSubtree(T1, T2 *BinTreeNode) bool {
	if T1 == T2 {
		return true
	}

	if T1 == nil || T2 == nil {
		return false
	}

	if equalTree(T1, T2) {
		return true
	}

	return CheckSubtree(T1.left, T2) || CheckSubtree(T1.right, T2)
}

// Check if two trees are equal.
func equalTree(T1, T2 *BinTreeNode) bool {
	if T1 == T2 {
		return true
	}

	if T1 == nil || T2 == nil {
		return false
	}

	return T1.name == T2.name &&
		equalTree(T1.left, T2.left) &&
		equalTree(T1.right, T2.right)
}

// A counted BST node keeps track of the total number of nodes in its subtree,
// including itself.
type CountedBSTNode struct {
	value int
	count int
	left  *CountedBSTNode
	right *CountedBSTNode
}

// Initialise the root node of a counted BST.
func InitCountedBST(rootValue int) *CountedBSTNode {
	return &CountedBSTNode{
		value: rootValue,
		count: 1,
		left:  nil,
		right: nil,
	}
}

// Get the node at the specified index in the in-order traversal.
func (root *CountedBSTNode) GetNodeAtIndex(index int) (*CountedBSTNode, error) {
	if root == nil {
		return nil, errors.New("Empty tree")
	}

	if index < 0 || index >= root.count {
		return nil, errors.New("Invalid index")
	}

	var leftCount int
	if root.left == nil {
		leftCount = 0
	} else {
		leftCount = root.left.count
	}

	if index < leftCount {
		return root.left.GetNodeAtIndex(index)
	} else if index == leftCount {
		return root, nil
	} else {
		return root.right.GetNodeAtIndex(index - leftCount - 1)
	}
}

// Get a random node from a counted BST.
func (root *CountedBSTNode) GetRandomNode() *CountedBSTNode {
	randIndex := rand.Intn(root.count)
	node, err := root.GetNodeAtIndex(randIndex)

	if err != nil {
		panic(err)
	}

	return node
}

// Insert a new value into a counted BST, making sure to update the count
// of all nodes as required.
func (root *CountedBSTNode) Insert(value int) {
	if root == nil {
		panic("Cannot insert value into empty tree")
	}

	root.count++

	if value < root.value {
		if root.left == nil {
			root.left = InitCountedBST(value)
		} else {
			root.left.Insert(value)
		}
	} else {
		if root.right == nil {
			root.right = InitCountedBST(value)
		} else {
			root.right.Insert(value)
		}
	}
}

type IntBinTree struct {
	value int
	left  *IntBinTree
	right *IntBinTree
}

func PathsWithSum(node *IntBinTree, targetSum int) int {
	pathCount := make(map[int]int)
	return countPathsWithSumRecur(node, targetSum, 0, pathCount)
}

func countPathsWithSumRecur(node *IntBinTree,
	targetSum,
	runningSum int,
	pathCount map[int]int) int {
	if node == nil {
		return 0
	}

	runningSum += node.value
	sum := runningSum - targetSum
	totalPaths := pathCount[sum]

	// If runningSum == targetSum, there is an additional path starting at root.
	if runningSum == targetSum {
		totalPaths++
	}

	pathCount[runningSum]++
	totalPaths += countPathsWithSumRecur(
		node.left,
		targetSum,
		runningSum,
		pathCount,
	)
	totalPaths += countPathsWithSumRecur(
		node.right,
		targetSum,
		runningSum,
		pathCount,
	)

	if pathCount[runningSum] == 1 {
		delete(pathCount, runningSum)
	} else {
		pathCount[runningSum]--
	}

	return totalPaths
}
