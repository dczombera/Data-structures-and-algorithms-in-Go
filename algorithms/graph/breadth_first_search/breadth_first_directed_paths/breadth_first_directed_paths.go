package breadth_first_directed_paths

import (
	"errors"
	"fmt"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/directed_graph"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/stack"
)

type BreadthFirstDirectedPaths struct {
	edgeTo       []int
	marked       []bool
	distTo       []int
	sourceVertex int
}

func NewBFDP(g *directed_graph.Digraph, sourceVertex int) *BreadthFirstDirectedPaths {
	length := g.Vertices()
	bfdp := BreadthFirstDirectedPaths{make([]int, length), make([]bool, length), make([]int, length), sourceVertex}
	bfdp.validateVertex(sourceVertex)
	bfdp.bfs(g, sourceVertex)

	return &bfdp
}

// TODO: Add constructor for multiple source vertices

func (bfdp *BreadthFirstDirectedPaths) bfs(g *directed_graph.Digraph, v int) {
	bfdp.marked[v] = true
	bfdp.distTo[v] = 0
	bfdp.edgeTo[v] = v
	q := []int{v}
	for len(q) > 0 {
		v, q = q[0], q[1:]
		for _, w := range g.AdjacencyList(v) {
			if !bfdp.marked[w] {
				bfdp.marked[w] = true
				bfdp.edgeTo[w] = v
				bfdp.distTo[w] = bfdp.distTo[v] + 1
				q = append(q, w)
			}
		}
	}
}

func (bfdp *BreadthFirstDirectedPaths) HasPathTo(v int) bool {
	bfdp.validateVertex(v)
	return bfdp.marked[v]
}

func (bfdp *BreadthFirstDirectedPaths) PathTo(v int) (stack.Stack, error) {
	bfdp.validateVertex(v)
	stack := stack.NewEmptyStack()

	if !bfdp.HasPathTo(v) {
		return stack, errors.New(fmt.Sprintf("No directed shortest path found from %v to %v\n", bfdp.sourceVertex, v))
	}

	for x := v; bfdp.distTo[x] > 0; x = bfdp.edgeTo[x] {
		stack.Push(x)
	}
	stack.Push(bfdp.sourceVertex)

	return stack, nil
}

func (bfdp *BreadthFirstDirectedPaths) DistTo(v int) (int, error) {
	bfdp.validateVertex(v)
	if !bfdp.HasPathTo(v) {
		return -1, errors.New(fmt.Sprintf("No path between %v and %v", bfdp.sourceVertex, v))
	}

	return bfdp.distTo[v], nil
}

func (bfdp *BreadthFirstDirectedPaths) validateVertex(v int) {
	if v < 0 || v >= len(bfdp.marked) {
		panic("Vertex out of bounds")
	}
}
