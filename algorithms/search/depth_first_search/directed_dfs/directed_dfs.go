package directed_dfs

import "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/directed_graph"

type DirectedDFS struct {
	marked []bool
	count  int
}

func NewDirectedDFS(g *directed_graph.Digraph, sourceVertices ...int) *DirectedDFS {
	dfs := DirectedDFS{make([]bool, g.Vertices()), 0}
	dfs.validateVertices(sourceVertices)
	for _, s := range sourceVertices {
		if !dfs.marked[s] {
			dfs.dfs(g, s)
		}
	}
	return &dfs
}

func (dfs *DirectedDFS) dfs(g *directed_graph.Digraph, v int) {
	dfs.marked[v] = true
	dfs.count++
	for _, w := range g.AdjacencyList(v) {
		if !dfs.marked[w] {
			dfs.dfs(g, w)
		}
	}
}

func (dfs *DirectedDFS) IsConnected(v int) bool {
	dfs.validateVertex(v)
	return dfs.marked[v]
}

func (dfs *DirectedDFS) Count() int {
	return dfs.count
}

func (dfs *DirectedDFS) validateVertex(v int) {
	if v < 0 || v >= len(dfs.marked) {
		panic("Vertex out of bounds")
	}
}

func (dfs *DirectedDFS) validateVertices(vv []int) {
	for _, v := range vv {
		dfs.validateVertex(v)
	}
}
