package sort

type MyInt int

func Selection(ii []MyInt) {
	for i := 0; i < len(ii); i++ {
		min := i

		for j := i + 1; j < len(ii); j++ {
			if ii[j].Compare(ii[min]) < 0 {
				min = j
			}
		}

		swap(i, min, ii)
	}
}

func (this MyInt) Compare(that MyInt) int {
	if this > that {
		return 1
	} else if this < that {
		return -1
	}
	return 0
}

func swap(x, y int, ii []MyInt) {
	ii[x], ii[y] = ii[y], ii[x]
}
