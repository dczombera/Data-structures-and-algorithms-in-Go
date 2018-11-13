// bipartite either finds a bipartition or an odd-length cycle of an undirected graph
package bipartite

import (
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/graph"
)

type Bipartite struct {
	marked      []bool
	color       []bool
	edgeTo      []int
	oddCycle    *Bag
	isBipartite bool
}

type Bag struct {
	items []int
}

func (b *Bag) Add(v int) {
	b.items = append(b.items, v)
}

func NewBipartite(g *graph.Graph) *Bipartite {
	size := g.Vertices()
	bp := Bipartite{make([]bool, size), make([]bool, size), make([]int, size), nil, true}
	for v := 0; v < g.Vertices(); v++ {
		if !bp.marked[v] {
			bp.dfs(g, v)
		}
	}

	return &bp
}

func (bp *Bipartite) dfs(g *graph.Graph, v int) {
	bp.marked[v] = true
	for _, w := range g.AdjacencyList(v) {
		if bp.HasOddCycle() {
			return
		}

		if !bp.marked[w] {
			bp.edgeTo[w] = v
			bp.color[w] = !bp.color[v]
			bp.dfs(g, w)
		}

		// check for odd-length cycle
		if bp.color[v] == bp.color[w] {
			bp.initCycle()
			bp.oddCycle.Add(w)
			for x := v; x != w; x = bp.edgeTo[x] {
				bp.oddCycle.Add(x)
			}
			bp.oddCycle.Add(w)
		}
	}
}

func (bp *Bipartite) OddCycle() []int {
	if bp.HasOddCycle() {
		return bp.oddCycle.items
	}
	return nil
}

func (bp *Bipartite) HasOddCycle() bool {
	return bp.oddCycle != nil
}

func (bp *Bipartite) Color(v int) bool {
	bp.validateVertex(v)
	return bp.color[v]
}

func (bp *Bipartite) IsBipartite() bool {
	return bp.isBipartite
}

func (bp *Bipartite) initCycle() {
	b := Bag{make([]int, 0, len(bp.color))}
	bp.oddCycle = &b
}

func (bp *Bipartite) validateVertex(v int) {
	if v < 0 || v > len(bp.marked) {
		panic("Vertex out of bounds")
	}
}
