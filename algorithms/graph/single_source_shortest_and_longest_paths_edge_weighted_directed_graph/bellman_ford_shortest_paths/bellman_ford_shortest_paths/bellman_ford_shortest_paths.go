package bellman_ford_shortest_paths

import (
	"errors"
	"fmt"
	"math"

	c "github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/graph/single_source_shortest_and_longest_paths_edge_weighted_directed_graph/acyclic_shortest_paths"

	graph "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_directed_graph"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/queue"
)

type BellmanFordSP struct {
	edgeTo  []graph.DirectedEdge
	distTo  []float64
	onQueue []bool
	queue   queue.Queue
	source  int
	cost    int
	cycle   graph.Stack
}

var infinity = math.Inf(1)

func NewBellmanFordSp(g *graph.EdgeWeightedDigraph, s int) *BellmanFordSP {
	size := g.VerticesCount()
	distTo := make([]float64, size)
	for i := 0; i < size; i++ {
		distTo[i] = infinity
	}
	sp := &BellmanFordSP{make([]graph.DirectedEdge, size), distTo, make([]bool, size), queue.NewEmptyQueue(), s, 0, graph.Stack{}}
	sp.distTo[s] = 0.0
	sp.queue.Push(s)
	sp.onQueue[s] = true
	for !sp.queue.IsEmpty() && !sp.hasNegativeCycle() {
		v, err := sp.queue.Pop()
		checkError(err)
		sp.onQueue[v] = false
		sp.relax(g, v)
	}
	return sp
}

func (sp *BellmanFordSP) relax(g *graph.EdgeWeightedDigraph, v int) {
	for _, e := range g.AdjacencyList(v) {
		w := e.To
		if sp.distTo[w] > sp.distTo[v]+e.Weight {
			sp.edgeTo[w] = e
			sp.distTo[w] = sp.distTo[v] + e.Weight
			sp.queue.Push(w)
			if !sp.onQueue[w] {
				sp.onQueue[w] = true
			}
		}
		sp.cost++
		if sp.cost%g.VerticesCount() == 0 {
			if sp.findNegativeCycle() {
				return
			}
		}
	}
}

func (sp *BellmanFordSP) findNegativeCycle() bool {
	size := len(sp.edgeTo)
	g := graph.NewEdgeWeightedDigraph(size)
	for i := 0; i < size; i++ {
		g.AddEdge(sp.edgeTo[i])
	}
	cycle := c.NewDirectedEdgeWeightedCycle(g)
	sp.cycle = cycle.Cycle()
	return sp.cycle.Size > 0
}

func (sp BellmanFordSP) hasNegativeCycle() bool {
	return sp.cycle.Size > 0
}

func (sp BellmanFordSP) HasPathTo(v int) bool {
	sp.validateVertex(v)
	return sp.distTo[v] < infinity
}

func (sp BellmanFordSP) PathTo(v int) (graph.Stack, error) {
	sp.validateVertex(v)
	if sp.hasNegativeCycle() {
		return graph.Stack{}, errors.New("No shortest paths found because graph has negative cycle")
	}
	pathStack := graph.Stack{}
	curr := sp.edgeTo[v]
	for ; curr.From != sp.source; curr = sp.edgeTo[curr.From] {
		pathStack.Push(curr)
	}
	pathStack.Push(curr)
	return pathStack, nil
}

func (sp BellmanFordSP) DistTo(v int) (float64, error) {
	sp.validateVertex(v)
	if sp.hasNegativeCycle() {
		return infinity, errors.New("No shortest paths found because graph has negative cycle")
	}
	return sp.distTo[v], nil
}

func (sp BellmanFordSP) validateVertex(v int) {
	if v < 0 || v >= len(sp.edgeTo) {
		panic(fmt.Sprintf("Vertex out of bounds"))
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
