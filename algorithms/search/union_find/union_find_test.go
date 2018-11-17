package union_find

import "testing"

type testCase struct {
	sites        int
	union        [][]int
	connected    [][]int
	disconnected [][]int
	count        int
}

var testCases = []testCase{
	{
		sites:        10,
		union:        [][]int{{4, 3}, {3, 8}, {6, 5}, {9, 4}, {2, 1}, {8, 9}, {5, 0}, {7, 2}, {6, 1}, {1, 0}, {6, 7}},
		connected:    [][]int{{7, 0}, {3, 9}, {1, 2}, {2, 5}, {9, 8}, {4, 8}},
		disconnected: [][]int{{2, 3}, {8, 7}, {1, 9}, {4, 0}, {5, 3}},
		count:        2,
	},
}

func TestUnionFind(t *testing.T) {
	for _, tc := range testCases {
		uf := NewUnionFind(tc.sites)
		for _, u := range tc.union {
			uf.Union(u[0], u[1])
		}

		for _, c := range tc.connected {
			if !uf.Connected(c[0], c[1]) {
				t.Errorf("Got %v and %v disconnected, want them to be connected", c[0], c[1])
			}
		}

		for _, d := range tc.disconnected {
			if uf.Connected(d[0], d[1]) {
				t.Errorf("Got %v and %v connected, want them to be disconnected", d[0], d[1])
			}
		}
	}
}
