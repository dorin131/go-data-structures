/*
Package heap provides helper methods for a Min Heap or a Max Heap
Not to be used on its own.
*/
package heap

// Heap : contains a slice which holds the underlying heap data
type Heap struct { // the total number of elements, which doesn't go down on extract
	Items []int
}

// GetLeftIndex : get left index of a Heap node
func (h *Heap) GetLeftIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

// GetRightIndex : get right index of a Heap node
func (h *Heap) GetRightIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

// GetParentIndex : get parent index of a Heap node
func (h *Heap) GetParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

// HasLeft : check if node at index has left node
func (h *Heap) HasLeft(index int) bool {
	return h.GetLeftIndex(index) < len(h.Items)
}

// HasRight : check if node at index has right node
func (h *Heap) HasRight(index int) bool {
	return h.GetRightIndex(index) < len(h.Items)
}

// HasParent : check if node at index has parent node
func (h *Heap) HasParent(index int) bool {
	return h.GetParentIndex(index) >= 0
}

// Left : get left node value, given an index
func (h *Heap) Left(index int) int {
	return h.Items[h.GetLeftIndex(index)]
}

// Right : get right node value, given an index
func (h *Heap) Right(index int) int {
	return h.Items[h.GetRightIndex(index)]
}

// Parent : get parent node value, given an index
func (h *Heap) Parent(index int) int {
	return h.Items[h.GetParentIndex(index)]
}

// Swap : swap values of two nodes at specified indeces
func (h *Heap) Swap(indexOne, indexTwo int) {
	h.Items[indexOne], h.Items[indexTwo] = h.Items[indexTwo], h.Items[indexOne]
}
