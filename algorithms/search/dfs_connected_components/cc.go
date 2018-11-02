package dfs_connected_components

import "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/graph"

type CC struct {
	id      []int
	countId []int
	marked  []bool
	countCC int
}

func NewCC(g *graph.Graph) *CC {
	size := g.Vertices() + 1
	cc := CC{make([]int, size), make([]int, size), make([]bool, size), 0}
	for v := 0; v < g.Vertices(); v++ {
		if !cc.marked[v] {
			cc.countCC++
			cc.dfs(g, v)
		}
	}
	return &cc
}

func (cc *CC) dfs(g *graph.Graph, v int) {
	cc.marked[v] = true
	cc.id[v] = cc.countCC
	cc.countId[cc.countCC]++
	for _, w := range g.AdjacencyList(v) {
		if !cc.marked[w] {
			cc.dfs(g, w)
		}
	}
}

func (cc *CC) Connected(v, w int) bool {
	cc.validateVertices(v, w)
	return cc.id[v] == cc.id[w]
}

func (cc *CC) CountCC() int {
	return cc.countCC
}

func (cc *CC) Id(v int) int {
	cc.validateVertex(v)
	return cc.id[v]
}

func (cc *CC) CountVerticesWithId(i int) int {
	cc.validateId(i)
	return cc.countId[i]
}

func (cc *CC) validateVertex(v int) {
	if v < 0 || v > len(cc.marked) {
		panic("Vertex out of bounds")
	}
}

func (cc *CC) validateVertices(vv ...int) {
	for _, v := range vv {
		cc.validateVertex(v)
	}
}

func (cc *CC) validateId(i int) {
	cc.validateVertex(i)
}
