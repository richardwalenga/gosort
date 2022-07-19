package sorters

import "golang.org/x/exp/constraints"

const (
	RootIndex int = 1
	Down      int = 0
	Up        int = 1
)

type Heap[T constraints.Ordered] struct {
	size     *int
	storage  []T
	comparer func(T, T) bool
}

type HeapNode[T constraints.Ordered] struct {
	isRoot bool
	index  int
	heap   *Heap[T]
}

func newHeapNode[T constraints.Ordered](heap *Heap[T], index int) *HeapNode[T] {
	node := &HeapNode[T]{isRoot: index == RootIndex, index: index, heap: heap}
	return node
}

func getValue[T constraints.Ordered](node *HeapNode[T]) T {
	return node.heap.storage[node.index]
}

func setValue[T constraints.Ordered](node *HeapNode[T], newVal T) {
	node.heap.storage[node.index] = newVal
}

func heapifyDown[T constraints.Ordered](node *HeapNode[T]) {
	left, right := left(node), right(node)
	if left == nil && right == nil {
		return
	}

	other := right
	if left != nil && right != nil {
		/* Favor the smallest or largest child node as a swap partner
		 * depending on if one is working with a min or max heap.
		 * The comparer will return true if the first value meets this
		 * criteria. */
		if node.heap.comparer(getValue(left), getValue(right)) {
			other = left
		}
	} else if left != nil {
		other = left
	}
	trySwapValue(node, other, Down)
}

func heapifyUp[T constraints.Ordered](node *HeapNode[T]) {
	parent := parent(node)
	if parent == nil {
		return
	}
	trySwapValue(node, parent, Up)
}

func left[T constraints.Ordered](node *HeapNode[T]) *HeapNode[T] {
	return fromIndex(node, 2*node.index)
}

func right[T constraints.Ordered](node *HeapNode[T]) *HeapNode[T] {
	return fromIndex(node, 2*node.index+1)
}

func parent[T constraints.Ordered](node *HeapNode[T]) *HeapNode[T] {
	var parent *HeapNode[T] = nil
	if !node.isRoot {
		parent = newHeapNode(node.heap, node.index/2)
	}
	return parent
}

func fromIndex[T constraints.Ordered](node *HeapNode[T], index int) *HeapNode[T] {
	var ret *HeapNode[T] = nil
	if !isOutOfRange(node.heap, index) {
		ret = newHeapNode(node.heap, index)
	}
	return ret
}

func trySwapValue[T constraints.Ordered](node, other *HeapNode[T], direction int) {
	if node == nil || other == nil {
		return
	}

	val, otherVal := getValue(node), getValue(other)
	comparer := node.heap.comparer
	if direction == Down && comparer(otherVal, val) {
		setValue(node, otherVal)
		setValue(other, val)
		heapifyDown(other)
	} else if direction == Up && comparer(val, otherVal) {
		setValue(node, otherVal)
		setValue(other, val)
		heapifyUp(other)
	}
}

func minComparer[T constraints.Ordered](x, y T) bool {
	return x < y
}

func maxComparer[T constraints.Ordered](x, y T) bool {
	return x > y
}

func newHeap[T constraints.Ordered](isMin bool, capacity int) *Heap[T] {
	size := 0
	heap := &Heap[T]{storage: make([]T, capacity+1)}
	heap.size = &size
	if isMin {
		heap.comparer = minComparer[T]
	} else {
		heap.comparer = maxComparer[T]
	}
	return heap
}

func store[T constraints.Ordered](heap *Heap[T], num T) {
	*heap.size++
	heap.storage[*heap.size] = num
	settingRoot := *heap.size == RootIndex
	if !settingRoot {
		added := newHeapNode(heap, *heap.size)
		heapifyUp(added)
	}
}

func isOutOfRange[T constraints.Ordered](heap *Heap[T], index int) bool {
	return index > *heap.size
}

func take[T constraints.Ordered](heap *Heap[T]) T {
	taken := heap.storage[RootIndex]
	heap.storage[RootIndex] = heap.storage[*heap.size]
	*heap.size--
	if *heap.size > 1 {
		rootNode := newHeapNode(heap, RootIndex)
		heapifyDown(rootNode)
	}
	return taken
}

func HeapSort[T constraints.Ordered](toSort []T) {
	count := len(toSort)
	heap := newHeap[T](true, count)
	i := 0
	for ; i < count; i++ {
		store(heap, toSort[i])
	}
	for i = 0; i < count; i++ {
		toSort[i] = take(heap)
	}
}
