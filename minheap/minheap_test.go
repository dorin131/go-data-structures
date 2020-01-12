package minheap

import (
	"fmt"
	"sort"
	"testing"
)

func ExampleMinHeap() {
	h := New([]int{100, 500})
	h.Insert(333).Insert(-3).Insert(0).Insert(1000).Insert(-44).Insert(5)
	fmt.Println(h.ExtractMin())
	// Output:
	// -44
}

func TestMinHeap(t *testing.T) {
	tests := []struct {
		initial []int
		toAdd   []int
	}{
		{[]int{}, []int{4, 9, 10, 0, -4, 7}},
		{[]int{}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{300, 5, 77, -8, 0, 50}},
		{[]int{}, []int{-1000, 1000}},
		{[]int{}, []int{1000, -1000}},
		{[]int{4, 9, 10, 0, -4, 7}, []int{}},
		{[]int{0, 7, 10}, []int{1, 2, 3, 4, 5}},
		{[]int{100}, []int{300, 5, 77, -8, 0, 50}},
		{[]int{-2000, 0, 800}, []int{-1000, 1000}},
		{[]int{5000, 10000}, []int{1000, -1000}},
	}
	for i, test := range tests {
		h := New(test.initial)
		for _, n := range test.toAdd {
			h.Insert(n)
		}

		checkMinHeap(i, h, t)
	}
}

func checkMinHeap(n int, h *MinHeap, t *testing.T) {
	result := []int{}
	correctlySorted := make([]int, len(h.Items))
	copy(correctlySorted, h.Items)

	length := len(h.Items)
	for i := 0; i < length; i++ {
		result = append(result, h.ExtractMin())
	}
	sort.Slice(correctlySorted, func(i, j int) bool {
		return correctlySorted[i] < correctlySorted[j]
	})

	for k := range correctlySorted {
		if correctlySorted[k] != result[k] {
			t.Errorf("[%v] Expected: %v, Got: %v\n", n, correctlySorted, result)
			return
		}
	}
	fmt.Printf("[%v] Correctly sorted: %v\n", n, result)
}
