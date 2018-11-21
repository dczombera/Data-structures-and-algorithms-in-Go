package acyclic_longest_paths

import (
	"errors"
	"fmt"
	"math"

	as "github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/graph/single_source_shortest_and_longest_paths_edge_weighted_directed_graph/acyclic_shortest_paths"
	graph "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_directed_graph"
)

type AcyclicLP struct {
	edgeTo []graph.DirectedEdge
	distTo []float64
	source int
}

var minInfinity = math.Inf(-1)

func NewAcyclicLP(g *graph.EdgeWeightedDigraph, s int) (AcyclicLP, error) {
	t, err := as.NewTopologicalOrder(g)
	if err != nil {
		return AcyclicLP{}, err
	}

	distTo := make([]float64, g.VerticesCount())
	for i := range distTo {
		distTo[i] = minInfinity
	}
	distTo[s] = 0.0
	alp := AcyclicLP{make([]graph.DirectedEdge, g.VerticesCount()), distTo, s}
	for curr := t.Order().First; curr != nil; curr = curr.Next {
		for _, e := range g.AdjacencyList(curr.Item) {
			alp.relax(e)
		}
	}

	return alp, nil
}

func (alp *AcyclicLP) relax(e graph.DirectedEdge) {
	v := e.From
	w := e.To
	if alp.distTo[w] < alp.distTo[v]+e.Weight {
		alp.edgeTo[w] = e
		alp.distTo[w] = alp.distTo[v] + e.Weight
	}
}

func (alp AcyclicLP) HasPathTo(v int) bool {
	alp.validateVertex(v)
	return alp.distTo[v] > minInfinity
}

func (alp AcyclicLP) PathTo(v int) (graph.Stack, error) {
	alp.validateVertex(v)
	if !alp.HasPathTo(v) {
		return graph.Stack{}, errors.New(fmt.Sprintf("No longest path found from source %v to %v\n", alp.source, v))
	}

	pathStack := graph.Stack{}
	curr := alp.edgeTo[v]
	for ; curr.From != alp.source; curr = alp.edgeTo[curr.From] {
		pathStack.Push(curr)
	}
	pathStack.Push(curr)
	return pathStack, nil
}

func (alp AcyclicLP) DistTo(v int) (float64, error) {
	alp.validateVertex(v)
	if !alp.HasPathTo(v) {
		return minInfinity, errors.New(fmt.Sprintf("No longest path found from source %v to %v\n", alp.source, v))
	}
	return alp.distTo[v], nil
}

func (alp AcyclicLP) validateVertex(v int) {
	if v < 0 || v >= len(alp.edgeTo) {
		panic(fmt.Sprintf("Vertex %v out of bounds", v))
	}
}
