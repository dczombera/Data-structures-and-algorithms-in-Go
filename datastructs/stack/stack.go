package stack

import (
	"errors"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/node"
)

type Stack struct {
	First *node.Node
	Size  int
}

func NewStack(n *node.Node) Stack {
	return Stack{n, 1}
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
	var n *node.Node
	if s.IsEmpty() {
		return n.Item, errors.New("Stack is empty!")
	}

	n = s.First
	s.First = n.Next
	s.Size--
	return n.Item, nil
}

func (s *Stack) Peek() (int, error) {
	var n *node.Node
	if s.IsEmpty() {
		return n.Item, errors.New("Stack is empty!")
	}
	return s.First.Item, nil
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}
