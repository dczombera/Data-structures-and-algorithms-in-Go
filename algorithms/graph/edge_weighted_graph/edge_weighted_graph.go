package edge_weighted_graph

import "github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/graph/edge"

type EdgeWeightedGraph struct {
	vertices int
	edges    int
	adj      [][]edge.Edge
}

func NewEdgeWeightedGraph(sizeV int) *EdgeWeightedGraph {
	return &EdgeWeightedGraph{sizeV, 0, make([][]edge.Edge, sizeV)}
}

func (g *EdgeWeightedGraph) AddEdge(e edge.Edge) {
	v := e.Either()
	w := e.Other(v)
	g.validateVertices(v, w)
	g.adj[v] = append(g.adj[v], e)
	g.adj[w] = append(g.adj[w], e)
	g.edges++
}

func (g *EdgeWeightedGraph) VerticesCount() int {
	return g.vertices
}

func (g *EdgeWeightedGraph) EdgesCount() int {
	return g.edges
}

func (g *EdgeWeightedGraph) AdjacencyList(v int) []edge.Edge {
	g.validateVertex(v)
	return g.adj[v]
}

func (g *EdgeWeightedGraph) Edges() []edge.Edge {
	edges := []edge.Edge{}
	for v := 0; v < g.VerticesCount(); v++ {
		selfLoops := 0
		for _, e := range g.adj[v] {
			if e.Other(v) > v {
				edges = append(edges, e)
			} else if v == e.Other(v) {
				// Append only one copy of each self loop
				if selfLoops%2 == 0 {
					edges = append(edges, e)
				}
				selfLoops++
			}
		}
	}
	return edges
}

func (g *EdgeWeightedGraph) Degree(v int) int {
	g.validateVertex(v)
	return len(g.adj[v])
}

func (g *EdgeWeightedGraph) validateVertex(v int) {
	if v < 0 || v >= g.vertices {
		panic("Vertex out of bounds")
	}
}

func (g *EdgeWeightedGraph) validateVertices(vv ...int) {
	for _, v := range vv {
		g.validateVertex(v)
	}
}
