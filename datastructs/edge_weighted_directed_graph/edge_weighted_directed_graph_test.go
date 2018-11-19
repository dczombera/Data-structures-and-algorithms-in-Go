package edge_weighted_directed_graph

import "testing"

type testCase struct {
	size          int
	edges         []DirectedEdge
	wantEdges     []VertexEdges
	wantOutDegree []VertexDegree
	wantInDegree  []VertexDegree
}

type VertexEdges struct {
	vertex int
	adj    []DirectedEdge
}

type VertexDegree struct {
	vertex int
	degree int
}

var testCases = []testCase{
	{
		size:          4,
		edges:         []DirectedEdge{{0, 1, 0.23}, {1, 2, 42.0}, {2, 3, -1.0}},
		wantEdges:     []VertexEdges{{0, []DirectedEdge{{0, 1, 0.23}}}, {1, []DirectedEdge{{1, 2, 42.0}}}, {2, []DirectedEdge{{2, 3, -1.0}}}},
		wantOutDegree: []VertexDegree{{0, 1}, {1, 1}, {2, 1}, {3, 0}},
		wantInDegree:  []VertexDegree{{0, 0}, {1, 1}, {2, 1}, {3, 1}},
	},
	{
		size:          9,
		edges:         []DirectedEdge{{0, 1, 1.0}, {1, 0, 1.0}, {1, 1, 0.1}, {1, 2, 10.0}, {1, 3, 0.42}, {1, 4, 5.65}, {1, 5, 33.22}, {1, 6, -12.0}, {1, 7, 1.1}, {2, 1, 2.1}, {3, 7, -3.1}, {4, 6, 9.8}, {3, 5, 0.23}},
		wantEdges:     []VertexEdges{{0, []DirectedEdge{{0, 1, 1.0}}}, {1, []DirectedEdge{{1, 0, 1.0}, {1, 1, 0.1}, {1, 2, 10.0}, {1, 3, 0.42}, {1, 4, 5.65}, {1, 5, 33.22}, {1, 6, -12.0}, {1, 7, 1.1}}}, {2, []DirectedEdge{{2, 1, 2.1}}}, {3, []DirectedEdge{{3, 7, -3.1}, {3, 5, 0.23}}}, {4, []DirectedEdge{{4, 6, 9.8}}}},
		wantOutDegree: []VertexDegree{{0, 1}, {1, 8}, {2, 1}, {3, 2}, {4, 1}, {5, 0}, {6, 0}, {7, 0}, {8, 0}},
		wantInDegree:  []VertexDegree{{0, 1}, {1, 3}, {2, 1}, {3, 1}, {4, 1}, {5, 2}, {6, 2}, {7, 2}, {8, 0}},
	},
}

func containsSameVertices(got, want []DirectedEdge) bool {
	if len(got) != len(want) {
		return false
	}

	for i, w := range want {
		if w != got[i] {
			return false
		}
	}
	return true
}

func TestDigraph(t *testing.T) {
	for _, tc := range testCases {
		dg := NewEdgeWeightedDigraph(tc.size)
		for _, e := range tc.edges {
			dg.AddEdge(e)
		}

		for _, wantEdge := range tc.wantEdges {
			if !containsSameVertices(dg.AdjacencyList(wantEdge.vertex), wantEdge.adj) {
				t.Errorf("Adjacency list of vertex %v do not include correct vertices", wantEdge.vertex)
			}
		}

		for _, vd := range tc.wantOutDegree {
			if dg.Outdegree(vd.vertex) != vd.degree {
				t.Errorf("Got outdegree of %v for vertex %v, want %v", dg.Outdegree(vd.vertex), vd.vertex, vd.degree)
			}
		}

		for _, vd := range tc.wantInDegree {
			if dg.Indegree(vd.vertex) != vd.degree {
				t.Errorf("Got indegree of %v for vertex %v, want %v", dg.Indegree(vd.vertex), vd.vertex, vd.degree)
			}
		}
	}
}
