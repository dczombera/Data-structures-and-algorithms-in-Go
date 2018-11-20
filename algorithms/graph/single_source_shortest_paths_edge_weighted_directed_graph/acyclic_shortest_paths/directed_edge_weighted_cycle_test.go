package acyclic_shortest_paths

import (
	"testing"

	graph "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_directed_graph"
)

type testCaseCycle struct {
	graphSize int
	edges     []graph.DirectedEdge
	cyclePath []graph.DirectedEdge
}

func TestDirectedEdgeWeightedCycle(t *testing.T) {
	var testCases = []testCaseCycle{
		{
			graphSize: 4,
			edges:     []graph.DirectedEdge{{0, 1, 0.1}, {1, 2, 1.2}, {2, 3, 2.3}, {3, 0, 3.0}},
			cyclePath: []graph.DirectedEdge{{0, 1, 0.1}, {1, 2, 1.2}, {2, 3, 2.3}, {3, 0, 3.0}},
		},
		{
			graphSize: 5,
			edges:     []graph.DirectedEdge{{0, 1, 0.1}, {1, 2, 1.2}, {2, 3, 2.3}, {3, 1, 3.1}, {3, 4, 3.4}},
			cyclePath: []graph.DirectedEdge{{1, 2, 1.2}, {2, 3, 2.3}, {3, 1, 3.1}},
		},
	}

	for _, tc := range testCases {
		g := graph.NewEdgeWeightedDigraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e)
		}

		c := NewDirectedEdgeWeightedCycle(g)
		if !c.HasCycle() {
			t.Errorf("Found no cycle, want directed edge weighted cycle")
		}

		cycle := c.Cycle()
		curr := cycle.First
		for _, cp := range tc.cyclePath {
			if cp != curr.Item {
				t.Errorf("Got %v for cycle path, want %v", curr.Item, cp)
			}
			curr = curr.Next
		}
	}
}

func TestNoDirectedEdgeWeightedCycle(t *testing.T) {
	var testCases = []testCaseCycle{
		{
			graphSize: 4,
			edges:     []graph.DirectedEdge{{3, 1, 3.1}, {2, 0, 2.0}, {2, 1, 2.1}},
			cyclePath: []graph.DirectedEdge{},
		},
		{
			graphSize: 8,
			edges:     []graph.DirectedEdge{{0, 1, 0.1}, {1, 2, 1.2}, {2, 3, 2.3}, {3, 4, 3.4}, {4, 5, 4.5}, {5, 6, 5.6}, {6, 7, 6.7}},
			cyclePath: []graph.DirectedEdge{},
		},
	}

	for _, tc := range testCases {
		g := graph.NewEdgeWeightedDigraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e)
		}

		c := NewDirectedEdgeWeightedCycle(g)
		if c.HasCycle() {
			t.Errorf("Found directed edge weighted cycle, want no cycle")
		}
	}
}
