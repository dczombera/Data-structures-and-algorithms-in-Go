package edge_weighted_graph

import (
	"testing"

	"github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/graph/edge"
)

type testCase struct {
	size       int
	edges      []edge.Edge
	wantEdges  []VertexEdges
	wantDegree []VertexDegree
	allEdges   []edge.Edge
}

type VertexEdges struct {
	vertex int
	adj    []edge.Edge
}

type VertexDegree struct {
	vertex int
	degree int
}

var edges = []edge.Edge{edge.NewEdge(0, 1, 42.0), edge.NewEdge(1, 2, -3.2), edge.NewEdge(2, 3, 13.37), edge.NewEdge(3, 0, 0.1), edge.NewEdge(4, 4, 23.23)}
var testCases = []testCase{
	{
		size:       5,
		edges:      edges,
		wantEdges:  []VertexEdges{{0, []edge.Edge{edges[0], edges[3]}}, {1, []edge.Edge{edges[0], edges[1]}}, {2, []edge.Edge{edges[1], edges[2]}}, {3, []edge.Edge{edges[2], edges[3]}}, {4, []edge.Edge{edges[4], edges[4]}}},
		wantDegree: []VertexDegree{{0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2}},
		allEdges:   []edge.Edge{edges[0], edges[3], edges[1], edges[2], edges[4]},
	},
}

func hasSameEdges(got, want []edge.Edge) bool {
	if len(got) != len(want) {
		return false
	}

	for i, w := range want {
		v1 := w.Either()
		w1 := w.Other(v1)
		v2 := got[i].Either()
		w2 := got[i].Other(v2)
		if v1 != v2 || w1 != w2 || w.Weight() != got[i].Weight() {
			return false
		}
	}
	return true
}

func TestEdgeWeightedGraph(t *testing.T) {
	for _, tc := range testCases {
		g := NewEdgeWeightedGraph(tc.size)
		for _, e := range tc.edges {
			g.AddEdge(e)
		}

		for _, wantEdge := range tc.wantEdges {
			if !hasSameEdges(g.AdjacencyList(wantEdge.vertex), wantEdge.adj) {
				t.Errorf("Adjacency list of vertex %v do not include all necessary edges", wantEdge.vertex)
			}
		}

		for _, degree := range tc.wantDegree {
			if g.Degree(degree.vertex) != degree.degree {
				t.Errorf("Got degree of %v for vertex %v, want %v", g.Degree(degree.vertex), degree.vertex, degree.degree)
			}
		}

		if !hasSameEdges(tc.allEdges, g.Edges()) {
			t.Errorf("Got incomplete slice of edges, want slice of all edges in given edge weighted graph")
		}
	}
}
