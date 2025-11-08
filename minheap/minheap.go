package minheap

import (
	"cmp"
)

// Define the generic Heap structure.
// T is a type parameter constrained by cmp.Ordered, meaning it must support
// standard comparison operators (<, >, ==, etc.).
type MinHeap[T cmp.Ordered] []T

// Len returns the number of elements.
func (h MinHeap[T]) Len() int {
	return len(h)
}

// Less reports whether the element at index i should sort before the element at index j.
// For a MIN heap, we return true if h[i] is less than h[j].
func (h MinHeap[T]) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the elements at indexes i and j.
func (h MinHeap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds x to the heap.
// Note: x must be asserted to type T.
func (h *MinHeap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

// Pop removes and returns the minimum element (the root) from the heap.
func (h *MinHeap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
