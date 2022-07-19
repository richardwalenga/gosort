package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/richardwalenga/gosort/sorters"
)

func getRandoms(capacity, maxExclusive int) []int {
	rand.Seed(time.Now().UnixNano())
	randoms := make([]int, capacity)
	for i := 0; i < capacity; i++ {
		randoms[i] = rand.Intn(maxExclusive)
	}
	return randoms
}

func main() {
	const Capacity int = 20000
	const MaxExclusive int = 100000
	insertion := sorters.InsertionSort[int]
	mySorters := map[string]func([]int){
		"Bubble":          sorters.BubbleSort[int],
		"Cocktail Shaker": sorters.CocktailShakerSort[int],
		"Selection":       sorters.SelectionSort[int],
		"Heap":            sorters.HeapSort[int],
		"Insertion":       insertion,
		"Merge":           sorters.GetMergeSort(insertion),
		"Quick":           sorters.QuickSort[int]}
	for sortName, sortFunc := range mySorters {
		randoms := getRandoms(Capacity, MaxExclusive)
		start := time.Now().UnixMilli()
		sortFunc(randoms)
		elapsed := time.Now().UnixMilli() - start
		var result string
		if sorters.IsSorted(randoms) {
			result = "successfully"
		} else {
			result = "unsuccessfully"
		}
		fmt.Printf("%s Sort finished %s in %d milliseconds\n",
			sortName, result, elapsed)
	}
}
