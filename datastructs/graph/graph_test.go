package graph

import "testing"

type testCase struct {
	size       int
	edges      [][]int
	wantEdges  []VertexEdges
	wantDegree []VertexDegree
}

type VertexEdges struct {
	vertex int
	adj    []Value
}

type VertexDegree struct {
	vertex int
	degree int
}

var testCases = []testCase{
	{8, [][]int{[]int{0, 1}, []int{2, 4}, []int{1, 5}, []int{7, 5}, []int{7, 1}}, []VertexEdges{{0, []Value{1}}, {1, []Value{7, 5, 0}}, {2, []Value{4}}, {4, []Value{2}}, {5, []Value{7, 1}}, {7, []Value{1, 5}}}, []VertexDegree{{0, 1}, {1, 3}, {2, 1}, {4, 1}, {5, 2}, {7, 2}}},
}

func containsVertices(has Bag, want []Value) bool {
	for _, w := range want {
		foundVertex := false
		for _, v := range has.items {
			if w == v {
				foundVertex = true
				break
			}
		}
		if !foundVertex {
			return false
		}
	}
	return true
}

func TestGraph(t *testing.T) {
	for _, tc := range testCases {
		g := NewGraph(tc.size)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}
		for _, we := range tc.wantEdges {
			adj := g.AdjacencyList(we.vertex)
			if !containsVertices(adj, we.adj) {
				t.Errorf("Adjacency lists do not include correct vertices")
			}
		}

		for _, wd := range tc.wantDegree {
			degree := g.AdjacencyList(wd.vertex).Size
			if degree != wd.degree {
				t.Errorf("Got degree of %v for vertex %v, want degree of %v", degree, wd.vertex, wd.degree)
			}
		}
	}
}
