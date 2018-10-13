package binary_search

import (
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
}
var errorTestCases = []testCase{
	{[]Node{{1, "hola"}, {2, "hello"}}, []Node{{5, ""}, {42, ""}}},
}

var emptyTestCases = []testCase{
	{[]Node{}, []Node{}},
}

func TestBinarySearchPut(t *testing.T) {
	for _, tc := range testCases {
		bs := NewEmptyBinarySearchST()
		for _, n := range tc.data {
			bs.Put(n.Key, n.Val)
		}
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

func TestBinarySearchEmptyPut(t *testing.T) {
	for _, tc := range emptyTestCases {
		bs := NewEmptyBinarySearchST()
		for _, n := range tc.data {
			bs.Put(n.Key, n.Val)
		}
		if bs.Size() != 0 {
			t.Errorf("Got %d, want %d", bs.Size(), 0)
		}
	}
}

func TestErrorBinarySearch(t *testing.T) {
	for _, tc := range errorTestCases {
		bs := NewEmptyBinarySearchST()
		for _, n := range tc.data {
			bs.Put(n.Key, n.Val)
		}
		for _, w := range tc.want {
			_, err := bs.Get(w.Key)
			if err == nil {
				t.Errorf("Got %v, want error message", err)
			}
		}
	}
}

func TestBinarySearchDelete(t *testing.T) {
	for _, tc := range testCases {
		bs := NewEmptyBinarySearchST()
		for _, n := range tc.data {
			bs.Put(n.Key, n.Val)
		}
		bs.Delete(Key(tc.data[0].Key))
		bs.Delete(tc.data[len(tc.data)-1].Key)
		want := len(tc.want) - 2
		if bs.Size() != want {
			t.Errorf("Got %d, want %d", bs.Size(), want)
		}
	}
}
