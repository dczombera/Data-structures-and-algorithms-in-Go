package directed_graph

import "testing"

type testCase struct {
	size          int
	edges         [][]int
	wantEdges     []VertexEdges
	wantOutDegree []VertexDegree
	wantInDegree  []VertexDegree
}

type VertexEdges struct {
	vertex int
	adj    []int
}

type VertexDegree struct {
	vertex int
	degree int
}

var testCases = []testCase{
	{
		size:          4,
		edges:         [][]int{{0, 1}, {1, 2}, {2, 3}},
		wantEdges:     []VertexEdges{{0, []int{1}}, {1, []int{2}}, {2, []int{3}}},
		wantOutDegree: []VertexDegree{{0, 1}, {1, 1}, {2, 1}, {3, 0}},
		wantInDegree:  []VertexDegree{{0, 0}, {1, 1}, {2, 1}, {3, 1}},
	},
	{
		size:          9,
		edges:         [][]int{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}, {2, 1}, {3, 7}, {4, 6}, {3, 5}},
		wantEdges:     []VertexEdges{{0, []int{1}}, {1, []int{0, 1, 2, 3, 4, 5, 6, 7}}, {2, []int{1}}, {3, []int{7, 5}}, {4, []int{6}}},
		wantOutDegree: []VertexDegree{{0, 1}, {1, 8}, {2, 1}, {3, 2}, {4, 1}, {5, 0}, {6, 0}, {7, 0}, {8, 0}},
		wantInDegree:  []VertexDegree{{0, 1}, {1, 3}, {2, 1}, {3, 1}, {4, 1}, {5, 2}, {6, 2}, {7, 2}, {8, 0}},
	},
}

func containsSameVertices(got, want []int) bool {
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
		dg := NewDigraph(tc.size)
		for _, e := range tc.edges {
			dg.AddEdge(e[0], e[1])
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

func TestReverseDigraph(t *testing.T) {
	testCases := []struct {
		size         int
		edges        [][]int
		reverseEdges []VertexEdges
	}{
		{

			size:         5,
			edges:        [][]int{{4, 3}, {3, 2}, {2, 1}, {1, 0}},
			reverseEdges: []VertexEdges{{0, []int{1}}, {1, []int{2}}, {2, []int{3}}, {3, []int{4}}},
		},
		{
			size:         5,
			edges:        [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			reverseEdges: []VertexEdges{{1, []int{0}}, {2, []int{1}}, {3, []int{2}}, {4, []int{3}}},
		},
	}

	for _, tc := range testCases {
		dg := NewDigraph(tc.size).Reverse()
		for _, e := range tc.edges {
			dg.AddEdge(e[0], e[1])
		}

		reverseDg := dg.Reverse()
		for _, reverseEdge := range tc.reverseEdges {
			if !containsSameVertices(reverseDg.AdjacencyList(reverseEdge.vertex), reverseEdge.adj) {
				t.Errorf("Adjacency list of vertex %v do not include correct vertices", reverseEdge.vertex)
			}
		}
	}
}
