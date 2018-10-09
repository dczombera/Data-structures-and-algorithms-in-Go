package merge

func less(i, j Key) bool {
	return i.Compare(j) < 0
}

type Key int

func (this Key) Compare(that Key) int {
	if this > that {
		return 1
	} else if this < that {
		return -1
	}
	return 0
}

func Sort(cc []Key) {
	aux := make([]Key, len(cc))
	sort(cc, aux, 0, len(cc)-1)
}

func sort(cc, aux []Key, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := (hi + lo) / 2
	sort(cc, aux, lo, mid)
	sort(cc, aux, mid+1, hi)
	merge(cc, aux, lo, mid, hi)
}

func merge(cc, aux []Key, lo, mid, hi int) {
	for k := lo; k <= hi; k++ {
		aux[k] = cc[k]
	}

	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			cc[k] = aux[j]
			j++
		} else if j > hi {
			cc[k] = aux[i]
			i++
		} else if less(aux[j], aux[i]) {
			cc[k] = aux[j]
			j++
		} else {
			cc[k] = aux[i]
			i++
		}
	}
}
