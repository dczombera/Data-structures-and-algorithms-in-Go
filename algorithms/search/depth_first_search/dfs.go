package dfs

import "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/graph"

type DepthFirstSearch struct {
	connected []bool
	count     int
}

func NewDFS(g *graph.Graph, sourceVertex int) *DepthFirstSearch {
	dfs := DepthFirstSearch{make([]bool, g.Vertices()), 0}
	dfs.validateVertex(sourceVertex)
	dfs.dfs(g, sourceVertex)
	return &dfs
}

func (dfs *DepthFirstSearch) dfs(g *graph.Graph, vertex int) {
	dfs.connected[vertex] = true
	dfs.count++
	for _, w := range g.AdjacencyList(vertex) {
		if !dfs.connected[w] {
			dfs.dfs(g, w)
		}
	}
}

func (dfs *DepthFirstSearch) IsConnected(vertex int) bool {
	dfs.validateVertex(vertex)
	return dfs.connected[vertex]
}

func (dfs *DepthFirstSearch) Count() int {
	return dfs.count
}

func (dfs *DepthFirstSearch) validateVertex(v int) {
	if v < 0 || v >= len(dfs.connected) {
		panic("Vertex out of bounds")
	}
}
