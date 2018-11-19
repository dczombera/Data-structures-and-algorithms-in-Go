package dijkstra_shortest_paths

import (
	"errors"
	"log"
	"math"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/stack"

	pq "github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/graph/minimum_spanning_tree/priority_queue"
	graph "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_directed_graph"
)

type DijkstraSP struct {
	edgeTo []graph.DirectedEdge
	distTo []float64
	pq     *pq.IndexMinPriorityQueue
	source int
}

func NewDijkstraSP(g *graph.EdgeWeightedDigraph, s int) *DijkstraSP {
	size := g.VerticesCount()
	distTo := make([]float64, size)
	for i := range distTo {
		distTo[i] = math.Inf(1)
	}
	sp := &DijkstraSP{make([]graph.DirectedEdge, size), distTo, pq.NewIndexMinPriorityQueue(size), s}

	sp.distTo[s] = 0.0
	sp.pq.Insert(s, 0.0)
	for !sp.pq.Empty() {
		v, err := sp.pq.DelMin()
		if err != nil {
			panic(err)
		}
		for _, e := range g.AdjacencyList(v) {
			sp.relax(e)
		}
	}
}

func (sp *DijkstraSP) relax(e graph.DirectedEdge) {
	v := e.From()
	w := e.To()
	if sp.distTo[w] > sp.distTo[v]+e.Weight() {
		sp.distTo[w] = sp.distTo[v] + e.Weight()
		sp.edgeTo[w] = e
		if sp.pq.Contains(w) {
			sp.pq.DecreaseWeight(w, pq.Weight(e.Weight()))
		} else {
			sp.pq.Insert(w, pq.Weight(e.Weight()))
		}
	}
}

func (sp *DijkstraSP) HasPathTo(v int) bool {
	sp.validateVertex(v)
	return sp.distTo[v] < math.Inf(1)
}

func (sp *DijkstraSP) PathTo(v int) ([]graph.DirectedEdge, error) {
	sp.validateVertex(v)
	if !sp.HasPathTo(v) {
		return graph.DirectedEdge(), errors.New("There is no shortest path between source vertex %v and %v", sp.source, v)
	}

	// TODO Implement stack for directed edge
	pathStack := stack.NewEmptyStack()
	curr := graph.DirectedEdge{}
	for curr = sp.edgeTo[v]; curr.From() != sp.source; curr = sp.edgeTo[v] {
		pathStack.Push(curr)
	}
	pathStack.Push(curr)
	// We want the convenience of using range, therefore we return a slice
	path := make([]graph.DirectedEdge, 0, pathStack.Size())
	for curr = pathStack.First; curr != nil; curr = curr.Next {
		path = append(path, curr)
	}

	return path, nil
}

func (sp *DijkstraSP) validateVertex(v int) {
	if v < 0 || v >= len(sp.edgeTo) {
		log.Fatalf("Vertext %v out of bounds", v)
	}
}
