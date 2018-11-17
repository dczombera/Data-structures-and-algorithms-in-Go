package kruskal_mst

import (
	"github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/graph/minimum_spanning_tree/priority_queue"
	"github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/search/union_find"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_graph"
)

// KruskalMST is a data type that computes the minimum spanning tree of an undirected edge weighted graph.
// The implementation uses Kruskal's algorithm, a minimum priority queue and the union find data type.
type KruskalMST struct {
	mst    Queue
	weight float64
	pq     priority_queue.MinPriorityQueue
	uf     *union_find.UF
}

func NewKruskalMST(g *edge_weighted_graph.EdgeWeightedGraph) *KruskalMST {
	uf := union_find.NewUnionFind(g.VerticesCount())
	pq := priority_queue.NewMinPriorityQueue()
	for _, e := range g.Edges() {
		pq.Insert(e)
	}

	mst := KruskalMST{NewEmptyQueue(), 0.0, pq, uf}
	for !pq.IsEmpty() && (mst.mst.Size < g.VerticesCount()-1) {
		edge, err := pq.DelMin()
		if err != nil {
			panic(err)
		}
		v := edge.Either()
		w := edge.Other(v)
		if !uf.Connected(v, w) {
			mst.mst.Push(edge)
			mst.weight += edge.Weight()
			uf.Union(v, w)
		}
	}

	return &mst
}

func (mst *KruskalMST) Weight() float64 {
	return mst.weight
}

func (mst *KruskalMST) Edges() Queue {
	return mst.mst
}
