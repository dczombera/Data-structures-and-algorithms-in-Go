package bst

import "errors"

type BST struct {
	Root *Node
}

type Node struct {
	Key   Key
	Val   Value
	Left  *Node
	Right *Node
	Size  int
}

type Key int
type Value string

func (this Key) CompareTo(that Key) int {
	if this > that {
		return 1
	} else if this < that {
		return -1
	}
	return 0
}

func (bst BST) Size() int {
	return bst.size(bst.Root)
}

func (bst BST) size(n *Node) int {
	if n == nil {
		return 0
	}
	return n.Size
}

func (bst BST) contains(k Key) bool {
	_, error := bst.Get(k)
	if error != nil {
		return false
	}
	return true
}

func (bst BST) Get(k Key) (Value, error) {
	return bst.get(bst.Root, k)
}

func (bst BST) get(n *Node, k Key) (Value, error) {
	if n == nil {
		return Value(""), errors.New("Key not found")
	}

	cmp := n.Key.CompareTo(k)
	if cmp < 0 {
		return bst.get(n.Left, k)
	} else if cmp > 0 {
		return bst.get(n.Right, k)
	}
	return n.Val, nil
}

// put

// del

// delMin

// delMax

// max

// min
