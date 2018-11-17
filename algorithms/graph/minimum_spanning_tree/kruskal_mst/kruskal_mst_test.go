package kruskal_mst

import (
	"testing"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_graph"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_graph/edge"
)

type testCase struct {
	graphSize int
	edges     []edge.Edge
	mst       []edge.Edge
	weight    float64
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

func TestKruskalMST(t *testing.T) {
	var testCases = []testCase{
		{
			graphSize: 8,
			edges: []edge.Edge{edge.NewEdge(4, 5, 0.35), edge.NewEdge(4, 7, 0.37), edge.NewEdge(5, 7, 0.28), edge.NewEdge(0, 7, 0.16),
				edge.NewEdge(1, 5, 0.32), edge.NewEdge(0, 4, 0.38), edge.NewEdge(2, 3, 0.17), edge.NewEdge(1, 7, 0.19), edge.NewEdge(0, 2, 0.26),
				edge.NewEdge(1, 2, 0.36), edge.NewEdge(1, 3, 0.29), edge.NewEdge(2, 7, 0.34), edge.NewEdge(6, 2, 0.40), edge.NewEdge(3, 6, 0.52),
				edge.NewEdge(6, 0, 0.58), edge.NewEdge(6, 4, 0.93)},
			mst:    []edge.Edge{edge.NewEdge(0, 7, 0.16), edge.NewEdge(2, 3, 0.17), edge.NewEdge(1, 7, 0.19), edge.NewEdge(0, 2, 0.26), edge.NewEdge(5, 7, 0.28), edge.NewEdge(4, 5, 0.35), edge.NewEdge(6, 2, 0.40)},
			weight: 1.81,
		},
	}

	for _, tc := range testCases {
		g := edge_weighted_graph.NewEdgeWeightedGraph(tc.graphSize)
		for _, e := range tc.edges {
			g.AddEdge(e)
		}

		mst := NewKruskalMST(g)
		queue := mst.Edges()
		edges := make([]edge.Edge, 0, g.VerticesCount()-1)
		for curr := queue.First; curr != nil; curr = curr.Next {
			edges = append(edges, curr.Item)
		}

		if !hasSameEdges(tc.mst, edges) {
			t.Errorf("Got minimum spanning tree of %v, want %v", edges, tc.mst)
		}

		if tc.weight != mst.Weight() {
			t.Errorf("Got minimum spanning tree weight of %v, want %v", mst.Weight(), tc.weight)
		}
	}
}
