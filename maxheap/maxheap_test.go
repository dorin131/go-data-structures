package maxheap

import (
	"fmt"
	"testing"
)

func ExampleMaxHeap() {
	h := New([]int{100, 500})
	h.Insert(333).Insert(-3).Insert(0).Insert(1000).Insert(-44).Insert(5)
	fmt.Println(h.ExtractMax())
	// Output:
	// 1000
}

func TestMaxHeap(t *testing.T) {
	tests := []struct {
		initial     []int
		toAdd       []int
		expectedMax int
	}{
		{[]int{}, []int{1, 2, 3, 4, 5}, 5},
		{[]int{}, []int{300, 5, 77, -8, 0, 50}, 300},
		{[]int{}, []int{-1000, 1000}, 1000},
		{[]int{}, []int{1000, -1000}, 1000},
		{[]int{0, 7, 10}, []int{1, 2, 3, 4, 5}, 10},
		{[]int{100}, []int{300, 5, 77, -8, 0, 50}, 300},
		{[]int{-2000, 0, 800}, []int{-1000, 1000}, 1000},
		{[]int{5000, 10000}, []int{1000, -1000}, 10000},
	}
	for i, test := range tests {
		h := New(test.initial)
		for _, n := range test.toAdd {
			h.Insert(n)
		}
		result := h.ExtractMax()
		if result != test.expectedMax {
			t.Errorf("[%d] Expected Max to be: %v. Got: %v", i, test.expectedMax, result)
		}
	}
}
