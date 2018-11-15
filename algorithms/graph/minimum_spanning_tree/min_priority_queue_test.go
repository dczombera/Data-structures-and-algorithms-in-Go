package minimum_spanning_tree

import (
	"testing"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_graph/edge"
)

type testCaseQueue struct {
	data []edge.Edge
	want []edge.Edge
}

func TestMinPriorityQueue(t *testing.T) {
	testCases := []testCaseQueue{
		{[]edge.Edge{edge.NewEdge(0, 1, 2.3), edge.NewEdge(2, 3, 0.42), edge.NewEdge(1, 2, 13.37)}, []edge.Edge{edge.NewEdge(2, 3, 0.42), edge.NewEdge(0, 1, 2.3), edge.NewEdge(1, 2, 13.37)}},
		{[]edge.Edge{edge.NewEdge(0, 1, 0.0), edge.NewEdge(3, 8, -2.3), edge.NewEdge(4, 5, 2.3), edge.NewEdge(5, 6, -2.3)}, []edge.Edge{edge.NewEdge(3, 8, -2.3), edge.NewEdge(5, 6, -2.3), edge.NewEdge(0, 1, 0.0), edge.NewEdge(4, 5, 2.3)}},
	}

	pq := NewMinPriorityQueue()
	for _, tc := range testCases {
		for _, e := range tc.data {
			pq.Insert(e)
		}

		if pq.Size() != len(tc.data) {
			t.Errorf("Got size %d, want %d", len(tc.data), pq.Size())
		}

		for _, e := range tc.want {
			min, _ := pq.DelMin()
			if min != e {
				t.Errorf("Got %v, want %v", min, e)
			}
		}
	}
}

func TestMinPriorityQueueError(t *testing.T) {
	pq := NewMinPriorityQueue()
	_, e := pq.DelMin()
	if e == nil {
		t.Errorf("Got %v, want error message", e)
	}
}
