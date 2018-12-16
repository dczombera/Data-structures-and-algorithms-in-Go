package ternary_search_trie

import "errors"

type TST struct {
	root *Node
}

type Node struct {
	char  byte
	val   *Value
	left  *Node
	mid   *Node
	right *Node
}

type Value int

func Constructor() *TST {
	return &TST{}
}

func (tst TST) Get(key string) (int, error) {
	n := tst.get(tst.root, key, 0)
	if n == nil {
		return -1, errors.New("Key not found")
	}
	return int(*n.val), nil
}

func (tst TST) get(n *Node, key string, pos int) *Node {
	if n == nil {
		return nil
	}
	char := key[pos]
	if char < n.char {
		return tst.get(n.left, key, pos)
	} else if char > n.char {
		return tst.get(n.right, key, pos)
	}

	if pos == len(key)-1 {
		return n
	}

	return tst.get(n.mid, key, pos+1)
}

func (tst *TST) Put(key string, val int) {
	tst.root = tst.put(tst.root, key, val, 0)
}

func (tst *TST) put(n *Node, key string, val, pos int) *Node {
	char := key[pos]
	if n == nil {
		n = &Node{char, nil, nil, nil, nil}
	}
	if char < n.char {
		n.left = tst.put(n.left, key, val, pos)
	} else if char > n.char {
		n.right = tst.put(n.right, key, val, pos)
	} else if pos < len(key)-1 {
		n.mid = tst.put(n.mid, key, val, pos+1)
	} else {
		value := Value(val)
		n.val = &value
	}

	return n
}
