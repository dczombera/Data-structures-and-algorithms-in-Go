package acyclic_shortest_paths

import (
	"testing"

	graph "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_directed_graph"
)

type testCaseTopological struct {
	graphSize int
	edges     []graph.DirectedEdge
	order     []int
	rank      []int
}

func TestTopologicalOrder(t *testing.T) {
	var testCases = []testCaseTopological{
		{
			graphSize: 13,
			edges:     []graph.DirectedEdge{{0, 5, 0.5}, {5, 4, 5.4}, {0, 1, 0.1}, {0, 6, 0.6}, {6, 9, 6.9}, {9, 11, 9.11}, {9, 10, 9.10}, {11, 12, 11.12}, {2, 3, 2.3}, {8, 7, 8.7}},
			order:     []int{8, 7, 2, 3, 0, 6, 9, 10, 11, 12, 1, 5, 4},
			rank:      []int{4, 10, 2, 3, 12, 11, 5, 1, 0, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		g := graph.NewEdgeWeightedDigraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e)
		}

		to, err := NewTopologicalOrder(g)
		if err != nil {
			t.Errorf("Got error %v, want topological order struct", err)
		}

		if !to.HasOrder() {
			t.Errorf("Got no order, want topological order")
		}

		curr := to.Order().First
		for _, o := range tc.order {
			if curr.Item != o {
				t.Errorf("Got %v in topological order, want %v", curr.Item, o)
			}
			curr = curr.Next
		}

		for i, r := range tc.rank {
			if to.Rank(i) != r {
				t.Errorf("Got position %v for vertex %v, want %v", to.Rank(i), i, r)
			}
		}
	}
}

func TestNoTopologicalOrder(t *testing.T) {
	testCase := struct {
		graphSize int
		edges     []graph.DirectedEdge
	}{
		graphSize: 4,
		edges:     []graph.DirectedEdge{{0, 1, 0.1}, {1, 2, 1.2}, {2, 3, 2.3}, {3, 1, 3.1}},
	}

	g := graph.NewEdgeWeightedDigraph(testCase.graphSize)
	for _, e := range testCase.edges {
		g.AddEdge(e)
	}

	to, err := NewTopologicalOrder(g)
	if err == nil {
		t.Errorf("Got no error, want error for non-existing topological order")
	}

	if to.HasOrder() {
		t.Errorf("Got topological order, even though no topological order exists")
	}
}
