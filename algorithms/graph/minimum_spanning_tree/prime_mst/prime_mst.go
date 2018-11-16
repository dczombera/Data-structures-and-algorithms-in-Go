package prime_mst

import (
	"log"
	"math"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_graph"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_graph/edge"
)

// PrimeMST is a data type for computing the minimum spanning tree/forest in an edge weighted undirected graph using Prime’s algorithm with an indexed binary heap
type PrimeMST struct {
	edgeTo []*edge.Edge
	mst    []edge.Edge
	distTo []Weight
	weight float64
	marked []bool
	pq     *IndexMinPriorityQueue
}

func NewPrimeMST(g *edge_weighted_graph.EdgeWeightedGraph) PrimeMST {
	maxSize := g.VerticesCount()
	distTo := make([]Weight, maxSize)
	for i := 0; i < len(distTo); i++ {
		distTo[i] = Weight(math.Inf(1))
	}

	mst := PrimeMST{make([]*edge.Edge, maxSize), make([]edge.Edge, 0), distTo, 0.0, make([]bool, maxSize), NewIndexMinPriorityQueue(maxSize)}
	for v := 0; v < maxSize; v++ {
		// Run from each vertex to find minimum spanning forest
		if !mst.marked[v] {
			mst.prime(g, v)
		}
	}

	mst.initMST()
	return mst
}

func (mst *PrimeMST) prime(g *edge_weighted_graph.EdgeWeightedGraph, s int) {
	mst.distTo[s] = 0.0
	mst.pq.Insert(s, mst.distTo[s])
	for !mst.pq.Empty() {
		v, err := mst.pq.DelMin()
		if err != nil {
			log.Fatalln(err)
		}
		mst.scan(g, v)
	}
}

func (mst *PrimeMST) scan(g *edge_weighted_graph.EdgeWeightedGraph, v int) {
	mst.marked[v] = true
	for _, e := range g.AdjacencyList(v) {
		w := e.Other(v)
		if mst.marked[w] {
			// v-w is obsolete
			continue
		}

		weight := Weight(e.Weight())
		if weight < mst.distTo[w] {
			mst.edgeTo[w] = &e
			mst.distTo[w] = weight
			if mst.pq.Contains(w) {
				mst.pq.DecreaseWeight(w, weight)
			} else {
				mst.pq.Insert(w, weight)
			}
		}
	}
}

func (mst *PrimeMST) initMST() {
	mst.normalizeMST()
	mst.calculateWeight()
}

func (mst *PrimeMST) normalizeMST() {
	for _, e := range mst.edgeTo {
		if e != nil {
			mst.mst = append(mst.mst, *e)
		}
	}
}

func (mst *PrimeMST) calculateWeight() {
	for _, e := range mst.mst {
		mst.weight += e.Weight()
	}
}

func (mst *PrimeMST) Weight() float64 {
	return mst.weight
}

func (mst *PrimeMST) Edges() []edge.Edge {
	return mst.mst
}
