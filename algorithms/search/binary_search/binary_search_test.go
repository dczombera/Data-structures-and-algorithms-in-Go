package binary_search

import (
	"log"
	"testing"
)

type Node struct {
	Key Key
	Val Value
}

type testCase struct {
	data []Node
	want []Node
}

var testCases = []testCase{
	{[]Node{{1, "hola"}, {2, "hello"}}, []Node{{1, "hola"}, {2, "hello"}}},
	{[]Node{{1, "hola"}, {2, "hello"}, {1, "bye"}, {42, "hola"}}, []Node{{42, "hola"}, {1, "bye"}, {2, "hello"}}},
	{[]Node{}, []Node{}},
}
var errorTestCases = []testCase{
	{[]Node{{1, "hola"}, {2, "hello"}}, []Node{{5, ""}, {42, ""}}},
}

func TestBinarySearch(t *testing.T) {
	for _, tc := range testCases {
		bs := NewEmptyBinarySearchST()
		for _, n := range tc.data {
			bs.Put(n.Key, n.Val)
		}
		log.Println(bs)
		for _, w := range tc.want {
			v, err := bs.Get(w.Key)
			if err != nil {
				t.Errorf("Got %v, want %v", err, w.Val)
			}
			if v != w.Val {
				t.Errorf("Got %v, want %v", v, w.Val)
			}
		}
	}
}
