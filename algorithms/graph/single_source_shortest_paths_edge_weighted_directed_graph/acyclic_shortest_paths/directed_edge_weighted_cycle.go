package acyclic_shortest_paths

import (
	graph "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_directed_graph"
)

type DirectedEdgeWeightedCycle struct {
	marked  []bool
	edgeTo  []graph.DirectedEdge
	onStack []bool
	cycle   graph.Stack
}

func NewDirectedEdgeWeightedCycle(g *graph.EdgeWeightedDigraph) *DirectedEdgeWeightedCycle {
	size := g.VerticesCount()
	dc := &DirectedEdgeWeightedCycle{make([]bool, size), make([]graph.DirectedEdge, size), make([]bool, size), graph.Stack{}}
	for v := 0; v < size; v++ {
		if dc.HasCycle() {
			break
		}
		if !dc.marked[v] {
			dc.dfs(g, v)
		}
	}
	return dc
}

func (dc *DirectedEdgeWeightedCycle) dfs(g *graph.EdgeWeightedDigraph, v int) {
	dc.marked[v] = true
	dc.onStack[v] = true
	for _, e := range g.AdjacencyList(v) {
		w := e.To
		if dc.HasCycle() {
			return
		} else if !dc.marked[w] {
			dc.edgeTo[w] = e
			dc.dfs(g, w)
		} else if dc.onStack[w] {
			curr := graph.DirectedEdge{}
			for curr = e; curr.From != w; curr = dc.edgeTo[curr.From] {
				dc.cycle.Push(curr)
			}
			dc.cycle.Push(curr)
		}
	}
	dc.onStack[v] = false
}

func (dc *DirectedEdgeWeightedCycle) HasCycle() bool {
	return dc.cycle.Size > 0
}

func (dc *DirectedEdgeWeightedCycle) Cycle() graph.Stack {
	return dc.cycle
}
