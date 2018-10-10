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

func Sort(kk []Key) {
	aux := make([]Key, len(kk))
	sort(kk, aux, 0, len(kk)-1)
}

func sort(kk, aux []Key, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := (hi + lo) / 2
	sort(kk, aux, lo, mid)
	sort(kk, aux, mid+1, hi)
	merge(kk, aux, lo, mid, hi)
}

func merge(kk, aux []Key, lo, mid, hi int) {
	for k := lo; k <= hi; k++ {
		aux[k] = kk[k]
	}

	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			kk[k] = aux[j]
			j++
		} else if j > hi {
			kk[k] = aux[i]
			i++
		} else if less(aux[j], aux[i]) {
			kk[k] = aux[j]
			j++
		} else {
			kk[k] = aux[i]
			i++
		}
	}
}
