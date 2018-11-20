package acyclic_shortest_paths

import (
	"errors"
	"fmt"

	graph "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_directed_graph"
	s "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/stack"
)

type TopologicalOrder struct {
	order s.Stack
	rank  []int
}

func NewTopologicalOrder(g *graph.EdgeWeightedDigraph) (TopologicalOrder, error) {
	dc := NewDirectedEdgeWeightedCycle(g)
	if dc.HasCycle() {
		return TopologicalOrder{}, errors.New("Edge weighted Digraph has a cycle and therefore is not an directed acyclic graph (DAG). However, a DAG is prerequisite for calculating the topological order")
	}
	reverse := NewDepthFirstOrder(g).ReversePostOrder()
	rank := make([]int, reverse.Size)
	t := TopologicalOrder{reverse, rank}
	counter := 0
	for curr := reverse.First; curr != nil; curr = curr.Next {
		t.rank[curr.Item] = counter
		counter++
	}
	return t, nil
}

func (t TopologicalOrder) HasOrder() bool {
	return t.order.Size > 0
}

func (t TopologicalOrder) Order() s.Stack {
	return t.order
}

func (t TopologicalOrder) Rank(v int) int {
	t.validateVertex(v)
	return t.rank[v]
}

func (t TopologicalOrder) validateVertex(v int) {
	if v < 0 || v >= t.order.Size {
		panic(fmt.Sprintf("Vertex %v out of bounds", v))
	}
}
