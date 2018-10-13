package binary_search

import (
	"errors"
)

type Key int

type Value string

type BinarySearchST struct {
	keys   []Key
	values []Value
	n      int
}

var initSize = 2

func NewEmptyBinarySearchST() BinarySearchST {
	return BinarySearchST{make([]Key, initSize), make([]Value, initSize), 0}
}

func NewBinarySearchST(key Key, val Value) BinarySearchST {
	return BinarySearchST{[]Key{key}, []Value{val}, 1}
}

func (this Key) CompareTo(that Key) int {
	if this > that {
		return 1
	} else if this < that {
		return -1
	}
	return 0
}

func (bs *BinarySearchST) Put(key Key, val Value) {
	r := bs.Rank(key)
	if r < bs.n && key.CompareTo(bs.keys[r]) == 0 {
		bs.values[r] = val
		return
	}
	if bs.n == cap(bs.keys) {
		bs.resize(cap(bs.keys) * 2)
	}
	for i := bs.n; i > r; i-- {
		bs.keys[i] = bs.keys[bs.n-1]
		bs.values[i] = bs.values[bs.n-1]
	}
	bs.keys[r] = key
	bs.values[r] = val
	bs.n++
}

func (bs *BinarySearchST) Get(key Key) (Value, error) {
	r := bs.Rank(key)
	if r < bs.n && key.CompareTo(bs.keys[r]) == 0 {
		return bs.values[r], nil
	}
	return Value(-1), errors.New("Key not found.")
}

func (bs *BinarySearchST) Rank(key Key) int {
	lo := 0
	hi := bs.n - 1
	for lo <= hi {
		mid := (lo + hi) / 2
		cmp := key.CompareTo(bs.keys[mid])
		if cmp < 0 {
			hi = mid - 1
		} else if cmp > 0 {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return lo
}

func (bs *BinarySearchST) Delete(key Key) {
	r := bs.Rank(key)
	if r < bs.n && key.CompareTo(bs.keys[r]) == 0 {
		bs.keys = append(bs.keys[:r], bs.keys[r+1:]...)
		bs.values = append(bs.values[:r], bs.values[r+1:]...)
		bs.n--
	}
}

func (bs *BinarySearchST) Size() int {
	return bs.n
}

func (bs *BinarySearchST) Contains(key Key) bool {
	if _, err := bs.Get(key); err != nil {
		return false
	}
	return true
}

func (bs *BinarySearchST) IsEmpty() bool {
	return bs.n == 0
}

func (bs *BinarySearchST) resize(size int) {
	n := (cap(bs.keys) + 1) * size
	tempK := make([]Key, n)
	tempV := make([]Value, n)
	copy(tempK, bs.keys)
	copy(tempV, bs.values)
	bs.keys = tempK
	bs.values = tempV
}
