package ctci

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
