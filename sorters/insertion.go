package sorters

import "golang.org/x/exp/constraints"

func InsertionSort[T constraints.Ordered](toSort []T) {
	count := len(toSort)
	if count < 2 {
		return
	}

	for i := 1; i < count; i++ {
		j, old := i-1, toSort[i]
		for ; j >= 0 && toSort[j] > old; j-- {
			SwapValues(toSort, j, j+1)
		}
		shifted := j != i-1
		if shifted {
			// Need to compensate for the last decrement of j
			// in the loop above
			toSort[j+1] = old
		}
	}
}
