package kosaraju_sharir_scc

import (
	"testing"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/directed_graph"
)

type testCase struct {
	graphSize         int
	edges             [][]int
	connectedVertices [][]int
	verticesIds       []VerticeId
	count             int
}

type VerticeId struct {
	vertex int
	id     int
}

var testCases = []testCase{
	{
		graphSize:         6,
		edges:             [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 0}},
		connectedVertices: [][]int{{0, 1}, {2, 5}, {5, 1}, {3, 2}, {2, 4}},
		verticesIds:       []VerticeId{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}},
		count:             1,
	},
	{
		graphSize: 8,
		edges:     [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {4, 5}, {5, 6}, {6, 7}, {7, 4}},
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
		verticesIds: []VerticeId{{0, 1}, {1, 1}, {2, 1}, {3, 1}, {4, 0}, {5, 0}, {6, 0}, {7, 0}},
		count:       2,
	},
	{
		graphSize:         5,
		edges:             [][]int{},
		connectedVertices: [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}},
		verticesIds:       []VerticeId{{0, 4}, {1, 3}, {2, 2}, {3, 1}, {4, 0}},
		count:             5,
	},
}

func TestKosarajuSharirCount(t *testing.T) {
	for _, tc := range testCases {
		g := directed_graph.NewDigraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		scc := NewKosarajuSharirSCC(g)
		if scc.Count() != tc.count {
			t.Errorf("Got strong components count of %v, want %v", scc.Count(), tc.count)
		}
	}
}

func TestKosarajuSharirConnected(t *testing.T) {
	for _, tc := range testCases {
		g := directed_graph.NewDigraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		scc := NewKosarajuSharirSCC(g)
		for _, vv := range tc.connectedVertices {
			if !scc.StronglyConnected(vv[0], vv[1]) {
				t.Errorf("Got vertices %v and %v are not strongly connected, want them to be strongly connected", vv[0], vv[1])
			}
		}
	}
}

func TestKosarajuSharirId(t *testing.T) {
	for _, tc := range testCases {
		g := directed_graph.NewDigraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		scc := NewKosarajuSharirSCC(g)
		for _, vid := range tc.verticesIds {
			id := scc.Id(vid.vertex)
			if id != vid.id {
				t.Errorf("Got id %v for vertex %v, want %v", id, vid.vertex, vid.id)
			}
		}
	}
}
