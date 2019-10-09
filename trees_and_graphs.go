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
