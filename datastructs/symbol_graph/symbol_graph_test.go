package symbol_graph

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	t.Log(os.Getwd())
}

type testCase struct {
	wantEdges  []VertexEdges
	wantDegree []VertexDegree
}

type VertexEdges struct {
	vertex int
	adj    []int
}

type VertexDegree struct {
	vertex int
	degree int
}

var testCases = []testCase{
	{
		wantEdges:  []VertexEdges{{0, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49}}},
		wantDegree: []VertexDegree{{0, 49}},
	},
	{
		wantEdges:  []VertexEdges{{50, []int{51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 13, 63, 15, 64, 65, 21, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 46, 91, 92, 93, 94, 95, 96}}},
		wantDegree: []VertexDegree{{50, 50}},
	},
}

func isSameListOfVertices(got, want []int) bool {
	for i, w := range want {
		if w != got[i] {
			return false
		}
	}
	return true
}

func TestSymbolGraph(t *testing.T) {
	sg := NewSymbolGraph("fixtures/movies.txt", "/")
	for _, tc := range testCases {
		for _, we := range tc.wantEdges {
			adj := sg.Graph().AdjacencyList(we.vertex)
			if !isSameListOfVertices(adj, we.adj) {
				t.Errorf("Adjacency list not include correct vertices")
			}
		}

		for _, wd := range tc.wantDegree {
			degree := sg.Graph().Degree(wd.vertex)
			if degree != wd.degree {
				t.Errorf("Got degree of %v for vertex %v, want degree of %v", degree, wd.vertex, wd.degree)
			}
		}
	}
}

func TestMultipleOccurenceOfSameVertex(t *testing.T) {
	tc := struct {
		name string
		adj  []int
	}{
		"Ford, Harrison",
		[]int{0, 50},
	}

	sg := NewSymbolGraph("fixtures/movies.txt", "/")
	index, err := sg.IndexOf(tc.name)
	if err != nil {
		t.Errorf("Got error %v, want index of actor %v", err, tc.name)
	}

	adj := sg.Graph().AdjacencyList(index)
	if !isSameListOfVertices(tc.adj, adj) {
		t.Errorf("Adjacency list of %v does not include correct vertices", tc.name)
	}
}
