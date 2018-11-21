package acyclic_shortest_paths

import (
	"testing"

	graph "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_directed_graph"
)

type shortestPath struct {
	to     int
	distTo float32
	path   []graph.DirectedEdge
}

type testCaseAcyclicSP struct {
	graphSize    int
	edges        []graph.DirectedEdge
	sourceVertex int
	paths        []shortestPath
}

func TestAcyclicShortestPaths(t *testing.T) {
	var testCases = []testCaseAcyclicSP{
		{
			graphSize:    5,
			edges:        []graph.DirectedEdge{{0, 1, 0.1}, {1, 2, 0.1}, {2, 3, 0.1}, {1, 3, 0.4}, {3, 4, 0.1}, {0, 4, 0.2}},
			sourceVertex: 0,
			paths: []shortestPath{
				{
					to:     1,
					distTo: 0.1,
					path:   []graph.DirectedEdge{{0, 1, 0.1}},
				},
				{
					to:     2,
					distTo: 0.2,
					path:   []graph.DirectedEdge{{0, 1, 0.1}, {1, 2, 0.1}},
				},
				{
					to:     3,
					distTo: 0.3,
					path:   []graph.DirectedEdge{{0, 1, 0.1}, {1, 2, 0.1}, {2, 3, 0.1}},
				},
				{
					to:     4,
					distTo: 0.2,
					path:   []graph.DirectedEdge{{0, 4, 0.2}},
				},
			},
		},
	}

	for _, tc := range testCases {
		g := graph.NewEdgeWeightedDigraph(tc.graphSize)

		for _, e := range tc.edges {
			g.AddEdge(e)
		}

		asp, err := NewAcyclicSP(g, tc.sourceVertex)
		if err != nil {
			t.Error(err)
		}
		for _, p := range tc.paths {
			if !asp.HasPathTo(p.to) {
				t.Errorf("Got no path from source vertex %v to vertex %v, want shortest path", tc.sourceVertex, p.to)
			}

			dist, err := asp.DistTo(p.to)
			if err != nil {
				t.Errorf("Got error %v for distance from %v to %v, want %v", err, tc.sourceVertex, p.to, p.distTo)
			}

			if float32(dist) != p.distTo {
				t.Errorf("Got a distance of %v from %v to %v, want distance %v", dist, tc.sourceVertex, p.to, p.distTo)
			}

			path, err := asp.PathTo(p.to)
			if err != nil {
				t.Errorf("Got error %v, want path from %v to %v", err, tc.sourceVertex, p.to)
			}
			curr := path.First
			for _, e := range p.path {
				if curr.Item != e {
					t.Errorf("Got directed edge %v, want %v", curr.Item, e)
				}
				curr = curr.Next
			}
		}
	}
}

func TestAcyclicShortestPathsWithCycle(t *testing.T) {
	var testCases = []testCaseAcyclicSP{
		{
			graphSize:    8,
			edges:        []graph.DirectedEdge{{4, 5, 0.35}, {5, 4, 0.35}, {4, 7, 0.37}, {5, 7, 0.28}, {7, 5, 0.28}, {5, 1, 0.32}, {0, 4, 0.38}, {0, 2, 0.26}, {7, 3, 0.39}, {1, 3, 0.29}, {2, 7, 0.34}, {6, 2, 0.40}, {3, 6, 0.52}, {6, 0, 0.58}, {6, 4, 0.93}},
			sourceVertex: 0,
			paths: []shortestPath{
				{
					to:     2,
					distTo: 0.26,
					path:   []graph.DirectedEdge{{0, 2, 0.26}},
				},
				{
					to:     6,
					distTo: 1.51,
					path:   []graph.DirectedEdge{{0, 2, 0.26}, {2, 7, 0.34}, {7, 3, 0.39}, {3, 6, 0.52}},
				},
				{
					to:     5,
					distTo: 0.73,
					path:   []graph.DirectedEdge{{0, 4, 0.38}, {4, 5, 0.35}},
				},
				{
					to:     1,
					distTo: 1.05,
					path:   []graph.DirectedEdge{{0, 4, 0.38}, {4, 5, 0.35}, {5, 1, 0.32}},
				},
			},
		},
	}

	for _, tc := range testCases {
		g := graph.NewEdgeWeightedDigraph(tc.graphSize)

		for _, e := range tc.edges {
			g.AddEdge(e)
		}

		_, err := NewAcyclicSP(g, tc.sourceVertex)
		if err == nil {
			t.Error("Got new acyclic shortest path data type, want error due to cycle in graph")
		}
	}
}
