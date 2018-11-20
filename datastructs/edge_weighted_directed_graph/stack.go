package edge_weighted_directed_graph

import "errors"

type Stack struct {
	First *Node
	Size  int
}

type Node struct {
	Item DirectedEdge
	Next *Node
}

func NewStack(e DirectedEdge) Stack {
	n := &Node{e, nil}
	return Stack{n, 1}
}

func (s *Stack) Push(e DirectedEdge) {
	n := &Node{e, s.First}
	s.First = n
	s.Size++
}

func (s *Stack) Pop() (DirectedEdge, error) {
	if s.IsEmpty() {
		return DirectedEdge{}, errors.New("Cannot pop item. Stack is empty")
	}
	e := s.First.Item
	s.First = s.First.Next
	s.Size--
	return e, nil
}

func (s *Stack) Peek() (DirectedEdge, error) {
	if s.IsEmpty() {
		return DirectedEdge{}, errors.New("Cannot peek. Stack is empty")
	}
	return s.First.Item, nil
}

func (s Stack) IsEmpty() bool {
	return s.Size == 0
}
