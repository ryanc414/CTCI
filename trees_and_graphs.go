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
