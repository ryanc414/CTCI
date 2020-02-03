package ctci

import (
	"container/list"
	"testing"
)

func TestRouteExists(t *testing.T) {
	graph := Graph{
		nodes: make([]GraphNode, 6),
	}

	graph.nodes[0].name = "0"
	graph.nodes[0].adjacent = []*GraphNode{
		&graph.nodes[1],
		&graph.nodes[4],
		&graph.nodes[5],
	}

	graph.nodes[1].name = "1"
	graph.nodes[1].adjacent = []*GraphNode{&graph.nodes[3], &graph.nodes[4]}

	graph.nodes[2].name = "2"
	graph.nodes[2].adjacent = []*GraphNode{&graph.nodes[1]}

	graph.nodes[3].name = "3"
	graph.nodes[3].adjacent = []*GraphNode{&graph.nodes[2], &graph.nodes[4]}

	graph.nodes[4].name = "4"
	graph.nodes[4].adjacent = []*GraphNode{}

	graph.nodes[5].name = "5"
	graph.nodes[5].adjacent = []*GraphNode{}

	if !graph.RouteExists(&graph.nodes[0], &graph.nodes[1]) {
		t.Error()
	}

	if !graph.RouteExists(&graph.nodes[0], &graph.nodes[2]) {
		t.Error()
	}

	if !graph.RouteExists(&graph.nodes[3], &graph.nodes[1]) {
		t.Error()
	}

	if graph.RouteExists(&graph.nodes[5], &graph.nodes[0]) {
		t.Error()
	}

	if graph.RouteExists(&graph.nodes[3], &graph.nodes[5]) {
		t.Error()
	}
}

func TestShortestRoute(t *testing.T) {
	graph := Graph{
		nodes: make([]GraphNode, 6),
	}

	graph.nodes[0].name = "A"
	graph.nodes[0].adjacent = []*GraphNode{
		&graph.nodes[1],
		&graph.nodes[2],
		&graph.nodes[4],
	}

	graph.nodes[1].name = "B"
	graph.nodes[1].adjacent = []*GraphNode{
		&graph.nodes[0],
		&graph.nodes[2],
		&graph.nodes[3],
		&graph.nodes[4],
	}

	graph.nodes[2].name = "C"
	graph.nodes[2].adjacent = []*GraphNode{
		&graph.nodes[0],
		&graph.nodes[1],
		&graph.nodes[4],
	}

	graph.nodes[3].name = "D"
	graph.nodes[3].adjacent = []*GraphNode{
		&graph.nodes[1],
		&graph.nodes[5],
	}

	graph.nodes[4].name = "E"
	graph.nodes[4].adjacent = []*GraphNode{
		&graph.nodes[0],
		&graph.nodes[1],
		&graph.nodes[2],
		&graph.nodes[5],
	}

	graph.nodes[5].name = "F"
	graph.nodes[5].adjacent = []*GraphNode{
		&graph.nodes[3],
		&graph.nodes[4],
	}

	path, err := graph.FindShortestPath(&graph.nodes[0], &graph.nodes[1])
	if err != nil || len(path) != 2 {
		t.Errorf("err = %v, path = %v", err, path)
	}

	path, err = graph.FindShortestPath(&graph.nodes[0], &graph.nodes[5])
	if err != nil || len(path) != 3 {
		t.Errorf("err = %v, path = %v", err, path)
	}
}

// Test the GenerateBST function.
func TestGenerateBST(t *testing.T) {
	inputArr := []int{1, 3, 4, 8, 10, 11, 15, 21}
	expectedBST := &BSTNode{
		value: 10,
		left: &BSTNode{
			value: 4,
			left: &BSTNode{
				value: 3,
				left: &BSTNode{
					value: 1,
					left:  nil,
					right: nil,
				},
				right: nil,
			},
			right: &BSTNode{
				value: 8,
				left:  nil,
				right: nil,
			},
		},
		right: &BSTNode{
			value: 15,
			left: &BSTNode{
				value: 11,
				left:  nil,
				right: nil,
			},
			right: &BSTNode{
				value: 21,
				left:  nil,
				right: nil,
			},
		},
	}

	actualBST := GenerateBST(inputArr)
	if !isEqualBST(actualBST, expectedBST) {
		t.Error(actualBST)
	}
}

func isEqualBST(nodeX, nodeY *BSTNode) bool {
	if nodeX == nodeY {
		return true
	}

	if nodeX == nil || nodeY == nil {
		return false
	}

	if nodeX.value != nodeY.value {
		return false
	}

	return isEqualBST(nodeX.left, nodeY.left) &&
		isEqualBST(nodeX.right, nodeY.right)
}

// Test the ListOfDepths function.
func TestListOfDepths(t *testing.T) {
	inputArr := []int{1, 3, 4, 8, 10, 11, 15, 21}
	tree := GenerateBST(inputArr)
	depths := tree.ListOfDepths()

	expectedDepthVals := [][]*BSTNode{
		{tree},
		{tree.left, tree.right},
		{tree.left.left, tree.left.right, tree.right.left, tree.right.right},
		{tree.left.left.left},
	}

	if len(depths) != 4 {
		t.Error(depths)
	} else {
		for i := range depths {
			if !compareDepths(depths[i], expectedDepthVals[i]) {
				t.Error(depths[i])
			}
		}
	}
}

// Compare a linked list of nodes at a particular depth with an expected slice.
func compareDepths(actual *list.List, expected []*BSTNode) bool {
	if actual.Len() != len(expected) {
		return false
	}

	i := 0
	for el := actual.Front(); el != nil; el = el.Next() {
		if el.Value != expected[i] {
			return false
		} else {
			i++
		}
	}

	return true
}

// Test the CheckBalanced function.
func TestCheckBalanced(t *testing.T) {
	inputArr := []int{1, 3, 4, 8, 10, 11, 15, 21}
	tree := GenerateBST(inputArr)
	if !tree.CheckBalanced() {
		t.Error()
	}

	tree.left.left.left.right = &BSTNode{
		value: 2,
		left:  nil,
		right: nil,
	}

	if tree.CheckBalanced() {
		t.Error()
	}
}

// Test the ValidateBST function.
func TestValidateBST(t *testing.T) {
	inputArr := []int{1, 3, 4, 8, 10, 11, 15, 21}
	tree := GenerateBST(inputArr)
	if !tree.ValidateBST() {
		t.Error()
	}

	// Add an extra node to invalidate the BST.
	tree.right.right.right = &BSTNode{
		value: 0,
		left:  nil,
		right: nil,
	}

	if tree.ValidateBST() {
		t.Error()
	}
}

// Test the Successor function.
func TestSuccessor(t *testing.T) {
	inputArr := []int{1, 3, 4, 8, 10, 11, 15, 21}
	tree := GenerateBST(inputArr)

	if tree.Successor() != tree.right.left {
		t.Error(tree.Successor())
	}

	if tree.left.left.left.Successor() != tree.left.left {
		t.Error(tree.left.left.left.Successor())
	}

	if tree.left.right.Successor() != tree {
		t.Error(tree.left.right.Successor())
	}

	if tree.right.left.Successor() != tree.right {
		t.Error(tree.right.left.Successor())
	}

	if tree.right.right.Successor() != nil {
		t.Error(tree.right.right.Successor())
	}
}

// Test the FindBuildOrder function
func TestFindBuildOrder(t *testing.T) {
	projects := []string{"a", "b", "c", "d", "e", "f"}
	dependencies := [][]string{
		{"a", "d"},
		{"f", "b"},
		{"b", "d"},
		{"f", "a"},
		{"d", "c"},
	}
	expectedOrder := []string{"f", "e", "b", "a", "d", "c"}

	order, err := FindBuildOrder(projects, dependencies)

	if err != nil {
		t.Error(err)
	}

	if len(order) != len(expectedOrder) {
		t.Error(order)
	} else {
		for i := range order {
			if order[i] != expectedOrder[i] {
				t.Error(order)
			}
		}
	}

	circularDeps := [][]string{
		{"a", "b"},
		{"b", "c"},
		{"c", "a"},
	}
	order, err = FindBuildOrder(projects, circularDeps)

	if err == nil {
		t.Error(order)
	}
}

// Test finding the common ancestor of two nodes in a binary tree.
func TestFindCommonAncestor(t *testing.T) {
	tree := &BinTreeNode{
		name: "A",
		left: &BinTreeNode{
			name: "B",
			left: &BinTreeNode{
				name:  "C",
				left:  nil,
				right: nil,
			},
			right: nil,
		},
		right: &BinTreeNode{
			name: "D",
			left: &BinTreeNode{
				name: "E",
				left: &BinTreeNode{
					name:  "F",
					left:  nil,
					right: nil,
				},
				right: &BinTreeNode{
					name:  "G",
					left:  nil,
					right: nil,
				},
			},
			right: &BinTreeNode{
				name:  "H",
				left:  nil,
				right: nil,
			},
		},
	}

	ancestor, err := FindCommonAncestor(
		tree,
		tree.left.left,
		tree.right.left.left,
	)

	if err != nil {
		t.Error(err)
	}

	if ancestor != tree {
		t.Error(ancestor)
	}

	ancestor, err = FindCommonAncestor(
		tree,
		tree.right.left.left,
		tree.right.left.right,
	)

	if err != nil {
		t.Error(err)
	}

	if ancestor != tree.right.left {
		t.Error(ancestor)
	}

	ancestor, err = FindCommonAncestor(
		tree.right,
		tree.right.right,
		tree.left,
	)

	if err == nil {
		t.Error(ancestor)
	}
}

// Test that the correct BST sequences can be found.
func TestBSTSequences(t *testing.T) {
	emptySeqs := FindBSTSequences(nil)
	if emptySeqs != nil {
		t.Error(emptySeqs)
	}

	basicBST := &BSTNode{
		value: 2,
		left: &BSTNode{
			value: 1,
			left:  nil,
			right: nil,
		},
		right: &BSTNode{
			value: 3,
			left:  nil,
			right: nil,
		},
	}

	expectedSeqs := [][]int{
		{2, 1, 3},
		{2, 3, 1},
	}
	foundSeqs := FindBSTSequences(basicBST)

	if !compareSeqs(foundSeqs, expectedSeqs) {
		t.Error(foundSeqs)
	}

	nonBasicBST := &BSTNode{
		value: 5,
		left: &BSTNode{
			value: 3,
			left: &BSTNode{
				value: 1,
				left:  nil,
				right: nil,
			},
			right: nil,
		},
		right: &BSTNode{
			value: 7,
			left: &BSTNode{
				value: 6,
				left:  nil,
				right: nil,
			},
			right: nil,
		},
	}
	expectedSeqs = [][]int{
		{5, 3, 7, 1, 6},
		{5, 3, 7, 6, 1},
		{5, 3, 1, 7, 6},
		{5, 7, 3, 6, 1},
		{5, 7, 3, 1, 6},
		{5, 7, 6, 3, 1},
	}

	foundSeqs = FindBSTSequences(nonBasicBST)

	if !compareSeqs(foundSeqs, expectedSeqs) {
		t.Error(foundSeqs)
	}
}

func compareSeqs(foundSeqs, expectedSeqs [][]int) bool {
	if len(foundSeqs) != len(expectedSeqs) {
		return false
	} else {
		for i := range expectedSeqs {
			if len(foundSeqs[i]) != len(expectedSeqs[i]) {
				return false
			} else {
				for j := range expectedSeqs[i] {
					if foundSeqs[i][j] != expectedSeqs[i][j] {
						return false
					}
				}
			}
		}
	}

	return true
}

// Test the CheckSubtree function.
func TestCheckSubtree(t *testing.T) {
	if !CheckSubtree(nil, nil) {
		t.Error()
	}

	T1 := &BinTreeNode{
		name: "A",
		left: &BinTreeNode{
			name: "B",
			left: &BinTreeNode{
				name:  "C",
				left:  nil,
				right: nil,
			},
			right: &BinTreeNode{
				name:  "D",
				left:  nil,
				right: nil,
			},
		},
		right: &BinTreeNode{
			name: "E",
			left: &BinTreeNode{
				name:  "F",
				left:  nil,
				right: nil,
			},
			right: nil,
		},
	}

	T2 := &BinTreeNode{
		name: "B",
		left: &BinTreeNode{
			name:  "C",
			left:  nil,
			right: nil,
		},
		right: &BinTreeNode{
			name:  "D",
			left:  nil,
			right: nil,
		},
	}

	if CheckSubtree(T1, nil) {
		t.Error()
	}

	if CheckSubtree(nil, T2) {
		t.Error()
	}

	if !CheckSubtree(T1, T2) {
		t.Error()
	}

	T3 := &BinTreeNode{
		name: "B",
		left: &BinTreeNode{
			name:  "C",
			left:  nil,
			right: nil,
		},
		right: &BinTreeNode{
			name: "D",
			left: &BinTreeNode{
				name:  "E",
				left:  nil,
				right: nil,
			},
			right: nil,
		},
	}

	if CheckSubtree(T1, T3) {
		t.Error()
	}
}

// Test getting a value from a counted BST at a specific index in the in-order
// traversal.
func TestGetNodeAtIndex(t *testing.T) {
	var root *CountedBSTNode
	node, err := root.GetNodeAtIndex(0)
	if err == nil {
		t.Error(node)
	}

	root = InitCountedBST(10)
	root.Insert(4)
	root.Insert(13)
	root.Insert(2)
	root.Insert(8)
	root.Insert(15)

	node, err = root.GetNodeAtIndex(3)
	if err != nil {
		t.Error(err)
	}
	if node == nil || node.value != 10 {
		t.Error(node)
	}

	node, err = root.GetNodeAtIndex(5)
	if err != nil {
		t.Error(err)
	}
	if node == nil || node.value != 15 {
		t.Error(node)
	}

	node, err = root.GetNodeAtIndex(-1)
	if err == nil {
		t.Error(node)
	}

	node, err = root.GetNodeAtIndex(6)
	if err == nil {
		t.Error(node)
	}
}

// Test finding all paths through a tree that sum to a given total.
func TestPathsWithSum(t *testing.T) {
	if PathsWithSum(nil, 9) != 0 {
		t.Error()
	}

	tree := &IntBinTree{
		value: 2,
		left: &IntBinTree{
			value: -2,
			left: &IntBinTree{
				value: 1,
				left:  nil,
				right: nil,
			},
			right: &IntBinTree{
				value: 1,
				left:  nil,
				right: nil,
			},
		},
		right: &IntBinTree{
			value: 3,
			left: &IntBinTree{
				value: -5,
				left: &IntBinTree{
					value: 3,
					left:  nil,
					right: nil,
				},
				right: nil,
			},
			right: &IntBinTree{
				value: -4,
				left:  nil,
				right: &IntBinTree{
					value: 2,
					left:  nil,
					right: nil,
				},
			},
		},
	}

	numPaths := PathsWithSum(tree, 1)
	if numPaths != 7 {
		t.Error(numPaths)
	}
}
