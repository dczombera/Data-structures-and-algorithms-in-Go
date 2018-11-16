package prime_mst

import "testing"

type testCaseQueue struct {
	data            []IndexWeight
	decreaseWeights []IndexWeight
	increaseWeights []IndexWeight
	wantWeight      []IndexWeight
	nextMin         []int
}

type IndexWeight struct {
	index  int
	weight Weight
}

func TestMinPriorityQueue(t *testing.T) {
	testCases := []testCaseQueue{
		{
			data:            []IndexWeight{{0, 2.3}, {1, 0.42}, {2, 13.37}, {3, 0.42}, {4, -0.42}, {5, 0.0}, {6, -0.41}},
			decreaseWeights: []IndexWeight{{1, -0.42}, {2, 3.1}},
			increaseWeights: []IndexWeight{{4, -0.41}, {2, 42.42}, {3, 41.14}},
			wantWeight:      []IndexWeight{{0, 2.3}, {1, -0.42}, {2, 42.42}, {3, 41.14}, {4, -0.41}, {5, 0.0}, {6, -0.41}},
			nextMin:         []int{1, 6, 4, 5, 0, 3, 2},
		},
	}

	for _, tc := range testCases {
		pq := NewIndexMinPriorityQueue(len(tc.data))
		for _, d := range tc.data {
			pq.Insert(d.index, d.weight)
		}

		for _, w := range tc.decreaseWeights {
			pq.DecreaseWeight(w.index, w.weight)
		}

		for _, w := range tc.increaseWeights {
			pq.IncreaseWeight(w.index, w.weight)
		}

		for _, w := range tc.wantWeight {
			if pq.WeightOf(w.index) != w.weight {
				t.Errorf("Got weight of %v for index %v, want %v", pq.WeightOf(w.index), w.index, w.weight)
			}
		}

		for _, i := range tc.nextMin {
			min, err := pq.DelMin()
			if err != nil {
				t.Errorf("Got error %v, want next minimum index in priority queue", err)
			}

			if min != i {
				t.Errorf("Got index %v as next minimum, want %v", min, i)
			}
		}
	}

}

func TestMinPriorityQueueError(t *testing.T) {
	pq := NewIndexMinPriorityQueue(0)
	_, e := pq.DelMin()
	if e == nil {
		t.Errorf("Got %v, want error message", e)
	}
}
