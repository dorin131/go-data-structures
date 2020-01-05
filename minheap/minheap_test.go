package minheap

import (
	"fmt"
	"testing"
)

func ExampleMinHeap() {
	h := New([]int{})
	h.Insert(333).Insert(-3).Insert(0).Insert(1000).Insert(-44).Insert(5)
	fmt.Println(h.ExtractMin())
	// Output:
	// -44
}

func TestMinHeap(t *testing.T) {
	tests := []struct {
		initial  []int
		toAdd    []int
		expected int
	}{
		{[]int{}, []int{1, 2, 3, 4, 5}, 1},
		{[]int{}, []int{300, 5, 77, -8, 0, 50}, -8},
		{[]int{}, []int{-1000, 1000}, -1000},
		{[]int{}, []int{1000, -1000}, -1000},
		{[]int{0, 7, 10}, []int{1, 2, 3, 4, 5}, 0},
		{[]int{100}, []int{300, 5, 77, -8, 0, 50}, -8},
		{[]int{-2000, 0, 800}, []int{-1000, 1000}, -2000},
		{[]int{5000, 10000}, []int{1000, -1000}, -1000},
	}
	for i, test := range tests {
		h := New(test.initial)
		for _, n := range test.toAdd {
			h.Insert(n)
		}
		result := h.ExtractMin()
		if result != test.expected {
			t.Errorf("[%d] Expected: %v. Got: %v", i, test.expected, result)
		}
	}
}
