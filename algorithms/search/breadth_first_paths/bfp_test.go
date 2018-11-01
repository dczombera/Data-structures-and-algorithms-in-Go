package breadth_first_paths

import (
	"testing"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/graph"
)

type shortestPath struct {
	target int
	distTo int
	path   []int
}
type testCase struct {
	graphSize    int
	edges        [][]int
	sourceVertex int
	paths        []shortestPath
}

var testCases = []testCase{
	{
		graphSize:    7,
		edges:        [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}},
		sourceVertex: 0,
		paths:        []shortestPath{{2, 2, []int{0, 1, 2}}, {6, 6, []int{0, 1, 2, 3, 4, 5, 6}}},
	},
	{
		graphSize:    6,
		edges:        [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {2, 5}, {3, 5}},
		sourceVertex: 1,
		paths:        []shortestPath{{4, 1, []int{1, 4}}, {3, 2, []int{1, 2, 3}}, {5, 2, []int{1, 2, 5}}},
	},
}

func TestBreadthFirstPaths(t *testing.T) {
	for _, tc := range testCases {
		g := graph.NewGraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		bfp := NewBreadthFirstPaths(&g, tc.sourceVertex)
		for _, shortestP := range tc.paths {
			if !bfp.HasPathTo(shortestP.target) {
				t.Errorf("Vertex %v has no path to source vertex %v", shortestP.target, tc.sourceVertex)
			}

			d, e := bfp.DistTo(shortestP.target)
			if e != nil {
				t.Errorf("Got error %v for distance from %v to %v, want %v", e, tc.sourceVertex, shortestP.target, shortestP.distTo)
			}
			if d != shortestP.distTo {
				t.Errorf("Got a distance of %v from %v to %v, want distance %v", d, tc.sourceVertex, shortestP.target, shortestP.distTo)
			}

			path, e := bfp.PathTo(shortestP.target)
			if e != nil {
				t.Errorf("Got error %v, want path from %v to %v", e, tc.sourceVertex, shortestP.target)
			}
			curr := path.First
			for _, w := range shortestP.path {
				if curr.Item != w {
					t.Errorf("Got vertex %v in path from %v to %v, want %v", curr.Item, tc.sourceVertex, shortestP.target, w)
				}
				curr = curr.Next
			}

		}
	}
}
