package transitive_closure

import (
	"testing"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/directed_graph"
)

type testCase struct {
	graphSize    int
	edges        [][]int
	reachability []Reachable
}

type Reachable struct {
	from int
	to   int
}

func TestTransitiveClosure(t *testing.T) {
	var testCases = []testCase{
		{
			graphSize:    4,
			edges:        [][]int{{0, 1}, {1, 2}, {2, 3}},
			reachability: []Reachable{{0, 1}, {1, 2}, {2, 3}},
		},
		{
			graphSize:    13,
			edges:        [][]int{{0, 5}, {5, 4}, {0, 1}, {0, 6}, {6, 4}, {6, 9}, {9, 11}, {9, 10}, {11, 12}, {2, 3}, {8, 7}, {7, 6}},
			reachability: []Reachable{{0, 4}, {6, 4}, {6, 12}, {8, 6}, {11, 12}},
		},
	}
	for _, tc := range testCases {
		g := directed_graph.NewDigraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		transitive := NewTransitiveClosure(g)
		for _, r := range tc.reachability {
			if !transitive.Reachable(r.from, r.to) {
				t.Errorf("Got vertex %v not reachable from vertex %v, want it to be reachable", r.to, r.from)
			}
		}
	}
}

func TestTransitiveClosureNonReachable(t *testing.T) {
	var testCases = []testCase{
		{
			graphSize:    4,
			edges:        [][]int{{0, 1}, {1, 2}, {2, 3}},
			reachability: []Reachable{{1, 0}, {3, 2}, {2, 0}},
		},
		{
			graphSize:    13,
			edges:        [][]int{{0, 5}, {5, 4}, {0, 1}, {0, 6}, {6, 4}, {6, 9}, {9, 11}, {9, 10}, {11, 12}, {2, 3}, {8, 7}, {7, 6}},
			reachability: []Reachable{{1, 4}, {9, 7}, {6, 8}, {5, 12}, {10, 8}},
		},
	}
	for _, tc := range testCases {
		g := directed_graph.NewDigraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		transitive := NewTransitiveClosure(g)
		for _, r := range tc.reachability {
			if transitive.Reachable(r.from, r.to) {
				t.Errorf("Got vertex %v reachable from vertex %v, want it to be non-reachable", r.to, r.from)
			}
		}
	}
}
