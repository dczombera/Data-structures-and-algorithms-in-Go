package depth_first_order

import (
	"log"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/directed_graph"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/queue"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/stack"
)

type DephtFirstOrder struct {
	marked           []bool
	preOrder         queue.Queue
	postOrder        queue.Queue
	reversePostOrder stack.Stack
	prePos           []int
	postPos          []int
	preCounter       int
	postCounter      int
}

func NewDepthFirstOrder(g *directed_graph.Digraph) *DephtFirstOrder {
	length := g.Vertices()
	dfo := DephtFirstOrder{make([]bool, length), queue.NewEmptyQueue(), queue.NewEmptyQueue(), stack.NewEmptyStack(), make([]int, length), make([]int, length), 0, 0}
	for v := 0; v < length; v++ {
		if !dfo.marked[v] {
			dfo.dfs(g, v)
		}
	}

	return &dfo
}

func (dfo *DephtFirstOrder) dfs(g *directed_graph.Digraph, v int) {
	dfo.marked[v] = true
	dfo.preOrder.Push(v)
	dfo.prePos[v] = dfo.preCounter
	dfo.preCounter++
	for _, w := range g.AdjacencyList(v) {
		if !dfo.marked[w] {
			log.Println(w)
			dfo.dfs(g, w)
		}
	}
	dfo.postOrder.Push(v)
	dfo.postPos[v] = dfo.postCounter
	dfo.postCounter++
}

func (dfo *DephtFirstOrder) PreOrder() queue.Queue {
	return dfo.preOrder
}

func (dfo *DephtFirstOrder) PostOrder() queue.Queue {
	return dfo.postOrder
}

func (dfo *DephtFirstOrder) ReversePostOrder() stack.Stack {
	if dfo.reversePostOrder.Size == 0 {
		for curr := dfo.postOrder.First; curr != nil; curr = curr.Next {
			dfo.reversePostOrder.Push(curr.Item)
		}
	}
	return dfo.reversePostOrder
}

func (dfo *DephtFirstOrder) PrePos(v int) int {
	dfo.validateVertex(v)
	return dfo.prePos[v]
}

func (dfo *DephtFirstOrder) PostPos(v int) int {
	dfo.validateVertex(v)
	return dfo.postPos[v]
}

func (dfo *DephtFirstOrder) ReversePostPos(v int) int {
	dfo.validateVertex(v)
	return dfo.postCounter - dfo.postPos[v]
}

func (dfo *DephtFirstOrder) validateVertex(v int) {
	if v < 0 || v >= len(dfo.marked) {
		panic("Vertex out of bounds")
	}
}
