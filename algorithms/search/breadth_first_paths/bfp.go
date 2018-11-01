package breadth_first_paths

import (
	"errors"
	"fmt"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/graph"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/queue"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/stack"
)

type BreadthFirstPaths struct {
	marked []bool
	edgeTo []int
	distTo []int
	source int
}

func NewBreadthFirstPaths(g *graph.Graph, sourceVertex int) *BreadthFirstPaths {
	size := g.Vertices()
	bfp := BreadthFirstPaths{make([]bool, size), make([]int, size), make([]int, size), sourceVertex}
	bfp.bfs(g, sourceVertex)
	return &bfp
}

func (bfp *BreadthFirstPaths) bfs(g *graph.Graph, s int) {
	bfp.validateVertex(s)
	q := queue.NewEmptyQueue()
	bfp.marked[s] = true
	bfp.distTo[s] = 0
	q.Push(s)
	for q.Size > 0 {
		v, _ := q.Pop()
		for _, w := range g.AdjacencyList(v) {
			if !bfp.marked[w] {
				bfp.marked[w] = true
				bfp.edgeTo[w] = v
				bfp.distTo[w] = bfp.distTo[v] + 1
				q.Push(w)
			}
		}
	}
}

func (bfp *BreadthFirstPaths) HasPathTo(v int) bool {
	bfp.validateVertex(v)
	return bfp.marked[v]
}

func (bfp *BreadthFirstPaths) DistTo(v int) (int, error) {
	bfp.validateVertex(v)
	if !bfp.HasPathTo(v) {
		return -1, errors.New(fmt.Sprintf("No path between %v and %v", bfp.source, v))
	}
	return bfp.distTo[v], nil
}

func (bfp *BreadthFirstPaths) PathTo(v int) (*stack.Stack, error) {
	bfp.validateVertex(v)
	s := stack.NewEmptyStack()
	if !bfp.HasPathTo(v) {
		return &s, errors.New(fmt.Sprintf("No path between %v and %v", bfp.source, v))
	}

	for i := v; bfp.distTo[i] != 0; i = bfp.edgeTo[i] {
		s.Push(i)
	}
	s.Push(bfp.source)
	return &s, nil
}

func (bfp *BreadthFirstPaths) validateVertex(v int) {
	if v < 0 || v >= len(bfp.marked) {
		panic("Vertex out of bounds")
	}
}
