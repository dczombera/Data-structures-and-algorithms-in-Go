package priority_queue

import (
	"errors"
)

type Comparator interface {
	Compare(Key) int
}

type Key int

func (this Key) Compare(that Key) int {
	if this > that {
		return 1
	} else if this < that {
		return -1
	}
	return 0
}

type MaxPriorityQueue struct {
	pq []Key
	n  int
}

func NewEmptyMaxPriorityQueue() MaxPriorityQueue {
	return MaxPriorityQueue{[]Key{-1}, 0}
}

func NewMaxPriorityQueue(k Key) MaxPriorityQueue {
	return MaxPriorityQueue{[]Key{-1, k}, 1}
}

func (pq *MaxPriorityQueue) Size() int {
	return pq.n
}

func (pq *MaxPriorityQueue) IsEmpty() bool {
	return pq.n == 0
}

func (pq *MaxPriorityQueue) swim(k int) {
	for k > 1 && pq.less(k/2, k) {
		pq.exch(k, k/2)
		k = k / 2
	}
}

func (pq *MaxPriorityQueue) sink(k int) {
	for k*2 <= pq.n {
		i := k * 2
		if i < pq.n && pq.less(i, i+1) {
			i++
		}
		if !pq.less(k, i) {
			break
		}
		pq.exch(k, i)
		k = i
	}
}

func (pq *MaxPriorityQueue) less(i, j int) bool {
	return pq.pq[i].Compare(pq.pq[j]) < 0
}

func (pq *MaxPriorityQueue) exch(i, j int) {
	pq.pq[i], pq.pq[j] = pq.pq[j], pq.pq[i]
}

func (pq *MaxPriorityQueue) Insert(i Key) {
	pq.n++
	pq.pq = append(pq.pq, i)
	pq.swim(pq.n)
}

func (pq *MaxPriorityQueue) Max() Key {
	return pq.pq[1]
}

func (pq *MaxPriorityQueue) DelMax() (Key, error) {
	if pq.IsEmpty() {
		return Key(-1), errors.New("Priority queue is empty!")
	}
	k := pq.pq[1]
	pq.exch(1, pq.n)
	pq.pq = append(pq.pq[:pq.n])
	pq.n--
	pq.sink(1)
	return k, nil
}
