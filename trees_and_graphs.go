package ctci

import "container/list"

type GraphNode struct {
	name     string
	adjacent []*GraphNode
	visited  bool
}

type Graph struct {
	nodes []GraphNode
}

func (graph Graph) RouteExists(nodeS, nodeE *GraphNode) bool {
	queue := NewBasicQueue()
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

// A node in a binary search tree.
type BSTNode struct {
	value int
	left  *BSTNode
	right *BSTNode
}

// Generates a binary search tree of minimal height from a sorted array of
// unique integer elements.
func GenerateBST(sortedArr []int) *BSTNode {
	if len(sortedArr) == 0 {
		return nil
	}

	midpoint := (len(sortedArr) / 2)

	return &BSTNode{
		value: sortedArr[midpoint],
		left:  GenerateBST(sortedArr[:midpoint]),
		right: GenerateBST(sortedArr[midpoint+1:]),
	}
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
