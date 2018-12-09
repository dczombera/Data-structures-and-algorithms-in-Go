package three_way_string_quicksort

import "testing"

type testCase struct {
	data []string
	want []string
}

var testCases = []testCase{
	{[]string{"aabb", "bbaa", "abab", "baba", "aaa", "bb", "aa", "b", "a", "z"}, []string{"a", "aa", "aaa", "aabb", "abab", "b", "baba", "bb", "bbaa", "z"}},
	{[]string{"a", "b", "c", "d", "aa"}, []string{"a", "aa", "b", "c", "d"}},
	{[]string{"z", "x", "y"}, []string{"x", "y", "z"}},
	{[]string{"zz", "x", "yy", "z", "xx", "y"}, []string{"x", "xx", "y", "yy", "z", "zz"}},
}

func TestThreeWayStringQuicksort(t *testing.T) {
	for _, tc := range testCases {
		Sort(tc.data)
		for i := 0; i < len(tc.data); i++ {
			if tc.data[i] != tc.want[i] {
				t.Errorf("Got %v, want %v", tc.data[i], tc.want[i])
			}
		}
	}
}
