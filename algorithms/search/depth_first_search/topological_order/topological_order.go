package topological_order

import (
	"errors"

	"github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/search/depth_first_search/depth_first_order"
	"github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/search/depth_first_search/directed_cycle"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/directed_graph"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/stack"
)

type Topological struct {
	order stack.Stack
	pos   []int
}

func NewTopologicalOrder(g *directed_graph.Digraph) (*Topological, error) {
	t := Topological{stack.NewEmptyStack(), make([]int, g.Vertices())}
	cycle := directed_cycle.NewDirectedCycle(g)
	if cycle.HasCycle() {
		return &t, errors.New("Graph has a cycle and therefore is not a directed acycle graph (DAG). However, a DAG is prerequisite for a topological order")
	}

	reverse := depth_first_order.NewDepthFirstOrder(g).ReversePostOrder()
	t.order = reverse
	pos := 0
	for curr := t.order.First; curr != nil; curr = curr.Next {
		t.pos[curr.Item] = pos
		pos++
	}
	return &t, nil
}

func (t *Topological) Order() stack.Stack {
	return t.order
}

func (t *Topological) HasOrder() bool {
	return t.order.Size > 0
}

func (t *Topological) IsDAG() bool {
	return t.order.Size > 0
}

func (t *Topological) Rank(v int) int {
	t.validateVertex(v)
	if t.HasOrder() {
		return t.pos[v]
	}
	return -1
}

func (t *Topological) validateVertex(v int) {
	if v < 0 || v >= len(t.pos) {
		panic("Vertex out of bounds")
	}
}
