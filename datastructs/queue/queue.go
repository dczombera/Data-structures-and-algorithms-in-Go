package queue

import (
	"errors"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/node"
)

type Queue struct {
	First *node.Node
	Last  *node.Node
	Size  int
}

func NewQueue(n *node.Node) Queue {
	return Queue{n, n, 1}
}

func NewEmptyQueue() Queue {
	return Queue{nil, nil, 0}
}

func (q *Queue) Push(i int) {
	oldLast := q.Last
	q.Last = &node.Node{i, nil}
	if q.IsEmpty() {
		q.First = q.Last
	} else {
		oldLast.Next = q.Last
	}
	q.Size++
}

func (q *Queue) Pop() (int, error) {
	var n *node.Node
	if q.IsEmpty() {
		return n.Item, errors.New("Queue is empty!")
	}

	n = q.First
	q.First = n.Next
	q.Size--

	if q.IsEmpty() {
		q.Last = nil
	}

	return n.Item, nil
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("There is nothing to peek, queue is empty!")
	}

	return q.First.Item, nil
}

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}
