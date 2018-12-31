package regular_expression_pattern_matching

import "testing"

func TestNFA(t *testing.T) {
	type Data struct {
		pattern string
		texts   []string
	}

	testCases := []Data{
		{"((A*B|AC)D)", []string{"AAAABD", "AABD", "BD", "ACD"}},
		{"(.*(A*B|AC)D.*)", []string{"ABCCBD", "AAABACDCAB"}},
		{"(0|1(01*0)*10*)", []string{"11", "1001", "1100"}},
		{"gcg(cgg|agg)*ctg", []string{"gcgaggaggcggcggctg", "gcgctg", "gcgcggctg", "gcgaggctg"}},
		{"(.*NEEDLE.*)", []string{"A NEEDLE IN A HAYSTACK", "NEEDLE", "NEEDLqweWENEEDLEBCQNEE"}},
	}

	for _, tc := range testCases {
		nfa := NFAConstructor(tc.pattern)
		for _, txt := range tc.texts {
			if !nfa.Recognizes(txt) {
				t.Errorf("Text %v was not recognized by NFA for pattern %v, want it to be recognized", txt, tc.pattern)
			}
		}
	}
}

func TestNFAError(t *testing.T) {
	type Data struct {
		pattern string
		texts   []string
	}

	testCases := []Data{
		{"((A*B|AC)D)", []string{"AAABCD", "BC", "D", "A", "BAAAD"}},
		{"(0|1(01*0)*1)", []string{"10", "1011", "10000"}},
		{"gcg(cgg|agg)*ctg", []string{"gcg", "ctg", "cgggcgctg"}},
		{"(.*NEEDLE.*)", []string{"A NEEDEL IN A HAYSTACK", "NEEDLLE", "NEEDLqweWENEDLEBCQNEE"}},
	}

	for _, tc := range testCases {
		nfa := NFAConstructor(tc.pattern)
		for _, txt := range tc.texts {
			if nfa.Recognizes(txt) {
				t.Errorf("Text %v was recognized by NFA for pattern %v, want it to be unrecognized", txt, tc.pattern)
			}
		}
	}
}
