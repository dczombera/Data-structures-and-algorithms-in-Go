package sort

func Insertion(ii []MyInt) {
	for i := 1; i < len(ii); i++ {
		for j := i; j > 0; j-- {
			if ii[j].Compare(ii[j-1]) > 0 {
				break
			}
			swap(j, j-1, ii)
		}
	}
}
