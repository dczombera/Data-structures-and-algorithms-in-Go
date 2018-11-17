package priority_queue

import (
	"errors"
	"fmt"
)

// IndexMinPriorityQueue is an indexed minimum priority queue using an indexed binary heap as underlying data structure
type IndexMinPriorityQueue struct {
	pq      []int
	qp      []int
	weights []*Weight
	size    int
	maxSize int
}

type Weight float64

func NewIndexMinPriorityQueue(maxSize int) *IndexMinPriorityQueue {
	adjustedSize := maxSize + 1
	pq := initIntSliceWithValue(adjustedSize, -1)
	qp := initIntSliceWithValue(adjustedSize, -1)
	return &IndexMinPriorityQueue{pq, qp, make([]*Weight, adjustedSize), 0, adjustedSize}
}

func (pq *IndexMinPriorityQueue) Empty() bool {
	return pq.size == 0
}

func (pq *IndexMinPriorityQueue) Contains(i int) bool {
	pq.validateIndex(i)
	if pq.qp[i] == -1 {
		return false
	}
	return true
}

func (pq *IndexMinPriorityQueue) Size() int {
	return pq.size
}

func (pq *IndexMinPriorityQueue) Insert(i int, weight Weight) {
	pq.validateIndex(i)
	if pq.Contains(i) {
		panic(fmt.Sprintf("Priority queue already contains given index %v", i))
	}
	pq.size++
	pq.pq[pq.size] = i
	pq.qp[i] = pq.size
	pq.weights[i] = &weight
	pq.swim(pq.size)
}

func (pq *IndexMinPriorityQueue) MinIndex() int {
	if pq.Empty() {
		panic("Priority queue is empty")
	}
	return pq.pq[1]
}

func (pq *IndexMinPriorityQueue) MinWeight() Weight {
	if pq.Empty() {
		panic("Priority queue is empty")
	}
	return *pq.weights[pq.pq[1]]
}

func (pq *IndexMinPriorityQueue) DelMin() (int, error) {
	if pq.Empty() {
		return -1, errors.New("Couldn't delete the minimum, priority queue is empty")
	}
	min := pq.pq[1]
	pq.exch(1, pq.size)
	pq.pq[pq.size] = -1
	pq.qp[min] = -1
	pq.weights[min] = nil
	pq.size--
	pq.sink(1)
	return min, nil
}

func (pq *IndexMinPriorityQueue) WeightOf(i int) Weight {
	pq.validateIndex(i)
	return *pq.weights[i]
}

func (pq *IndexMinPriorityQueue) ChangeWeight(i int, weight Weight) {
	pq.validateExistence(i)

	pq.weights[i] = &weight
	pq.swim(pq.qp[i])
	pq.sink(pq.qp[i])
}

func (pq *IndexMinPriorityQueue) DecreaseWeight(i int, weight Weight) {
	pq.validateExistence(i)
	if weight >= *pq.weights[i] {
		panic(fmt.Sprintf("Called DecreaseWeight with larger weight (%v) than currently exising in queue (%v)", weight, *pq.weights[i]))
	}
	pq.weights[i] = &weight
	pq.swim(pq.qp[i])
}

func (pq *IndexMinPriorityQueue) IncreaseWeight(i int, weight Weight) {
	pq.validateExistence(i)
	if weight <= *pq.weights[i] {
		panic(fmt.Sprintf("Called IncreaseWeight with smaller weight (%v) than currently exising in queue (%v)", weight, *pq.weights[i]))
	}
	pq.weights[i] = &weight
	pq.sink(pq.qp[i])
}

func (pq *IndexMinPriorityQueue) Delete(i int) {
	pq.validateExistence(i)
	index := pq.qp[i]
	pq.exch(index, pq.size)
	pq.pq[pq.size] = -1
	pq.qp[i] = -1
	pq.weights[i] = nil
	pq.size--
	pq.swim(index)
	pq.sink(index)
}

func (pq *IndexMinPriorityQueue) swim(i int) {
	for i > 1 && pq.greater(i/2, i) {
		pq.exch(i, i/2)
		i /= 2
	}
}

func (pq *IndexMinPriorityQueue) sink(i int) {
	for i*2 <= pq.size {
		j := i * 2
		if j < pq.size && pq.greater(j, j+1) {
			j++
		}
		if pq.greater(j, i) {
			break
		}
		pq.exch(i, j)
		i = j
	}
}

func (pq *IndexMinPriorityQueue) exch(i, j int) {
	pq.pq[i], pq.pq[j] = pq.pq[j], pq.pq[i]
	pq.qp[pq.pq[i]] = i
	pq.qp[pq.pq[j]] = j
}

func (pq *IndexMinPriorityQueue) greater(i, j int) bool {
	return pq.weights[pq.pq[i]].Compare(*pq.weights[pq.pq[j]]) > 0
}

func (pq *IndexMinPriorityQueue) validateIndex(i int) {
	if i < 0 || i >= pq.maxSize {
		panic("Index out of bounds")
	}
}

func (pq *IndexMinPriorityQueue) validateExistence(i int) {
	if !pq.Contains(i) {
		panic(fmt.Sprintf("Priority queue does not contain index %v", i))
	}
}

func initIntSliceWithValue(size, val int) []int {
	ii := make([]int, size)
	for i := 0; i < size; i++ {
		ii[i] = val
	}
	return ii
}

func (this Weight) Compare(that Weight) int {
	if this > that {
		return 1
	} else if this < that {
		return -1
	}
	return 1
}
