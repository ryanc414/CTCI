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

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
