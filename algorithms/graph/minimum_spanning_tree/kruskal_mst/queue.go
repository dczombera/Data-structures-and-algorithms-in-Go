package kruskal_mst

import (
	"errors"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_graph/edge"
)

type Queue struct {
	First *Node
	Last  *Node
	Size  int
}

type Node struct {
	Item edge.Edge
	Next *Node
}

func NewNode(e edge.Edge) *Node {
	return &Node{e, nil}
}

func NewEmptyQueue() Queue {
	return Queue{nil, nil, 0}
}

func (q *Queue) Push(e edge.Edge) {
	oldLast := q.Last
	q.Last = NewNode(e)
	if q.IsEmpty() {
		q.First = q.Last
	} else {
		oldLast.Next = q.Last
	}
	q.Size++
}

func (q *Queue) Pop() (edge.Edge, error) {
	if q.IsEmpty() {
		return edge.NewEmptyEdge(), errors.New("Queue is empty!")
	}

	n := q.First
	q.First = n.Next
	q.Size--

	if q.IsEmpty() {
		q.Last = nil
	}

	return n.Item, nil
}

func (q *Queue) Peek() (edge.Edge, error) {
	if q.IsEmpty() {
		return edge.NewEmptyEdge(), errors.New("There is nothing to peek, queue is empty!")
	}

	return q.First.Item, nil
}

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}
