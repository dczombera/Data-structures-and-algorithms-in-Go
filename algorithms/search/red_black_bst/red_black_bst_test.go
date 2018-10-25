package red_black_bst

import "testing"

type testCase struct {
	data []KV
	want []KV
}

type KV struct {
	key Key
	val Value
}

// Traverse tree without using get method of BST
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

func countTreeLevels(n *Node) int {
	if n == nil {
		return 0
	}

	q := []*Node{n}
	lvls := 0
	for len(q) > 0 {
		lvls++
		tmp := []*Node{}
		for len(q) > 0 {
			n, q = q[0], q[1:]
			if n.Left != nil {
				tmp = append(tmp, n.Left)
			}
			if n.Right != nil {
				tmp = append(tmp, n.Right)
			}
		}
		q = tmp
	}

	return lvls
}

func TestRedBlackBSTPut(t *testing.T) {
	testCases := []testCase{
		{[]KV{{1, "hola"}, {2, "hello"}, {3, "hallo"}}, []KV{{1, "hola"}, {2, "hello"}, {3, "hallo"}}},
		{[]KV{{42, "foo"}, {0, "bar"}, {42, "bar"}, {0, "foo"}}, []KV{{42, "bar"}, {0, "foo"}}},
	}

	for _, tc := range testCases {
		bst := NewEmptyRedBlackBST()
		for _, d := range tc.data {
			bst.Put(d.key, d.val)
		}

		for _, w := range tc.want {
			if !containsKey(bst.Root, w.key) {
				t.Errorf("Red Black BST does not contain key %v with value %v", w.key, w.val)
			}
		}
	}
}

func TestRedBlackBSTGet(t *testing.T) {
	testCases := []testCase{
		{[]KV{{1, "hola"}, {2, "hello"}, {3, "hallo"}}, []KV{{1, "hola"}, {2, "hello"}, {3, "hallo"}}},
		{[]KV{{42, "foo"}, {0, "bar"}, {42, "bar"}, {0, "foo"}}, []KV{{42, "bar"}, {0, "foo"}}},
		{[]KV{{-42, "han"}, {42, "shot"}, {0, "first"}, {0, "second"}, {42, "jumped"}}, []KV{{0, "second"}, {-42, "han"}, {42, "jumped"}}},
	}

	for _, tc := range testCases {
		bst := NewEmptyRedBlackBST()
		for _, d := range tc.data {
			bst.Put(d.key, d.val)
		}

		for _, w := range tc.want {
			v, err := bst.Get(w.key)

			if err != nil {
				t.Errorf("Got error %v, want key %v with value %v", err, w.key, w.val)
			}

			if v != w.val {
				t.Errorf("Got %v, want %v for key %v", v, w.val, w.key)
			}
		}
	}
}

func TestRedBlackBSTBalancing(t *testing.T) {
	testCases := []struct {
		data     []KV
		treeLvls int
	}{
		{[]KV{{83, "a"}, {69, "a"}, {65, "a"}, {82, "a"}, {67, "a"}, {72, "a"}, {88, "a"}, {77, "a"}}, 4},
		{[]KV{{83, "a"}, {69, "a"}, {65, "a"}, {82, "a"}, {67, "a"}, {72, "a"}, {88, "a"}, {77, "a"}, {80, "a"}, {76, "a"}}, 4},
		{[]KV{{83, "a"}, {69, "a"}, {65, "a"}, {82, "a"}, {67, "a"}, {72, "a"}, {88, "a"}, {77, "a"}, {83, "a"}, {69, "a"}, {65, "a"}, {82, "a"}, {67, "a"}, {72, "a"}, {88, "a"}, {77, "a"}}, 4},
		{[]KV{{65, "a"}, {67, "a"}, {69, "a"}, {72, "a"}, {76, "a"}, {77, "a"}, {80, "a"}, {82, "a"}, {83, "a"}, {88, "a"}}, 4},
		{[]KV{{65, "a"}}, 1},
		{[]KV{{}}, 1},
	}

	for _, tc := range testCases {
		bst := NewEmptyRedBlackBST()
		for _, d := range tc.data {
			bst.Put(d.key, d.val)
		}

		lvls := countTreeLevels(bst.Root)
		if lvls != tc.treeLvls {
			t.Errorf("Got tree with %v levels, want tree with %v levels", lvls, tc.treeLvls)
		}
	}
}
