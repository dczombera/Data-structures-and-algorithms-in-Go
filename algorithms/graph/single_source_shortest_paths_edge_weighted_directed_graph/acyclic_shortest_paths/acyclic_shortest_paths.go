package acyclic_shortest_paths

import (
	"errors"
	"fmt"
	"math"

	graph "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_directed_graph"
)

type AcyclicSP struct {
	edgeTo  []graph.DirectedEdge
	distoTo []float64
	source  int
}

var infinity = math.Inf(1)

func NewAcyclicSP(g *graph.EdgeWeightedDigraph, s int) (AcyclicSP, error) {
	t, err := NewTopologicalOrder(g)
	if err != nil {
		return AcyclicSP{}, err
	}

	distTo := make([]float64, g.VerticesCount())
	for i := range distTo {
		distTo[i] = infinity
	}
	distTo[s] = 0.0
	asp := AcyclicSP{make([]graph.DirectedEdge, g.VerticesCount()), distTo, s}
	for curr := t.Order().First; curr != nil; curr = curr.Next {
		for _, e := range g.AdjacencyList(curr.Item) {
			asp.relax(e)
		}
	}
	return asp, nil
}

func (asp *AcyclicSP) relax(e graph.DirectedEdge) {
	v := e.From
	w := e.To
	if asp.distoTo[w] > asp.distoTo[v]+e.Weight {
		asp.edgeTo[w] = e
		asp.distoTo[w] = asp.distoTo[v] + e.Weight
	}
}

func (asp AcyclicSP) HasPathTo(v int) bool {
	asp.validateVertex(v)
	return asp.distoTo[v] < infinity
}

func (asp AcyclicSP) PathTo(v int) (graph.Stack, error) {
	asp.validateVertex(v)
	if !asp.HasPathTo(v) {
		return graph.Stack{}, errors.New(fmt.Sprintf("No shortest path found from source %v to %v\n", asp.source, v))
	}

	stackPath := graph.Stack{}
	curr := asp.edgeTo[v]
	for ; curr.From != asp.source; curr = asp.edgeTo[curr.From] {
		stackPath.Push(curr)
	}
	stackPath.Push(curr)

	return stackPath, nil
}

func (asp *AcyclicSP) DistTo(v int) (float64, error) {
	asp.validateVertex(v)
	if !asp.HasPathTo(v) {
		return -1.0, errors.New(fmt.Sprintf("No shortest path found from source %v to %v\n", asp.source, v))
	}
	return asp.distoTo[v], nil
}

func (asp *AcyclicSP) validateVertex(v int) {
	if v < 0 || v >= len(asp.edgeTo) {
		panic(fmt.Sprintf("Vertex %v out of bounds", v))
	}
}
