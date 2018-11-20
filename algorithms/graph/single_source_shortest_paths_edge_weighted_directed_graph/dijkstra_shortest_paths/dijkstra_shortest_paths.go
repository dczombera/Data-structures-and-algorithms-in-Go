package dijkstra_shortest_paths

import (
	"errors"
	"fmt"
	"log"
	"math"

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

	return sp
}

func (sp *DijkstraSP) relax(e graph.DirectedEdge) {
	v := e.From
	w := e.To
	if sp.distTo[w] > (sp.distTo[v] + e.Weight) {
		sp.distTo[w] = sp.distTo[v] + e.Weight
		sp.edgeTo[w] = e
		if sp.pq.Contains(w) {
			sp.pq.DecreaseWeight(w, pq.Weight(sp.distTo[w]))
		} else {
			sp.pq.Insert(w, pq.Weight(sp.distTo[w]))
		}
	}
}

func (sp *DijkstraSP) HasPathTo(v int) bool {
	sp.validateVertex(v)
	return sp.distTo[v] < math.Inf(1)
}

func (sp *DijkstraSP) DistTo(v int) (float64, error) {
	sp.validateVertex(v)
	if !sp.HasPathTo(v) {
		return -1, errors.New(fmt.Sprintf("No shortest path found from source %v to %v\n", sp.source, v))
	}
	return sp.distTo[v], nil
}

func (sp *DijkstraSP) PathTo(v int) ([]graph.DirectedEdge, error) {
	sp.validateVertex(v)
	if !sp.HasPathTo(v) {
		return []graph.DirectedEdge{}, errors.New(fmt.Sprintf("There is no shortest path between source vertex %v and %v", sp.source, v))
	}

	pathStack := graph.Stack{}
	curr := graph.DirectedEdge{}
	for curr = sp.edgeTo[v]; curr.From != sp.source; curr = sp.edgeTo[curr.From] {
		pathStack.Push(curr)
	}
	pathStack.Push(curr)

	// We want the convenience of using range, therefore we return a slice
	path := make([]graph.DirectedEdge, 0, pathStack.Size)
	for curr := pathStack.First; curr != nil; curr = curr.Next {
		path = append(path, curr.Item)
	}
	return path, nil
}

func (sp *DijkstraSP) validateVertex(v int) {
	if v < 0 || v >= len(sp.edgeTo) {
		log.Fatalf("Vertext %v out of bounds", v)
	}
}
