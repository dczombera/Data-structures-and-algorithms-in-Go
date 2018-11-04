package dfs_cycle

import (
	"testing"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/graph"
)

type testCase struct {
	graphSize int
	edges     [][]int
	cyclePath []int
}

func TestCycleSelfLoop(t *testing.T) {
	var testCases = []testCase{
		{
			graphSize: 3,
			edges:     [][]int{{0, 1}, {1, 2}, {2, 2}},
			cyclePath: []int{2, 2},
		},
		{
			graphSize: 2,
			edges:     [][]int{{1, 1}, {2, 2}, {1, 2}},
			cyclePath: []int{1, 1},
		},
	}

	for _, tc := range testCases {
		g := graph.NewGraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		c := NewDFSCyle(&g)
		if !c.HasCycle() {
			t.Errorf("Found no cycle, want self loop cycle")
		}

		cycle := c.Cycle()
		for i, cp := range tc.cyclePath {
			if cp != cycle[i] {
				t.Errorf("Got %v for cycle path, want %v", cycle[i], cp)
			}
		}

	}
}

func TestCycleHasParallelEdges(t *testing.T) {
	var testCases = []testCase{
		{
			graphSize: 3,
			edges:     [][]int{{0, 1}, {0, 1}, {1, 2}},
			cyclePath: []int{0, 1, 0},
		},
		{
			graphSize: 4,
			edges:     [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 1}, {3, 2}},
			cyclePath: []int{1, 2, 1},
		},
	}

	for _, tc := range testCases {
		g := graph.NewGraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		c := NewDFSCyle(&g)
		if !c.HasCycle() {
			t.Errorf("Found no cycle, want parallel edges cycle")
		}

		cycle := c.Cycle()
		for i, cp := range tc.cyclePath {
			if cp != cycle[i] {
				t.Errorf("Got %v for cycle path, want %v", cycle[i], cp)
			}
		}

	}
}

func TestCycle(t *testing.T) {
	var testCases = []testCase{
		{
			graphSize: 4,
			edges:     [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}},
			cyclePath: []int{3, 2, 1, 0, 3},
		},
		{
			graphSize: 5,
			edges:     [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}, {3, 4}},
			cyclePath: []int{3, 2, 1, 3},
		},
	}

	for _, tc := range testCases {
		g := graph.NewGraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		c := NewDFSCyle(&g)
		if !c.HasCycle() {
			t.Errorf("Found no cycle, want regular cycle")
		}

		cycle := c.Cycle()
		for i, cp := range tc.cyclePath {
			if cp != cycle[i] {
				t.Errorf("Got %v for cycle path, want %v", cycle[i], cp)
			}
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
		g := graph.NewGraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		c := NewDFSCyle(&g)
		if c.HasCycle() {
			t.Errorf("Found cycle, want no cycle")
		}
	}
}
