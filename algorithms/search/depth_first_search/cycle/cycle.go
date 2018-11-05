// dfs_cycle finds a simple cycle in an undirected graph
package dfs_cycle

import (
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/graph"
)

// Cycle holds information about a possibly existing simple cycle
type Cycle struct {
	marked []bool
	edgeTo []int
	cycle  *Bag
}

type Bag struct {
	items []int
	size  int
}

func (b *Bag) Add(i int) {
	b.items = append(b.items, i)
	b.size++
}

func NewDFSCyle(g *graph.Graph) *Cycle {
	size := g.Vertices()
	c := &Cycle{make([]bool, size), make([]int, size), nil}

	if c.HasSelfLoop(g) {
		return c
	}

	if c.HasParallelEdges(g) {
		return c
	}

	for v := 0; v < g.Vertices(); v++ {
		if !c.marked[v] {
			c.dfs(g, -1, v)
		}
	}
	return c
}

func (c *Cycle) dfs(g *graph.Graph, u, v int) {
	c.marked[v] = true
	for _, w := range g.AdjacencyList(v) {
		if c.HasCycle() {
			return
		}

		if !c.marked[w] {
			c.edgeTo[w] = v
			c.dfs(g, v, w)
		} else if w != u {
			c.initCycle()
			for x := v; x != w; x = c.edgeTo[x] {
				c.cycle.Add(x)
			}
			c.cycle.Add(w)
			c.cycle.Add(v)
		}
	}

}

func (c *Cycle) HasSelfLoop(g *graph.Graph) bool {
	for v := 0; v < g.Vertices(); v++ {
		for _, w := range g.AdjacencyList(v) {
			if v == w {
				c.initCycle()
				c.cycle.Add(v)
				c.cycle.Add(v)
				return true
			}
		}
	}
	return false
}

func (c *Cycle) HasParallelEdges(g *graph.Graph) bool {
	for v := 0; v < g.Vertices(); v++ {
		for _, w := range g.AdjacencyList(v) {
			if c.marked[w] {
				c.initCycle()
				c.cycle.Add(v)
				c.cycle.Add(w)
				c.cycle.Add(v)
				return true
			}

			c.marked[w] = true
		}

		for _, w := range g.AdjacencyList(v) {
			c.marked[w] = false
		}
	}
	return false
}

func (c *Cycle) HasCycle() bool {
	return c.cycle != nil
}

func (c *Cycle) Cycle() []int {
	if c.HasCycle() {
		return c.cycle.items
	}
	return nil
}

func (c *Cycle) initCycle() {
	b := &Bag{make([]int, 0, len(c.marked)), 0}
	c.cycle = b
}
