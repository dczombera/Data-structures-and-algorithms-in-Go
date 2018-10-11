package symbol_table

import "errors"

type sequentialSearchST struct {
	first *Node
	n     int
}

type Node struct {
	Key  Key
	Val  Value
	Next *Node
}

type Key int

type Value string

func NewEmptySequentialSearchST() sequentialSearchST {
	return sequentialSearchST{nil, 0}
}

func NewSequentialSearchST(n *Node) sequentialSearchST {
	return sequentialSearchST{n, 1}
}

func (this Key) equals(that Key) bool {
	return this == that
}

func (seq *sequentialSearchST) Put(key Key, val Value) {
	for curr := seq.first; curr != nil; curr = curr.Next {
		if key.equals(curr.Key) {
			curr.Val = val
			return
		}
	}
	node := &Node{key, val, seq.first}
	seq.first = node
	seq.n++
}

func (seq *sequentialSearchST) Get(key Key) (Value, error) {
	for curr := seq.first; curr != nil; curr = curr.Next {
		if key.equals(curr.Key) {
			return curr.Val, nil
		}
	}
	return Value(""), errors.New("key not found.")
}

func (seq *sequentialSearchST) Delete(k Key) {
	seq.delete(seq.first, k)
}

func (seq *sequentialSearchST) delete(n *Node, key Key) *Node {
	if n == nil {
		return nil
	}
	if key.equals(n.Key) {
		seq.n--
		return n.Next
	}
	n.Next = seq.delete(n.Next, key)
	return n
}

func (seq *sequentialSearchST) Contains(key Key) bool {
	for curr := seq.first; curr != nil; curr = curr.Next {
		if key.equals(curr.Key) {
			return true
		}
	}
	return false
}

func (seq *sequentialSearchST) IsEmpty() bool {
	return seq.n == 0
}

func (seq *sequentialSearchST) Size() int {
	return seq.n
}
