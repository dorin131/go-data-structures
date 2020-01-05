/*
Package maxheap provides a Max Heap data structure
*/
package maxheap

import "github.com/dorin131/go-data-structures/heap"

// MaxHeap : represents a Max heap data structure
type MaxHeap struct {
	*heap.Heap
}

// New : returns a new instance of a Heap
func New(input []int) *MaxHeap {
	h := &MaxHeap{
		&heap.Heap{
			Size:     len(input),
			Elements: len(input),
			Items:    input,
		},
	}

	if h.Elements > 0 {
		h.buildMaxHeap()
	}

	return h
}

func (h *MaxHeap) buildMaxHeap() {
	for i := h.Elements / 2; i >= 0; i-- {
		h.maxHeapifyDown(i)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {

}

// Insert : adds an element to the heap
func (h *MaxHeap) Insert(element int) *MaxHeap {
	return h
}

// ExtractMax : returns the maximum element and removes it from the Heap
func (h *MaxHeap) ExtractMax() int {
	return 0
}
