package edge_weighted_directed_graph

import "log"

type EdgeWeightedDigraph struct {
	vertices int
	edges    int
	adj      [][]DirectedEdge
	indegree []int
}

func NewEdgeWeightedDigraph(sizeV int) *EdgeWeightedDigraph {
	return &EdgeWeightedDigraph{sizeV, 0, make([][]DirectedEdge, sizeV), make([]int, sizeV)}
}

func (g *EdgeWeightedDigraph) VerticesCount() int {
	return g.vertices
}

func (g *EdgeWeightedDigraph) EdgesCount() int {
	return g.edges
}

func (g *EdgeWeightedDigraph) AddEdge(e DirectedEdge) {
	g.validateEdge(e)
	g.adj[e.From()] = append(g.adj[e.From()], e)
	g.indegree[e.To()]++
	g.edges++
}

func (g *EdgeWeightedDigraph) AdjacencyList(v int) []DirectedEdge {
	g.validateVertex(v)
	return g.adj[v]
}

func (g *EdgeWeightedDigraph) Outdegree(v int) int {
	g.validateVertex(v)
	return len(g.adj[v])
}

func (g *EdgeWeightedDigraph) Indegree(v int) int {
	g.validateVertex(v)
	return g.indegree[v]
}

func (g *EdgeWeightedDigraph) Edges() []DirectedEdge {
	edges := make([]DirectedEdge, 0, g.EdgesCount())
	for v := 0; v < g.VerticesCount(); v++ {
		for _, e := range g.adj[v] {
			edges = append(edges, e)
		}
	}
	return edges
}

func (g *EdgeWeightedDigraph) validateEdge(e DirectedEdge) {
	g.validateVertex(e.From())
	g.validateVertex(e.To())
}

func (g *EdgeWeightedDigraph) validateVertex(v int) {
	if v < 0 || v >= g.VerticesCount() {
		log.Fatalf("Vertex %v out of bounds", v)
	}
}
