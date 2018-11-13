// Transitive Closure computes the transitive closure of a digraph.
// It is acceptable for small digraphs, but not so for larger ones
// since it has a time complexity of O(V(V+E)) and a space complexity of 0(V^2)

package transitive_closure

import (
	"github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/graph/depth_first_search/directed_dfs"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/directed_graph"
)

type TransitiveClosure struct {
	tc []*directed_dfs.DirectedDFS
}

func NewTransitiveClosure(g *directed_graph.Digraph) *TransitiveClosure {
	tc := TransitiveClosure{make([]*directed_dfs.DirectedDFS, g.Vertices())}
	for v := 0; v < g.Vertices(); v++ {
		tc.tc[v] = directed_dfs.NewDirectedDFS(g, v)
	}
	return &tc
}

func (tc *TransitiveClosure) Reachable(v, w int) bool {
	tc.validateVertices(v, w)
	return tc.tc[v].IsConnected(w)
}

func (tc *TransitiveClosure) validateVertices(vv ...int) {
	for _, v := range vv {
		tc.validateVertex(v)
	}
}

func (tc *TransitiveClosure) validateVertex(v int) {
	if v < 0 || v >= len(tc.tc) {
		panic("Vertex out of bounds")
	}
}
