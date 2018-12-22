package substring_search

type BoyerMoore struct {
	pattern string
	right   []int
	len     int
}

func BoyerMooreConstructor(pattern string) BoyerMoore {
	len := len(pattern)
	right := make([]int, R)
	for c := range right {
		right[c] = -1
	}
	for i, c := range pattern {
		right[c] = i
	}
	return BoyerMoore{pattern, right, len}
}

func (bm BoyerMoore) Search(text string) int {
	n := len(text)
	skip := 0
	for i := 0; i <= n-bm.len; i += skip {
		skip = 0
		for j := bm.len - 1; j >= 0; j-- {
			if bm.pattern[j] != text[i+j] {
				skip = j - bm.right[text[i+j]]
				if skip < 1 {
					skip = 1
				}
				break
			}
		}
		if skip == 0 {
			return i
		}
	}
	return n
}
