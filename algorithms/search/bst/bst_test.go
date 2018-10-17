package bst

import "testing"

type testCase struct {
	data []KV
	want []KV
}

type KV struct {
	key Key
	val Value
}

func containsKey(n *Node, k Key) bool {
	q := []*Node{n}
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

}
