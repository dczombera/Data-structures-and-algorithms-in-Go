package acyclic_shortest_paths

import (
	"fmt"

	graph "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_directed_graph"
	q "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/queue"
	s "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/stack"
)

type DepthFirstOrder struct {
	marked           []bool
	preOrder         q.Queue
	postOrder        q.Queue
	reversePostOrder s.Stack
	prePos           []int
	postPos          []int
	preCounter       int
	postCounter      int
}

func NewDepthFirstOrder(g *graph.EdgeWeightedDigraph) *DepthFirstOrder {
	size := g.VerticesCount()
	dfo := &DepthFirstOrder{make([]bool, size), q.Queue{}, q.Queue{}, s.Stack{}, make([]int, size), make([]int, size), 0, 0}

	for v := 0; v < size; v++ {
		if !dfo.marked[v] {
			dfo.dfs(g, v)
		}
	}

	return dfo
}

func (dfo *DepthFirstOrder) dfs(g *graph.EdgeWeightedDigraph, v int) {
	dfo.marked[v] = true
	dfo.preOrder.Push(v)
	dfo.prePos[v] = dfo.preCounter
	dfo.preCounter++
	for _, e := range g.AdjacencyList(v) {
		w := e.To
		if !dfo.marked[w] {
			dfo.dfs(g, w)
		}
	}
	dfo.postOrder.Push(v)
	dfo.postPos[v] = dfo.postCounter
	dfo.postCounter++
}

func (dfo *DepthFirstOrder) PreOrder() q.Queue {
	return dfo.preOrder
}

func (dfo *DepthFirstOrder) PostOrder() q.Queue {
	return dfo.postOrder
}

func (dfo *DepthFirstOrder) ReversePostOrder() s.Stack {
	reverse := s.Stack{}
	for curr := dfo.postOrder.First; curr != nil; curr = curr.Next {
		reverse.Push(curr.Item)
	}
	return reverse
}

func (dfo *DepthFirstOrder) PrePos(v int) int {
	dfo.validateVertex(v)
	return dfo.prePos[v]
}

func (dfo *DepthFirstOrder) PostPos(v int) int {
	dfo.validateVertex(v)
	return dfo.postPos[v]
}

func (dfo *DepthFirstOrder) ReversePostPos(v int) int {
	dfo.validateVertex(v)
	return dfo.postCounter - dfo.postPos[v] - 1
}

func (dfo *DepthFirstOrder) validateVertex(v int) {
	if v < 0 || v >= len(dfo.marked) {
		panic(fmt.Sprintf("Vertex %v out of bounds", v))
	}
}
