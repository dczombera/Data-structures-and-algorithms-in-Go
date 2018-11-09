package directed_cycle

import (
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/directed_graph"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/stack"
)

type DirectedCycle struct {
	edgeTo  []int
	marked  []bool
	onStack []bool
	cycle   *stack.Stack
}

func NewDirectedCycle(g *directed_graph.Digraph) *DirectedCycle {
	length := g.Vertices()
	dc := DirectedCycle{make([]int, length), make([]bool, length), make([]bool, length), nil}
	for v := 0; v < length; v++ {
		if !dc.marked[v] {
			dc.dfs(g, v)
		}
	}
	return &dc
}

func (dc *DirectedCycle) dfs(g *directed_graph.Digraph, v int) {
	dc.onStack[v] = true
	dc.marked[v] = true
	for _, w := range g.AdjacencyList(v) {
		if dc.HasCycle() {
			return
		}

		if !dc.marked[w] {
			dc.edgeTo[w] = v
			dc.dfs(g, w)
		}

		if dc.onStack[w] {
			dc.cycle = stack.NewStack()
			for x := v; x != w; x = dc.edgeTo[v] {
				dc.cycle.Push(x)
			}
			dc.cycle.Push(w)
			dc.cycle.Push(v)
		}
	}
	dc.onStack[v] = false
}

func (dc *DirectedCycle) HasCycle() bool {
	return dc.cycle != nil
}

func (dc *DirectedCycle) Cycle() *stack.Stack {
	return dc.cycle
}
