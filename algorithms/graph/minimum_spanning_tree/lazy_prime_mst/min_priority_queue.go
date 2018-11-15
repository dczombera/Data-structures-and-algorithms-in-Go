package lazy_prime_mst

import (
	"errors"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_graph/edge"
)

// MinPriorityQueue is a minimum priority queue using a binary heap as underlying data structure
type MinPriorityQueue struct {
	pq []edge.Edge
	n  int
}

func NewMinPriorityQueue() MinPriorityQueue {
	return MinPriorityQueue{[]edge.Edge{edge.NewEmptyEdge()}, 0}
}

func (pq *MinPriorityQueue) Size() int {
	return pq.n
}

func (pq *MinPriorityQueue) IsEmpty() bool {
	return pq.n == 0
}

func (pq *MinPriorityQueue) swim(k int) {
	for k > 1 && pq.greater(k/2, k) {
		pq.exch(k, k/2)
		k = k / 2
	}
}

func (pq *MinPriorityQueue) sink(k int) {
	for k*2 <= pq.n {
		i := k * 2
		if i < pq.n && pq.greater(i, i+1) {
			i++
		}
		if !pq.greater(k, i) {
			break
		}
		pq.exch(k, i)
		k = i
	}
}

func (pq *MinPriorityQueue) greater(i, j int) bool {
	return pq.pq[i].Compare(pq.pq[j]) > 0
}

func (pq *MinPriorityQueue) exch(i, j int) {
	pq.pq[i], pq.pq[j] = pq.pq[j], pq.pq[i]
}

func (pq *MinPriorityQueue) Insert(i edge.Edge) {
	pq.n++
	pq.pq = append(pq.pq, i)
	pq.swim(pq.n)
}

func (pq *MinPriorityQueue) Min() edge.Edge {
	return pq.pq[1]
}

func (pq *MinPriorityQueue) DelMin() (edge.Edge, error) {
	if pq.IsEmpty() {
		return edge.NewEmptyEdge(), errors.New("Priority queue is empty!")
	}
	k := pq.pq[1]
	pq.exch(1, pq.n)
	pq.pq = append(pq.pq[:pq.n])
	pq.n--
	pq.sink(1)
	return k, nil
}
