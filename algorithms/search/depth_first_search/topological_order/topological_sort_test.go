package topological_order

import (
	"testing"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/directed_graph"
)

type testCase struct {
	graphSize int
	edges     [][]int
	order     []int
	pos       []int
}

func TestTopologicalOrder(t *testing.T) {
	var testCases = []testCase{
		{
			graphSize: 13,
			edges:     [][]int{{0, 5}, {5, 4}, {0, 1}, {0, 6}, {6, 9}, {9, 11}, {9, 10}, {11, 12}, {2, 3}, {8, 7}},
			order:     []int{8, 7, 2, 3, 0, 6, 9, 10, 11, 12, 1, 5, 4},
			pos:       []int{4, 10, 2, 3, 12, 11, 5, 1, 0, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		g := directed_graph.NewDigraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		topo, err := NewTopologicalOrder(g)
		if err != nil {
			t.Errorf("Got error %v, want topological order struct", err)
		}

		if !topo.HasOrder() || !topo.IsDAG() {
			t.Errorf("Got no order, want topological order")
		}

		curr := topo.Order().First
		for _, o := range tc.order {
			if curr.Item != o {
				t.Errorf("Got %v in topological order, want %v", curr.Item, o)
			}
			curr = curr.Next
		}

		for i, r := range tc.pos {
			if topo.Rank(i) != r {
				t.Errorf("Got position %v for vertex %v, want %v", topo.Rank(i), i, r)
			}
		}
	}
}

func TestNoTopologicalOrder(t *testing.T) {
	testCase := struct {
		graphSize int
		edges     [][]int
	}{
		graphSize: 4,
		edges:     [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}},
	}

	g := directed_graph.NewDigraph(testCase.graphSize)
	for _, e := range testCase.edges {
		g.AddEdge(e[0], e[1])
	}

	topo, err := NewTopologicalOrder(g)
	if err == nil {
		t.Errorf("Got no error, want error for non-existing topological order")
	}

	if topo.HasOrder() || topo.IsDAG() {
		t.Errorf("Got topological order, even though no topological order exists")
	}
}
