package sort

type MyInt int

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
