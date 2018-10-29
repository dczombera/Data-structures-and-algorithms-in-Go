package linear_probing_hash_table

import (
	"errors"
	"hash/fnv"
)

type LinearProbingHT struct {
	keys        []*Key
	values      []*Value
	sizeTable   int
	sizeKVPairs int
}

var init_cap = 8

type Key string
type Value int

func NewEmptyLinearProbingHT() LinearProbingHT {
	return LinearProbingHT{make([]*Key, init_cap), make([]*Value, init_cap), init_cap, 0}
}

func NewSizedLinearProbingHT(size int) LinearProbingHT {
	return LinearProbingHT{make([]*Key, size), make([]*Value, size), size, 0}
}

func (ht *LinearProbingHT) Get(k Key) (Value, error) {
	for i := ht.hash(k); ht.keys[i] != nil; i = (i + 1) % ht.sizeTable {
		if *ht.keys[i] == k {
			return *ht.values[i], nil
		}
	}
	return Value(0), errors.New("Key not found")
}

func (ht *LinearProbingHT) Put(k Key, v Value) {
	i := ht.hash(k)
	for ; ht.keys[i] != nil; i = (i + 1) % ht.sizeTable {
		if *ht.keys[i] == k {
			ht.values[i] = &v
			return
		}
	}
	ht.keys[i] = &k
	ht.values[i] = &v
	ht.sizeKVPairs++

	// double table if 50% full or more
	if ht.sizeKVPairs >= ht.sizeTable/2 {
		ht.resize(ht.sizeTable * 2)
	}
}

func (ht *LinearProbingHT) Delete(k Key) {
	if !ht.Contains(k) {
		return
	}

	i := ht.hash(k)
	for *ht.keys[i] != k {
		i = (i + 1) % ht.sizeTable
	}
	ht.keys[i] = nil
	ht.values[i] = nil
	ht.sizeKVPairs--

	ht.rehashCluster((i + 1) % ht.sizeTable)

	// half size of table if it is 12.5% full or less
	if ht.sizeKVPairs > 0 && ht.sizeKVPairs*8 <= ht.sizeTable {
		ht.resize(ht.sizeTable / 2)
	}
}

func (ht *LinearProbingHT) rehashCluster(start int) {
	for i := start; ht.keys[i] != nil; i = (i + 1) % ht.sizeTable {
		keyToRehash := *ht.keys[i]
		valToRehash := *ht.values[i]
		ht.keys[i] = nil
		ht.values[i] = nil
		ht.sizeKVPairs--
		ht.Put(keyToRehash, valToRehash)
	}
}

func (ht *LinearProbingHT) Size() int {
	return ht.sizeKVPairs
}

func (ht *LinearProbingHT) IsEmpty() bool {
	return ht.Size() == 0
}

func (ht *LinearProbingHT) Contains(k Key) bool {
	_, err := ht.Get(k)
	if err != nil {
		return false
	}
	return true
}

func (ht *LinearProbingHT) hash(k Key) int {
	h := fnv.New32a()
	h.Write([]byte(k))
	val := h.Sum32() % uint32(ht.sizeTable)
	return int(val)
}

func (ht *LinearProbingHT) resize(size int) {
	tmp := NewSizedLinearProbingHT(size)
	for i := 0; i < ht.sizeTable; i++ {
		if ht.keys[i] != nil {
			tmp.Put(*ht.keys[i], *ht.values[i])
		}
	}
	ht.keys = tmp.keys
	ht.values = tmp.values
	ht.sizeKVPairs = tmp.sizeKVPairs
	ht.sizeTable = tmp.sizeTable
}
