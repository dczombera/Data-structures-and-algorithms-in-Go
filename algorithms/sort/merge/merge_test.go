package merge

import "testing"

type testCase struct {
	data []Key
	want []Key
}

var testCases = []testCase{
	{[]Key{5, 4, 3, 2, 1}, []Key{1, 2, 3, 4, 5}},
	{[]Key{42, -42}, []Key{-42, 42}},
	{[]Key{3, 5, 3, 5, -5, -3, 0, -3}, []Key{-5, -3, -3, 0, 3, 3, 5, 5}},
	{[]Key{-1, 0, -1}, []Key{-1, -1, 0}},
	{[]Key{}, []Key{}},
}

func TestMerge(t *testing.T) {
	for _, tc := range testCases {
		Sort(tc.data)
		for i := 0; i < len(tc.data); i++ {
			if tc.data[i] != tc.want[i] {
				t.Errorf("Got %v, want %v", tc.data[i], tc.want[i])
			}
		}
	}
}
