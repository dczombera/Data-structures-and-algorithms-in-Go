package substring_search

import (
	"testing"
)

func TestRabinKarp(t *testing.T) {
	var R uint64 = 256
	type Data struct {
		text     string
		pattern  string
		position int
		R        uint64
	}

	testCases := []Data{
		{"AABBCC", "BB", 2, R},
		{"ABABAC", "BAC", 3, R},
		{"AABRAACADABRAACAADABRA", "AACAA", 12, R},
		{"Shellsort", "sort", 5, R},
		{"Shell", "Shells", 5, R},
		{"Hello", "Hola", 5, R},
	}

	for _, tc := range testCases {
		rk := RabinKarpConstructor(tc.pattern, tc.R)
		pos := rk.Search(tc.text)
		if pos != tc.position {
			t.Errorf("Got position %v for pattern %v in text %v, want %v\n", pos, tc.pattern, tc.text, tc.position)
		}
	}
}
