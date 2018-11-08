package depth_first_directed_paths

import (
	"errors"
	"fmt"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/directed_graph"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/stack"
)

type DepthFirstDirectedPaths struct {
	edgeTo       []int
	marked       []bool
	sourceVertex int
}

func NewDFDP(g *directed_graph.Digraph, sourceVertex int) *DepthFirstDirectedPaths {
	dfdp := DepthFirstDirectedPaths{make([]int, g.Vertices()), make([]bool, g.Vertices()), sourceVertex}
	dfdp.validateVertex(sourceVertex)
	dfdp.dfs(g, sourceVertex)
	return &dfdp
}

func (dfdp *DepthFirstDirectedPaths) dfs(g *directed_graph.Digraph, v int) {
	dfdp.marked[v] = true
	for _, w := range g.AdjacencyList(v) {
		if !dfdp.marked[w] {
			dfdp.edgeTo[w] = v
			dfdp.dfs(g, w)
		}
	}
}

func (dfdp *DepthFirstDirectedPaths) HasPathTo(v int) bool {
	dfdp.validateVertex(v)
	return dfdp.marked[v]
}

func (dfdp *DepthFirstDirectedPaths) PathTo(v int) (stack.Stack, error) {
	dfdp.validateVertex(v)
	stack := stack.NewEmptyStack()

	if !dfdp.HasPathTo(v) {
		return stack, errors.New(fmt.Sprintf("No directed path found from %v to %v\n", dfdp.sourceVertex, v))
	}

	for x := v; x != dfdp.sourceVertex; x = dfdp.edgeTo[x] {
		stack.Push(x)
	}
	stack.Push(dfdp.sourceVertex)
	return stack, nil
}

func (dfdp *DepthFirstDirectedPaths) validateVertex(v int) {
	if v < 0 || v >= len(dfdp.marked) {
		panic("Vertex out of bounds")
	}
}
