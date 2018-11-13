package depth_first_directed_paths

import (
	"testing"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/directed_graph"
)

type pathInfo struct {
	target int
	path   []int
}
type testCase struct {
	graphSize         int
	edges             [][]int
	sourceVertex      int
	paths             []pathInfo
	connectedVertices []int
}

var testCases = []testCase{
	{
		graphSize:         7,
		edges:             [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}},
		sourceVertex:      0,
		paths:             []pathInfo{{2, []int{0, 1, 2}}, {6, []int{0, 1, 2, 3, 4, 5, 6}}},
		connectedVertices: []int{0, 1, 2, 3, 4, 5, 6},
	},
	{
		graphSize:         6,
		edges:             [][]int{{1, 2}, {0, 5}, {1, 4}, {5, 3}},
		sourceVertex:      1,
		paths:             []pathInfo{{4, []int{1, 4}}, {2, []int{1, 2}}},
		connectedVertices: []int{1, 2, 4},
	},
	{
		graphSize:         1,
		edges:             [][]int{{0, 0}},
		sourceVertex:      0,
		paths:             []pathInfo{{0, []int{0}}},
		connectedVertices: []int{0},
	},
}

func TestDepthFirstDirectedPaths(t *testing.T) {
	for _, tc := range testCases {
		g := directed_graph.NewDigraph(tc.graphSize)
		for _, edge := range tc.edges {
			g.AddEdge(edge[0], edge[1])
		}

		dfdp := NewDFDP(g, tc.sourceVertex)
		for _, w := range tc.connectedVertices {
			if !dfdp.HasPathTo(w) {
				t.Errorf("Got vertex %v not connected with source vertex %v, want it to be connected.", w, tc.sourceVertex)
			}
		}

		for _, pathInfo := range tc.paths {
			path, err := dfdp.PathTo(pathInfo.target)
			if err != nil {
				t.Errorf("Got error %v, want path from source %v to %v", err, tc.sourceVertex, pathInfo.target)
			}
			if path.Size != len(pathInfo.path) {
				t.Errorf("Got path size of %v, want %v", path.Size, len(pathInfo.path))
			}

			curr := path.First
			for _, w := range pathInfo.path {
				if curr.Item != w {
					t.Errorf("Got vertex %v for path from source %v to %v, want %v", curr.Item, dfdp.sourceVertex, pathInfo.target, w)
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
		g := directed_graph.NewDigraph(tc.graphSize)
		for _, edge := range tc.edges {
			g.AddEdge(edge[0], edge[1])
		}

		dfdp := NewDFDP(g, tc.sourceVertex)

		for _, w := range tc.notConntectedVertices {
			if dfdp.HasPathTo(w) {
				t.Errorf("Found path from source %v to %v, want no path", dfdp.sourceVertex, w)
			}
		}
	}
}
