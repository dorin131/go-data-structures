/*
Package heap provides an implementation of a Min Heap data structure
*/
package heap

import "log"

// Heap : contains a slice which holds the underlying heap data
type Heap struct {
	items []int
}

// New : returns a new instance of a Heap
func New() *Heap {
	return &Heap{
		items: []int{},
	}
}

func (h *Heap) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h *Heap) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *Heap) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *Heap) hasLeftChild(index int) bool {
	return h.getLeftChildIndex(index) < len(h.items)
}

func (h *Heap) hasRightChild(index int) bool {
	return h.getRightChildIndex(index) < len(h.items)
}

func (h *Heap) hasParent(index int) bool {
	return h.getParentIndex(index) >= 0
}

func (h *Heap) leftChild(index int) int {
	return h.items[h.getLeftChildIndex(index)]
}

func (h *Heap) rightChild(index int) int {
	return h.items[h.getRightChildIndex(index)]
}

func (h *Heap) parent(index int) int {
	return h.items[h.getParentIndex(index)]
}

func (h *Heap) swap(indexOne, indexTwo int) {
	h.items[indexOne], h.items[indexTwo] = h.items[indexTwo], h.items[indexOne]
}

// Min : returns the top element of the heap
func (h *Heap) Min() int {
	if len(h.items) == 0 {
		log.Fatal("No items in the heap")
	}
	return h.items[0]
}

// ExtractMin : removes top element of the heap and returns it
func (h *Heap) ExtractMin() int {
	if len(h.items) == 0 {
		log.Fatal("No items in the heap")
	}
	item := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]

	h.bubbleDown()

	return item
}

// Insert : adds a new element to the heap
func (h *Heap) Insert(item int) *Heap {
	h.items = append(h.items, item)

	h.bubbleUp()

	return h
}

func (h *Heap) bubbleDown() {
	index := 0
	for (h.hasLeftChild(index) && h.items[index] > h.leftChild(index)) ||
		(h.hasRightChild(index) && h.items[index] > h.rightChild(index)) {
		if h.hasLeftChild(index) && h.items[index] > h.leftChild(index) {
			h.swap(index, h.getLeftChildIndex(index))
			index = h.getLeftChildIndex(index)
		} else {
			h.swap(index, h.getRightChildIndex(index))
			index = h.getRightChildIndex(index)
		}
	}
}

func (h *Heap) bubbleUp() {
	index := len(h.items) - 1
	for h.hasParent(index) && h.parent(index) > h.items[index] {
		h.swap(h.getParentIndex(index), index)
		index = h.getParentIndex(index)
	}
}
