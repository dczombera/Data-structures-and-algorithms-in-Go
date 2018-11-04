package dfs

import (
	"testing"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/graph"
)

type testCase struct {
	graphSize               int
	edges                   [][]int
	sourceVertex            int
	connectedVertices       []int
	countConntectedVertices int
}

var testCases = []testCase{
	{7, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}}, 0, []int{0, 1, 2, 3, 4, 5, 6}, 7},
	{6, [][]int{{1, 2}, {0, 5}, {1, 4}, {5, 3}}, 1, []int{1, 2, 4}, 3},
}

func TestDepthFirstSearch(t *testing.T) {
	for _, tc := range testCases {
		g := graph.NewGraph(tc.graphSize)
		for _, edge := range tc.edges {
			g.AddEdge(edge[0], edge[1])
		}

		dfs := NewDFS(&g, tc.sourceVertex)
		for _, w := range tc.connectedVertices {
			if !dfs.IsConnected(w) {
				t.Errorf("Got vertex %v not connected with source vertex %v, want it to be connected.", w, tc.sourceVertex)
			}
		}

		if dfs.Count() != tc.countConntectedVertices {
			t.Errorf("Got number of connected vertices of %v, want %v", dfs.Count(), tc.countConntectedVertices)
		}
	}
}
