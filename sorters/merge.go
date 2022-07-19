package sorters

import "golang.org/x/exp/constraints"

func GetMergeSort[T constraints.Ordered](smallSorter func([]T)) func([]T) {
	return func(toSort []T) {
		sortImpl(toSort, len(toSort), smallSorter)
	}
}

func copySlice[T any](slice []T) []T {
	copied := make([]T, len(slice))
	copy(copied, slice)
	return copied
}

func sortImpl[T constraints.Ordered](ary []T, count int, smallSorter func([]T)) {
	if count < 10 {
		smallSorter(ary)
		return
	}

	mid := count / 2
	// Unlike Python, a slice doesn't make a copy.
	first, second := copySlice(ary[0:mid]), copySlice(ary[mid:])
	firstCount, secondCount := len(first), len(second)
	sortImpl(first, firstCount, smallSorter)
	sortImpl(second, secondCount, smallSorter)
	firstIndex, secondIndex := 0, 0
	for i := 0; i < count; i++ {
		canTakeFirst, canTakeSecond := firstIndex < firstCount, secondIndex < secondCount
		if canTakeFirst && (!canTakeSecond || first[firstIndex] <= second[secondIndex]) {
			ary[i] = first[firstIndex]
			if firstIndex < firstCount {
				firstIndex++
			}
		} else {
			ary[i] = second[secondIndex]
			if secondIndex < secondCount {
				secondIndex++
			}
		}
	}
}
