package sorters

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](toSort []T) {
	count := len(toSort)
	if count < 2 {
		return
	}

	for leftToRightSort(toSort, count) {
	}
}

func CocktailShakerSort[T constraints.Ordered](toSort []T) {
	count := len(toSort)
	if count < 2 {
		return
	}

	comparers := []func(ary []T, count int) bool{
		leftToRightSort[T], rightToLeftSort[T]}
	// By applying a bitmask of 1 less than a power of 2, I can cleanly
	// alternate sorting left to right followed by right to left.
	const Bitmask int = 1
	for i := 0; comparers[i](toSort, count); i = (i + 1) & Bitmask {
	}
}

func leftToRightSort[T constraints.Ordered](ary []T, count int) bool {
	swapped := false
	for i := 1; i < count; i++ {
		if ary[i-1] > ary[i] {
			SwapValues(ary, i-1, i)
			swapped = true
		}
	}
	return swapped
}

func rightToLeftSort[T constraints.Ordered](ary []T, count int) bool {
	swapped := false
	for i := count - 1; i > 0; i-- {
		if ary[i] < ary[i-1] {
			SwapValues(ary, i-1, i)
			swapped = true
		}
	}
	return swapped
}
