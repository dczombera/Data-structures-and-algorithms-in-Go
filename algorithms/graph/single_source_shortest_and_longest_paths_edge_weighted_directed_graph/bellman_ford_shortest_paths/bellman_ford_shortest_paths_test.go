package bellman_ford_shortest_paths

import (
	"testing"

	graph "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_directed_graph"
)

type shortestPath struct {
	to     int
	distTo float32
	path   []graph.DirectedEdge
}

type TestCase struct {
	graphSize    int
	edges        []graph.DirectedEdge
	sourceVertex int
	paths        []shortestPath
}

func TestBellmanFordShortestPaths(t *testing.T) {
	tc := TestCase{
		graphSize:    8,
		edges:        []graph.DirectedEdge{{4, 5, 0.35}, {5, 4, 0.35}, {4, 7, 0.37}, {5, 7, 0.28}, {7, 5, 0.28}, {5, 1, 0.32}, {0, 4, 0.38}, {0, 2, 0.26}, {7, 3, 0.39}, {1, 3, 0.29}, {2, 7, 0.34}, {6, 2, -1.20}, {3, 6, 0.52}, {6, 0, -1.40}, {6, 4, -1.25}},
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
				distTo: 0.61,
				path:   []graph.DirectedEdge{{0, 2, 0.26}, {2, 7, 0.34}, {7, 3, 0.39}, {3, 6, 0.52}, {6, 4, -1.25}, {4, 5, 0.35}},
			},
			{
				to:     1,
				distTo: 0.93,
				path:   []graph.DirectedEdge{{0, 2, 0.26}, {2, 7, 0.34}, {7, 3, 0.39}, {3, 6, 0.52}, {6, 4, -1.25}, {4, 5, 0.35}, {5, 1, 0.32}},
			},
		},
	}
	g := graph.NewEdgeWeightedDigraph(tc.graphSize)

	for _, e := range tc.edges {
		g.AddEdge(e)
	}

	sp := NewBellmanFordSP(g, tc.sourceVertex)

	for _, p := range tc.paths {
		if !sp.HasPathTo(p.to) {
			t.Errorf("Got no path from source vertex %v to vertex %v, want shortest path", tc.sourceVertex, p.to)
		}

		dist, err := sp.DistTo(p.to)
		if err != nil {
			t.Errorf("Got error %v for distance from %v to %v, want %v", err, tc.sourceVertex, p.to, p.distTo)
		}

		if float32(dist) != p.distTo {
			t.Errorf("Got a distance of %v from %v to %v, want distance %v", dist, tc.sourceVertex, p.to, p.distTo)
		}

		path, err := sp.PathTo(p.to)
		if err != nil {
			t.Errorf("Got error %v, want path from %v to %v", err, tc.sourceVertex, p.to)
		}
		curr := path.First
		for _, e := range p.path {
			if curr.Item != e {
				t.Errorf("Got directed edge %v in path from %v to %v, want %v", curr.Item, tc.sourceVertex, p.to, e)
			}
			curr = curr.Next
		}
	}
}

func TestBellmanFordShortestPathsNegativeCycle(t *testing.T) {
	tc := struct {
		graphSize    int
		edges        []graph.DirectedEdge
		sourceVertex int
		cycle        []graph.DirectedEdge
	}{
		graphSize:    8,
		edges:        []graph.DirectedEdge{{4, 5, 0.35}, {5, 4, -0.66}, {4, 7, 0.37}, {5, 7, 0.28}, {7, 5, 0.28}, {5, 1, 0.32}, {0, 4, 0.38}, {0, 2, 0.26}, {7, 3, 0.39}, {1, 3, 0.29}, {2, 7, 0.34}, {6, 2, 0.40}, {3, 6, 0.52}, {6, 0, 0.58}, {6, 4, 0.93}},
		sourceVertex: 0,
		cycle:        []graph.DirectedEdge{{4, 5, 0.35}, {5, 4, -0.66}},
	}

	g := graph.NewEdgeWeightedDigraph(tc.graphSize)
	for _, e := range tc.edges {
		g.AddEdge(e)
	}
	sp := NewBellmanFordSP(g, tc.sourceVertex)
	if !sp.HasNegativeCycle() {
		t.Errorf("Did not find negative cycle, even though edge-weighted digraph has negative cycle 4 -> 5 -> 4\n")
	}
	curr := sp.NegativeCycle().First
	for _, e := range tc.cycle {
		if curr.Item != e {
			t.Errorf("Got edge %v in negative cycle, want %v\n", curr.Item, e)
		}
		curr = curr.Next
	}
}
