package depth_first_order

import (
	"testing"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/directed_graph"
)

type testCase struct {
	graphSize        int
	edges            [][]int
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
		edges:            [][]int{{0, 1}, {1, 2}, {2, 3}},
		preOrder:         []int{0, 1, 2, 3},
		postOrder:        []int{3, 2, 1, 0},
		reversePostOrder: []int{0, 1, 2, 3},
		prePos:           []int{0, 1, 2, 3},
		postPos:          []int{3, 2, 1, 0},
		reversePostPos:   []int{0, 1, 2, 3},
	},
	{
		graphSize:        13,
		edges:            [][]int{{0, 5}, {5, 4}, {0, 1}, {0, 6}, {6, 9}, {9, 11}, {9, 10}, {11, 12}, {2, 3}, {8, 7}},
		preOrder:         []int{0, 5, 4, 1, 6, 9, 11, 12, 10, 2, 3, 7, 8},
		postOrder:        []int{4, 5, 1, 12, 11, 10, 9, 6, 0, 3, 2, 7, 8},
		reversePostOrder: []int{8, 7, 2, 3, 0, 6, 9, 10, 11, 12, 1, 5, 4},
		prePos:           []int{0, 3, 9, 10, 2, 1, 4, 11, 12, 5, 8, 6, 7},
		postPos:          []int{12, 9, 3, 2, 10, 11, 8, 1, 0, 7, 4, 6, 5},
		reversePostPos:   []int{5, 6, 4, 7, 0, 1, 8, 11, 10, 2, 3, 9, 12},
	},
}

func TestDepthFirstOrder(t *testing.T) {
	for _, tc := range testCases {
		digraph := directed_graph.NewDigraph(tc.graphSize)
		for _, e := range tc.edges {
			digraph.AddEdge(e[0], e[1])
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
	}
}
