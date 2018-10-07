package sort

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
