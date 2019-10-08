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
