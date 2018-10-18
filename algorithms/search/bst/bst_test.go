package bst

import (
	"testing"
)

type testCase struct {
	data []KV
	want []KV
}

type KV struct {
	key Key
	val Value
}

// Traverse tree without using get method of BST
func containsKey(bst *BST, k Key) bool {
	q := []*Node{bst.Root}
	var n *Node
	for len(q) > 0 {
		n, q = q[0], q[1:]
		if n == nil {
			continue
		}

		if n.Key == k {
			return true
		}
		q = append(q, n.Left, n.Right)
	}
	return false
}

func TestBSTPut(t *testing.T) {
	testCases := []testCase{
		{[]KV{{1, "hola"}, {2, "hello"}, {3, "hallo"}}, []KV{{1, "hola"}, {2, "hello"}, {3, "hallo"}}},
		{[]KV{{42, "foo"}, {0, "bar"}, {42, "bar"}, {0, "foo"}}, []KV{{42, "bar"}, {0, "foo"}}},
	}
	for _, tc := range testCases {
		bst := NewEmptyBST()
		for _, d := range tc.data {
			bst.Put(d.key, d.val)
		}
		for _, w := range tc.want {
			if !containsKey(&bst, w.key) {
				t.Errorf("BST does not contain key %v", w.key)
			}
		}
	}
}

func TestBSTGet(t *testing.T) {
	testCases := []testCase{
		{[]KV{{1, "hola"}, {2, "hello"}, {3, "hallo"}}, []KV{{1, "hola"}, {2, "hello"}, {3, "hallo"}}},
		{[]KV{{42, "foo"}, {0, "bar"}, {42, "bar"}, {0, "foo"}}, []KV{{42, "bar"}, {0, "foo"}}},
		{[]KV{{-42, "han"}, {42, "shot"}, {0, "first"}, {0, "second"}, {42, "jumped"}}, []KV{{0, "second"}, {-42, "han"}, {42, "jumped"}}},
	}

	for _, tc := range testCases {
		bst := NewEmptyBST()
		for _, d := range tc.data {
			bst.Put(d.key, d.val)
		}
		for _, w := range tc.want {
			v, err := bst.Get(w.key)
			if err != nil {
				t.Errorf("Got error '%v' while looking for value %v of key %v", err, w.val, w.key)
			} else if v != w.val {
				t.Errorf("Got %v, want %v", v, w.val)
			}
		}
	}
}

func TestBSTDelete(t *testing.T) {
	testCases := []testCase{
		{[]KV{{1, "hola"}, {2, "hello"}, {2, ""}}, []KV{{1, "hola"}, {2, ""}}},
		{[]KV{{42, "foo"}, {0, "bar"}, {42, "bar"}, {0, "foo"}, {42, ""}, {0, ""}}, []KV{{42, ""}, {0, ""}}},
		{[]KV{{-42, "han"}, {42, "shot"}, {0, "first"}, {0, ""}}, []KV{{-42, "han"}, {0, ""}, {42, "shot"}}},
	}

	for _, tc := range testCases {
		bst := NewEmptyBST()
		for _, d := range tc.data {
			bst.Put(d.key, d.val)
		}

		for _, w := range tc.want {
			v, err := bst.Get(w.key)
			if err != nil {
				if w.val != "" {
					t.Errorf("Got error '%v' while looking for value %v of key %v", err, w.val, w.key)
				}
			} else if v != w.val {
				t.Errorf("Got %v, want %v", v, w.val)
			}
		}
	}
}
