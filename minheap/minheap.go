/*
Package minheap provides an implementation of a Min Heap data structure
*/
package minheap

import (
	"log"

	"github.com/dorin131/go-data-structures/heap"
)

// MinHeap : represents a Min Heap data structure
type MinHeap struct {
	*heap.Heap
}

// New : returns a new instance of a Heap
func New(input []int) *MinHeap {
	h := &MinHeap{
		&heap.Heap{
			Size:  len(input),
			Items: input,
		},
	}

	if h.Size > 0 {
		h.buildMinHeap()
	}

	return h
}

// ExtractMin : removes top element of the heap and returns it
func (h *MinHeap) ExtractMin() int {
	if h.Size == 0 {
		log.Fatal("No items in the heap")
	}
	minItem := h.Items[0]
	lastIndex := h.Size - 1
	h.Items[0] = h.Items[lastIndex]

	// storing minimum at the end of the slice
	h.Items[lastIndex] = minItem

	h.minHeapifyDown(0)

	h.Size--

	return minItem
}

// Insert : adds a new element to the heap
func (h *MinHeap) Insert(item int) *MinHeap {
	h.Items = append(h.Items, item)
	h.Size++
	lastElementIndex := h.Size - 1
	h.minHeapifyUp(lastElementIndex)

	return h
}

// buildMinHeap : given a slice, arrange the elements so that
// they satisfy the Min Heap properties
func (h *MinHeap) buildMinHeap() {
	for i := h.Size / 2; i >= 0; i-- {
		h.minHeapifyUp(i)
	}
}

// minHeapifyDown : takes the top element and moves it down until
// the Min Heap properties are satisfied
func (h *MinHeap) minHeapifyDown(index int) {
	for (h.HasLeft(index) && h.Items[index] > h.Left(index)) ||
		(h.HasRight(index) && h.Items[index] > h.Right(index)) {
		if h.HasLeft(index) && h.Items[index] > h.Left(index) {
			h.Swap(index, h.GetLeftIndex(index))
			index = h.GetLeftIndex(index)
		} else {
			h.Swap(index, h.GetRightIndex(index))
			index = h.GetRightIndex(index)
		}
	}
}

// minHeapUp : takes the last element and moves it up until
// the Min Heap properties are satisfied
func (h *MinHeap) minHeapifyUp(index int) {
	for h.HasParent(index) && h.Parent(index) > h.Items[index] {
		h.Swap(h.GetParentIndex(index), index)
		index = h.GetParentIndex(index)
	}
}
