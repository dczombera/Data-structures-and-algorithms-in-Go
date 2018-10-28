package separate_chaining_hash_table

import (
	"errors"
	"hash/fnv"
)

type Node struct {
	Key  Key
	Val  Value
	Next *Node
}

type Key string
type Value int

var init_cap int = 8

type SeparateChainingHT struct {
	table     []*Node
	sizeTable int
	sizePairs int
}

func NewNode(k Key, v Value, next *Node) *Node {
	return &Node{k, v, next}
}

func NewEmptySeparateChainingHT() SeparateChainingHT {
	return SeparateChainingHT{make([]*Node, init_cap), init_cap, 0}
}

func NewSizedSeparateChainingHT(size int) SeparateChainingHT {
	return SeparateChainingHT{make([]*Node, size), size, 0}
}

func (ht *SeparateChainingHT) Get(k Key) (Value, error) {
	h := ht.hash(k)
	for curr := ht.table[h]; curr != nil; curr = curr.Next {
		if curr.Key == k {
			return curr.Val, nil
		}
	}
	return Value(0), errors.New("Key not found")
}

func (ht *SeparateChainingHT) Put(k Key, v Value) {
	h := ht.hash(k)
	first := ht.table[h]
	for curr := first; curr != nil; curr = curr.Next {
		if curr.Key == k {
			curr.Val = v
			return
		}
	}
	ht.table[h] = NewNode(k, v, first)
	ht.sizePairs++

	// resize table if avg. length of list is >= 10
	if ht.sizePairs >= ht.sizeTable*10 {
		ht.resize(2 * ht.sizeTable)
	}
}

func (ht *SeparateChainingHT) Delete(k Key) {
	h := ht.hash(k)
	ht.table[h] = ht.delete(ht.table[h], k)
	// resize if avg.length of list is <= 2
	if ht.sizeTable > init_cap && ht.sizePairs <= ht.sizeTable*2 {
		ht.resize(ht.sizeTable / 2)
	}
}

func (ht *SeparateChainingHT) delete(n *Node, k Key) *Node {
	if n == nil {
		return nil
	}
	if n.Key == k {
		ht.sizePairs--
		return n.Next
	}
	n.Next = ht.delete(n.Next, k)
	return n
}

func (ht *SeparateChainingHT) Size() int {
	return ht.sizePairs
}

func (ht *SeparateChainingHT) IsEmpty() bool {
	return ht.Size() == 0
}

func (ht *SeparateChainingHT) Contains(k Key) bool {
	_, err := ht.Get(k)
	if err != nil {
		return false
	}
	return true
}

func (ht *SeparateChainingHT) hash(k Key) int {
	h := fnv.New32a()
	h.Write([]byte(k))
	val := h.Sum32() % uint32(ht.sizeTable)
	return int(val)
}

func (ht *SeparateChainingHT) resize(size int) {
	tmp := NewSizedSeparateChainingHT(size)
	for _, curr := range ht.table {
		for ; curr != nil; curr = curr.Next {
			tmp.Put(curr.Key, curr.Val)
		}
	}
	ht.table = tmp.table
	ht.sizeTable = tmp.sizeTable
	ht.sizePairs = tmp.sizePairs
}
