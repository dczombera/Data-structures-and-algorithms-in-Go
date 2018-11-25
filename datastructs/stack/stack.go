package stack

import (
	"errors"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/node"
)

type Stack struct {
	First *node.Node
	Size  int
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func NewEmptyStack() Stack {
	return Stack{nil, 0}
}

func (s *Stack) Push(i int) {
	n := &node.Node{i, s.First}
	s.First = n
	s.Size++
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("Stack is empty")
	}

	n := s.First
	s.First = s.First.Next
	s.Size--
	return n.Item, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("Stack is empty")
	}
	return s.First.Item, nil
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}
