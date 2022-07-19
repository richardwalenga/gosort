package sorters

import "golang.org/x/exp/constraints"

func QuickSort[T constraints.Ordered](toSort []T) {
	sortBetweenIndexes(toSort, 0, len(toSort)-1)
}

// Organizes the values between the high and low indexes where the
// chosen pivot is moved to a new index where all values greater than
// the pivot are to its right. The new index for the pivot is returned.
func partition[T constraints.Ordered](ary []T, low, high int) int {
	pivot := ary[high]
	// initialize the index below low because the index is guaranteed
	// to be incremented before the pivot is moved to its new home.
	newPivotIndex := low - 1
	for i := low; i < high; i++ {
		if ary[i] <= pivot {
			newPivotIndex++
			SwapValues(ary, newPivotIndex, i)
		}
	}
	// There will always be at least one swap call since if this is the
	// first time, it means every value checked is greater than the pivot.
	newPivotIndex++
	SwapValues(ary, newPivotIndex, high)
	return newPivotIndex
}

func sortBetweenIndexes[T constraints.Ordered](ary []T, low, high int) {
	if low < high {
		pivotIndex := partition(ary, low, high)
		sortBetweenIndexes(ary, low, pivotIndex-1)
		sortBetweenIndexes(ary, pivotIndex+1, high)
	}
}
