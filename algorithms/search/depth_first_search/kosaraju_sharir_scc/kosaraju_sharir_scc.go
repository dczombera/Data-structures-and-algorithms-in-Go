// This implementation uses the Kosaraju Sharir Algorithm to find the strong components in a digraph.

package kosaraju_sharir_scc

import (
	"github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/search/depth_first_search/depth_first_order"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/directed_graph"
)

type KosarajuSharirSCC struct {
	marked []bool
	id     []int
	count  int
}

func NewKosarajuSharirSCC(g *directed_graph.Digraph) *KosarajuSharirSCC {
	scc := KosarajuSharirSCC{make([]bool, g.Vertices()), make([]int, g.Vertices()), 0}
	// It is essential to get the reverse post order of the reverse digraph
	// in order to find all strong components in a digraph
	reverse := depth_first_order.NewDepthFirstOrder(g.Reverse()).ReversePostOrder()

	for curr := reverse.First; curr != nil; curr = curr.Next {
		if !scc.marked[curr.Item] {
			scc.dfs(g, curr.Item)
			scc.count++
		}
	}

	return &scc
}

func (scc *KosarajuSharirSCC) dfs(g *directed_graph.Digraph, v int) {
	scc.marked[v] = true
	scc.id[v] = scc.count
	for _, w := range g.AdjacencyList(v) {
		if !scc.marked[w] {
			scc.dfs(g, w)
		}
	}
}

// StronglyConnected finds out wheter to vertices are strongly connected.
// Two vertices are strongly connected, if and only if
// there is a directed path from v to w and vice versa
func (scc *KosarajuSharirSCC) StronglyConnected(v, w int) bool {
	scc.validateVertices(v, w)
	return scc.id[v] == scc.id[w]
}

func (scc *KosarajuSharirSCC) Id(v int) int {
	scc.validateVertex(v)
	return scc.id[v]
}

func (scc *KosarajuSharirSCC) Count() int {
	return scc.count
}

func (scc *KosarajuSharirSCC) validateVertices(vv ...int) {
	for _, v := range vv {
		scc.validateVertex(v)
	}
}

func (scc *KosarajuSharirSCC) validateVertex(v int) {
	if v < 0 || v >= len(scc.marked) {
		panic("Vertex out of bounds")
	}
}
