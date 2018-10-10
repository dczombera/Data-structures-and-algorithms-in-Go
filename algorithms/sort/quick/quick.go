package quick

type Key int

func (this Key) Compare(that Key) int {
	if this > that {
		return 1
	} else if this < that {
		return -1
	}
	return 0
}

func Sort(kk []Key) {
	sort(kk, 0, len(kk)-1)
}

func sort(kk []Key, lo, hi int) {
	if hi <= lo {
		return
	}

	j := partition(kk, lo, hi)
	sort(kk, lo, j-1)
	sort(kk, j+1, hi)
}

func partition(kk []Key, lo, hi int) int {
	v := kk[lo]
	i := lo
	j := hi + 1
	for true {
		for {
			i++
			if !less(kk[i], v) || i == hi {
				break
			}
		}

		for {
			j--
			if !less(v, kk[j]) || j == lo {
				break
			}
		}

		if i >= j {
			break
		}

		exch(kk, i, j)
	}

	exch(kk, lo, j)
	return j
}

func exch(kk []Key, i, j int) {
	kk[i], kk[j] = kk[j], kk[i]
}

func less(i, j Key) bool {
	return i.Compare(j) < 0
}
