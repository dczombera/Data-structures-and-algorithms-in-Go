package graph

import (
	"fmt"
)

type Graph struct {
	vertices  int
	edges     int
	adjacency [][]int
}

func NewGraph(sizeV int) Graph {
	return Graph{sizeV, 0, make([][]int, sizeV)}
}

func (g *Graph) AddEdge(v, w int) {
	g.validateVertices(v, w)
	g.adjacency[v] = append(g.adjacency[v], w)
	g.adjacency[w] = append(g.adjacency[w], v)
}

func (g *Graph) Vertices() int {
	return g.vertices
}

func (g *Graph) Edges() int {
	return g.edges
}

func (g *Graph) AdjacencyList(vertex int) []int {
	g.validateVertices(vertex)
	return g.adjacency[vertex]
}

func (g *Graph) Degree(vertex int) int {
	g.validateVertices(vertex)
	return len(g.adjacency[vertex])
}

func (g *Graph) validateVertices(vertices ...int) {
	for _, v := range vertices {
		if v < 0 || v >= g.vertices {
			panic(fmt.Sprintf("vertix %v out of graph bounds", v))
		}
	}
}
