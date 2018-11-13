package bipartite

import (
	"testing"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/graph"
)

type testCase struct {
	graphSize    int
	edges        [][]int
	oddCyclePath []int
	colorPath    []bool
}

func TestOddCycle(t *testing.T) {
	var testCases = []testCase{
		{
			graphSize:    4,
			edges:        [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}},
			oddCyclePath: []int{1, 3, 2, 1},
			colorPath:    []bool{true, true, false, true},
		},
		{
			graphSize:    1,
			edges:        [][]int{{0, 0}},
			oddCyclePath: []int{0, 0},
			colorPath:    []bool{false, false},
		},
		{
			graphSize:    6,
			edges:        [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {1, 5}, {5, 2}},
			oddCyclePath: []int{1, 5, 2, 1},
			colorPath:    []bool{true, true, false, true},
		},
	}

	for _, tc := range testCases {
		g := graph.NewGraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		b := NewBipartite(&g)
		if !b.HasOddCycle() {
			t.Errorf("Found no cycle, want odd-lengthy cycle")
		}

		cycle := b.OddCycle()
		for i, cp := range tc.oddCyclePath {
			if cp != cycle[i] {
				t.Errorf("Got %v for oddcycle path, want %v", cycle[i], cp)
			}

			if b.Color(cp) != tc.colorPath[i] {
				t.Errorf("Got color %v for vertex %v, want %v", b.Color(i), cycle[i], tc.colorPath[i])
			}
		}
	}
}
