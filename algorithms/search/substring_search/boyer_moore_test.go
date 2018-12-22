package substring_search

import "testing"

func TestBoyerMoore(t *testing.T) {
	type Data struct {
		text     string
		pattern  string
		position int
	}

	testCases := []Data{
		{"AABBCC", "BB", 2},
		{"ABABAC", "BAC", 3},
		{"AABRAACADABRAACAADABRA", "AACAA", 12},
		{"Shellsort", "sort", 5},
		{"Shell", "Shells", 5},
		{"Hello", "Hola", 5},
	}

	for _, tc := range testCases {
		bm := BoyerMooreConstructor(tc.pattern)
		pos := bm.Search(tc.text)
		if pos != tc.position {
			t.Errorf("Got position %v for pattern %v in text %v, want %v\n", pos, tc.pattern, tc.text, tc.position)
		}
	}
}
