package dfp

import (
	"testing"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/graph"
)

type pathInfo struct {
	target int
	path   []int
}
type testCase struct {
	graphSize              int
	edges                  [][]int
	sourceVertex           int
	paths                  []pathInfo
	connectedVertices      []int
	countConnectedVertices int
}

var testCases = []testCase{
	{
		graphSize:              7,
		edges:                  [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}},
		sourceVertex:           0,
		paths:                  []pathInfo{{2, []int{0, 1, 2}}, {6, []int{0, 1, 2, 3, 4, 5, 6}}},
		connectedVertices:      []int{0, 1, 2, 3, 4, 5, 6},
		countConnectedVertices: 7,
	},
	{
		graphSize:              6,
		edges:                  [][]int{{1, 2}, {0, 5}, {1, 4}, {5, 3}},
		sourceVertex:           1,
		paths:                  []pathInfo{{4, []int{1, 4}}, {2, []int{1, 2}}},
		connectedVertices:      []int{1, 2, 4},
		countConnectedVertices: 3,
	},
}

func TestDepthFirstPaths(t *testing.T) {
	for _, tc := range testCases {
		g := graph.NewGraph(tc.graphSize)
		for _, edge := range tc.edges {
			g.AddEdge(edge[0], edge[1])
		}

		dfp := NewDFP(&g, tc.sourceVertex)
		for _, w := range tc.connectedVertices {
			if !dfp.HasPathTo(w) {
				t.Errorf("Got vertex %v not connected with source vertex %v, want it to be connected.", w, tc.sourceVertex)
			}
		}

		if dfp.Count() != tc.countConnectedVertices {
			t.Errorf("Got number of connected vertices of %v, want %v", dfp.Count(), tc.countConnectedVertices)
		}

		for _, pathInfo := range tc.paths {
			path := dfp.PathTo(pathInfo.target)
			if path.Size != len(pathInfo.path) {
				t.Errorf("Got path size of %v, want %v", path.Size, len(pathInfo.path))
			}

			curr := path.First
			for _, w := range pathInfo.path {
				if curr.Item != w {
					t.Errorf("Got vertex %v for path from source %v to %v, want %v", curr.Item, dfp.sourceVertex, pathInfo.target, w)
				}
				curr = curr.Next
			}
		}
	}
}

func TestDepthFirstPathsNotConnectedGraph(t *testing.T) {
	testCases := []struct {
		graphSize             int
		edges                 [][]int
		sourceVertex          int
		notConntectedVertices []int
	}{
		{
			graphSize:             6,
			edges:                 [][]int{{1, 2}, {0, 5}, {1, 4}, {5, 3}},
			sourceVertex:          1,
			notConntectedVertices: []int{0, 5, 3},
		},
	}

	for _, tc := range testCases {
		g := graph.NewGraph(tc.graphSize)
		for _, edge := range tc.edges {
			g.AddEdge(edge[0], edge[1])
		}

		dfp := NewDFP(&g, tc.sourceVertex)

		for _, w := range tc.notConntectedVertices {
			path := dfp.PathTo(w)
			if path != nil {
				t.Errorf("Got non-empty path from source %v to %v, want nil since there exists no path", dfp.sourceVertex, w)
			}
		}
	}
}
