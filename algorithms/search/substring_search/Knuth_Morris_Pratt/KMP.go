package substring_search

type KMP struct {
	dfa [][]int
	len int
}

var R int = 256

func KMPConstructor(pattern string) KMP {
	m := len(pattern)
	dfa := make([][]int, R)
	for i := range dfa {
		dfa[i] = make([]int, m)
	}
	dfa[pattern[0]][0] = 1
	for j, x := 1, 0; j < m; j++ {
		for c := 0; c < R; c++ {
			dfa[c][j] = dfa[c][x]
		}
		dfa[pattern[j]][j] = j + 1
		x = dfa[pattern[j]][x]
	}

	return KMP{dfa, m}
}

func (kmp KMP) Search(text string) int {
	n := len(text)
	i, j := 0, 0
	for ; i < n && j < kmp.len; i++ {
		j = kmp.dfa[text[i]][j]
	}
	if j == kmp.len {
		return i - kmp.len
	}
	return n
}
