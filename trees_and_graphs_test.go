package ctci

import "testing"

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
