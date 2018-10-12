package sequential_search

import "testing"

type testCase struct {
	data []Node
	want []Node
}

var n = Node{42, "42", nil}
var testCases = []testCase{
	{[]Node{{1, "hola", nil}, {2, "hello", nil}}, []Node{{1, "hola", nil}, {2, "hello", nil}}},
	{[]Node{{1, "hola", nil}, {2, "hello", nil}, {1, "bye", nil}, {42, "hola", nil}}, []Node{{42, "hola", nil}, {1, "bye", nil}, {2, "hello", nil}}},
	{[]Node{}, []Node{}},
}
var errorTestCases = []testCase{
	{[]Node{{1, "hola", nil}, {2, "hello", nil}}, []Node{{5, "", nil}, {42, "", nil}}},
}

func TestSequentialSearch(t *testing.T) {
	for _, tc := range testCases {
		sq := NewEmptySequentialSearchST()
		for _, n := range tc.data {
			sq.Put(n.Key, n.Val)
		}
		for _, w := range tc.want {
			v, err := sq.Get(w.Key)
			if err != nil {
				t.Errorf("Got %v, want %v", err, w.Val)
			}
			if v != w.Val {
				t.Errorf("Got %v, want %v", v, w.Val)
			}
		}
	}
}

func TestErrorSequentialSearch(t *testing.T) {
	for _, tc := range errorTestCases {
		sq := NewEmptySequentialSearchST()
		for _, n := range tc.data {
			sq.Put(n.Key, n.Val)
		}
		for _, w := range tc.want {
			_, err := sq.Get(w.Key)
			if err == nil {
				t.Errorf("Got %v, want error message", err)
			}
		}
	}
}
