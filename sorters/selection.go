package sorters

import "golang.org/x/exp/constraints"

func SelectionSort[T constraints.Ordered](toSort []T) {
	count := len(toSort)
	if count < 2 {
		return
	}

	for i := 0; i < count-1; i++ {
		minIdx := i
		for j := i + 1; j < count; j++ {
			if toSort[minIdx] > toSort[j] {
				minIdx = j
			}
		}
		if i != minIdx {
			SwapValues(toSort, i, minIdx)
		}
	}
}
