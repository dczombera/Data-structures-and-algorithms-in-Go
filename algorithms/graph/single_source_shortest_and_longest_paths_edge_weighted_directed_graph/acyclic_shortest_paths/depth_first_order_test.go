package acyclic_shortest_paths

import (
	"testing"

	graph "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_directed_graph"
)

type testCase struct {
	graphSize        int
	edges            []graph.DirectedEdge
	preOrder         []int
	postOrder        []int
	reversePostOrder []int
	prePos           []int
	postPos          []int
	reversePostPos   []int
}

var testCases = []testCase{
	{
		graphSize:        4,
		edges:            []graph.DirectedEdge{{0, 1, 0.1}, {1, 2, 1.2}, {2, 3, 2.3}},
		preOrder:         []int{0, 1, 2, 3},
		postOrder:        []int{3, 2, 1, 0},
		reversePostOrder: []int{0, 1, 2, 3},
		prePos:           []int{0, 1, 2, 3},
		postPos:          []int{3, 2, 1, 0},
		reversePostPos:   []int{0, 1, 2, 3},
	},
	{
		graphSize:        13,
		edges:            []graph.DirectedEdge{{0, 5, 0.5}, {5, 4, 5.4}, {0, 1, 0.1}, {0, 6, 0.6}, {6, 9, 6.9}, {9, 11, 9.11}, {9, 10, 9.10}, {11, 12, 11.12}, {2, 3, 2.3}, {8, 7, 8.7}},
		preOrder:         []int{0, 5, 4, 1, 6, 9, 11, 12, 10, 2, 3, 7, 8},
		postOrder:        []int{4, 5, 1, 12, 11, 10, 9, 6, 0, 3, 2, 7, 8},
		reversePostOrder: []int{8, 7, 2, 3, 0, 6, 9, 10, 11, 12, 1, 5, 4},
		prePos:           []int{0, 3, 9, 10, 2, 1, 4, 11, 12, 5, 8, 6, 7},
		postPos:          []int{8, 2, 10, 9, 0, 1, 7, 11, 12, 6, 5, 4, 3},
		reversePostPos:   []int{4, 10, 2, 3, 12, 11, 5, 1, 0, 6, 7, 8, 9},
	},
}

func TestDepthFirstOrder(t *testing.T) {
	for _, tc := range testCases {
		digraph := graph.NewEdgeWeightedDigraph(tc.graphSize)
		for _, e := range tc.edges {
			digraph.AddEdge(e)
		}

		dfo := NewDepthFirstOrder(digraph)
		preOrder := dfo.PreOrder()
		curr := preOrder.First
		for _, want := range tc.preOrder {
			if curr.Item != want {
				t.Errorf("Got vertex %v in preorder iteration, want %v", curr.Item, want)
			}
			curr = curr.Next
		}

		postOrder := dfo.PostOrder()
		curr = postOrder.First
		for _, want := range tc.postOrder {
			if curr.Item != want {
				t.Errorf("Got vertex %v in postorder iteration, want %v", curr.Item, want)
			}
			curr = curr.Next
		}

		reversePostOrder := dfo.ReversePostOrder()
		curr = reversePostOrder.First
		for _, want := range tc.reversePostOrder {
			if curr.Item != want {
				t.Errorf("Got vertex %v in reverse postorder iteration, want %v", curr.Item, want)
			}
			curr = curr.Next
		}

		for i, pos := range tc.prePos {
			if dfo.PrePos(i) != pos {
				t.Errorf("Got pre-position %v for vertex %v, want %v", dfo.PrePos(i), i, pos)
			}
		}

		for i, pos := range tc.postPos {
			if dfo.PostPos(i) != pos {
				t.Errorf("Got post-position %v for vertex %v, want %v", dfo.PostPos(i), i, pos)
			}
		}

		for i, pos := range tc.reversePostPos {
			if dfo.ReversePostPos(i) != pos {
				t.Errorf("Got reverse-position %v for vertex %v, want %v", dfo.ReversePostPos(i), i, pos)
			}
		}
	}
}
