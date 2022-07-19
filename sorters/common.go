package sorters

import "golang.org/x/exp/constraints"

func AreIdentical[T comparable](x, y []T) bool {
	xLength, yLength := len(x), len(y)
	if xLength != yLength {
		return false
	}
	for i := 0; i < xLength; i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func IsSorted[T constraints.Ordered](ary []T) bool {
	count := len(ary)
	if count < 2 {
		return true
	}
	for i := 0; i < count-1; i++ {
		if ary[i] > ary[i+1] {
			return false
		}
	}
	return true
}

func SwapValues[T any](ary []T, x, y int) {
	if x == y {
		return
	}
	ary[x], ary[y] = ary[y], ary[x]
}
