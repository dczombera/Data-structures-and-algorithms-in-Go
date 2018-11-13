package dfs_connected_components

import (
	"testing"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/graph"
)

type testCase struct {
	graphSize         int
	edges             [][]int
	connectedVertices [][]int
	verticesIds       []VerticeId
	countCC           int
}

type VerticeId struct {
	vertex int
	id     int
}

var testCases = []testCase{
	{
		graphSize:         7,
		edges:             [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}},
		connectedVertices: [][]int{{0, 1}, {2, 6}, {5, 0}, {3, 2}, {6, 0}},
		verticesIds:       []VerticeId{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}},
		countCC:           1,
	},
	{
		graphSize: 8,
		edges:     [][]int{{0, 1}, {1, 2}, {2, 3}, {4, 5}, {5, 6}, {6, 7}},
		connectedVertices: [][]int{
			{0, 0}, {0, 1}, {0, 2}, {0, 3},
			{1, 0}, {1, 1}, {1, 2}, {1, 3},
			{2, 0}, {2, 1}, {2, 2}, {2, 3},
			{3, 0}, {3, 1}, {3, 2}, {3, 3},
			{4, 4}, {4, 5}, {4, 6}, {4, 7},
			{5, 4}, {5, 5}, {5, 6}, {5, 7},
			{6, 4}, {6, 5}, {6, 6}, {6, 7},
			{7, 4}, {7, 5}, {7, 6}, {7, 7},
		},
		verticesIds: []VerticeId{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {5, 1}, {6, 1}, {7, 1}},
		countCC:     2,
	},
	{
		graphSize:         5,
		edges:             [][]int{},
		connectedVertices: [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}},
		verticesIds:       []VerticeId{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}},
		countCC:           5,
	},
}

func TestCCCount(t *testing.T) {
	for _, tc := range testCases {
		g := graph.NewGraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		cc := NewCC(&g)
		if cc.CountCC() != tc.countCC {
			t.Errorf("Got connected components count of %v, want %v", cc.CountCC(), tc.countCC)
		}
	}
}

func TestCCConnected(t *testing.T) {
	for _, tc := range testCases {
		g := graph.NewGraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		cc := NewCC(&g)
		for _, vv := range tc.connectedVertices {
			if !cc.Connected(vv[0], vv[1]) {
				t.Errorf("Got vertices %v and %v are not connected, want them to be connected", vv[0], vv[1])
			}
		}
	}
}

func TestCCId(t *testing.T) {
	for _, tc := range testCases {
		g := graph.NewGraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		cc := NewCC(&g)
		for _, vid := range tc.verticesIds {
			id := cc.Id(vid.vertex)
			if id != vid.id {
				t.Errorf("Got vertex id of %v, want %v", id, vid.id)
			}
		}
	}
}
