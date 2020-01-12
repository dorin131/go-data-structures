/*
Package maxheap provides a Max Heap data structure
*/
package maxheap

import (
	"log"

	"github.com/dorin131/go-data-structures/heap"
)

// MaxHeap : represents a Max heap data structure
type MaxHeap struct {
	*heap.Heap
}

// New : returns a new instance of a Heap
func New(input []int) *MaxHeap {
	h := &MaxHeap{
		&heap.Heap{
			Items: input,
		},
	}

	if len(h.Items) > 0 {
		h.buildMaxHeap()
	}

	return h
}

func (h *MaxHeap) buildMaxHeap() {
	for i := len(h.Items)/2 - 1; i >= 0; i-- {
		h.maxHeapifyDown(i)
	}
}

// Insert : adds an element to the heap
func (h *MaxHeap) Insert(item int) *MaxHeap {
	h.Items = append(h.Items, item)
	lastElementIndex := len(h.Items) - 1
	h.maxHeapifyUp(lastElementIndex)

	return h
}

// ExtractMax : returns the maximum element and removes it from the Heap
func (h *MaxHeap) ExtractMax() int {
	if len(h.Items) == 0 {
		log.Fatal("No items in the heap")
	}
	minItem := h.Items[0]
	lastIndex := len(h.Items) - 1
	h.Items[0] = h.Items[lastIndex]

	// shrinking slice
	h.Items = h.Items[:len(h.Items)-1]

	h.maxHeapifyDown(0)

	return minItem
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.HasParent(index) && h.Parent(index) < h.Items[index] {
		h.Swap(h.GetParentIndex(index), index)
		index = h.GetParentIndex(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	// iterate until we have a child node smaller than the index value
	for (h.HasLeft(index) && h.Items[index] < h.Left(index)) ||
		(h.HasRight(index) && h.Items[index] < h.Right(index)) {
		// if both children are smaller
		if (h.HasLeft(index) && h.Items[index] < h.Left(index)) &&
			(h.HasRight(index) && h.Items[index] < h.Right(index)) {
			// compare the children and swap with the largest
			if h.Left(index) > h.Right(index) {
				h.Swap(index, h.GetLeftIndex(index))
				index = h.GetLeftIndex(index)
			} else {
				h.Swap(index, h.GetRightIndex(index))
				index = h.GetRightIndex(index)
			}
			// else if only left child is smaller than swap with it
		} else if h.HasLeft(index) && h.Items[index] < h.Left(index) {
			h.Swap(index, h.GetLeftIndex(index))
			index = h.GetLeftIndex(index)
			// else it must be the right child that is smaller, so we swap with it
		} else {
			h.Swap(index, h.GetRightIndex(index))
			index = h.GetRightIndex(index)
		}
	}
}
