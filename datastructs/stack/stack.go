package stack

import (
	"errors"

	"github.com/dczomber/data-structs-and-algs/datastructs/node"
)

type stack struct {
	First *node.Node
	Size  int
}

func NewStack(n *node.Node) stack {
	return stack{n, 1}
}

func NewEmptyStack() stack {
	return stack{nil, 0}
}

func (s *stack) Push(i node.Item) {
	n := &node.Node{i, s.First}
	s.First = n
	s.Size++
}

func (s *stack) Pop() (node.Item, error) {
	var n *node.Node
	if s.IsEmpty() {
		return n.Item, errors.New("Stack is empty!")
	}

	n = s.First
	s.First = n.Next
	s.Size--
	return n.Item, nil
}

func (s *stack) Peek() (node.Item, error) {
	var n *node.Node
	if s.IsEmpty() {
		return n.Item, errors.New("Stack is empty!")
	}
	return s.First.Item, nil
}

func (s *stack) IsEmpty() bool {
	return s.Size == 0
}
