package dfp

import (
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/graph"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/stack"
)

type DepthFirstPaths struct {
	edgeTo       []int
	connected    []bool
	sourceVertex int
	count        int
}

func NewDFP(g *graph.Graph, sourceVertex int) *DepthFirstPaths {
	size := g.Vertices()
	dfp := DepthFirstPaths{make([]int, size), make([]bool, size), sourceVertex, 0}
	dfp.validateVertex(sourceVertex)
	dfp.dfs(g, sourceVertex)
	return &dfp
}

func (dfp *DepthFirstPaths) dfs(g *graph.Graph, v int) {
	dfp.connected[v] = true
	dfp.count++
	for _, w := range g.AdjacencyList(v) {
		if !dfp.connected[w] {
			dfp.edgeTo[w] = v
			dfp.dfs(g, w)
		}
	}
}

func (dfp *DepthFirstPaths) PathTo(v int) *stack.Stack {
	dfp.validateVertex(v)
	if !dfp.HasPathTo(v) {
		return nil
	}

	stack := stack.NewEmptyStack()
	for i := v; i != dfp.sourceVertex; i = dfp.edgeTo[i] {
		stack.Push(i)
	}

	stack.Push(dfp.sourceVertex)
	return &stack
}

func (dfp *DepthFirstPaths) HasPathTo(v int) bool {
	dfp.validateVertex(v)
	return dfp.connected[v]
}

func (dfp *DepthFirstPaths) Count() int {
	return dfp.count
}

func (dfp *DepthFirstPaths) validateVertex(v int) {
	if v < 0 || v >= len(dfp.connected) {
		panic("Vertex out of bounds")
	}
}
