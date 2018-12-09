package three_way_string_quicksort

func Sort(ss []string) {
	sort(ss, 0, len(ss)-1, 0)
}

func sort(ss []string, lo, hi, c int) {
	if hi <= lo {
		return
	}

	lt, i, gt := lo, lo+1, hi
	v := charAt(ss[lo], c)
	for gt >= i {
		w := charAt(ss[i], c)
		if w < v {
			swap(ss, i, lt)
			i++
			lt++
		} else if w > v {
			swap(ss, i, gt)
			gt--
		} else {
			i++
		}
	}
	sort(ss, lo, lt-1, c)
	if v >= 0 {
		sort(ss, lt, gt, c+1)
	}
	sort(ss, gt+1, hi, c)
}

func charAt(s string, c int) int {
	if len(s) <= c {
		return -1
	}
	return int(s[c])
}

func swap(ss []string, i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
