package directed_cycle

import (
	"testing"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/directed_graph"
)

type testCase struct {
	graphSize int
	edges     [][]int
	cyclePath []int
}

func TestCycle(t *testing.T) {
	var testCases = []testCase{
		{
			graphSize: 4,
			edges:     [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}},
			cyclePath: []int{3, 0, 1, 2, 3},
		},
		{
			graphSize: 5,
			edges:     [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}, {3, 4}},
			cyclePath: []int{3, 1, 2, 3},
		},
	}

	for _, tc := range testCases {
		g := directed_graph.NewDigraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		c := NewDirectedCycle(g)
		if !c.HasCycle() {
			t.Errorf("Found no cycle, want directed cycle")
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

func TestNoCycle(t *testing.T) {
	var testCases = []testCase{
		{
			graphSize: 4,
			edges:     [][]int{{3, 1}, {2, 0}, {2, 1}},
			cyclePath: []int{},
		},
		{
			graphSize: 8,
			edges:     [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}},
			cyclePath: []int{},
		},
	}

	for _, tc := range testCases {
		g := directed_graph.NewDigraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		c := NewDirectedCycle(g)
		if c.HasCycle() {
			t.Errorf("Found cycle, want no cycle")
		}
	}
}
