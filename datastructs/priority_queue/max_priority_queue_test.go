package priority_queue

import (
	"testing"
)

type testCase struct {
	data []Key
	want []Key
}

var testCases = []testCase{
	{[]Key{2, 4, 3, 1, 5, 0}, []Key{5, 4, 3, 2, 1, 0}},
	{[]Key{-3, 0, -42, 99, 0, 42}, []Key{99, 42, 0, 0, -3, -42}},
	{[]Key{-3, -9, -1, -12, -2, -8}, []Key{-1, -2, -3, -8, -9, -12}},
	{[]Key{0, 0, 0}, []Key{0, 0, 0}},
	{[]Key{-1, 0, 0, 0}, []Key{0, 0, 0, -1}},
}

func TestMaxPriorityQueue(t *testing.T) {
	pq := NewEmptyMaxPriorityQueue()
	for _, tc := range testCases {
		for _, e := range tc.data {
			pq.Insert(e)
		}

		if pq.Size() != len(tc.data) {
			t.Errorf("Got size %d, want %d", len(tc.data), pq.Size())
		}

		for _, e := range tc.want {
			max, _ := pq.DelMax()
			if max != e {
				t.Errorf("Got %d, want %d", max, e)
			}
		}
	}
}

func TestMaxPriorityQueueError(t *testing.T) {
	pq := NewEmptyMaxPriorityQueue()
	_, e := pq.DelMax()
	if e == nil {
		t.Errorf("Got %v, want error message", e)
	}
}
