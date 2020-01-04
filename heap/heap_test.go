package heap

import (
	"fmt"
	"testing"
)

func ExampleHeap() {
	h := New()
	h.Insert(333).Insert(-3).Insert(0).Insert(1000).Insert(-44).Insert(5)
	fmt.Println(h.Min())
	// Output:
	// -44
}

func TestHeap(t *testing.T) {
	tests := []struct {
		given    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{300, 5, 77, -8, 0, 50}, -8},
		{[]int{-1000, 1000}, -1000},
		{[]int{1000, -1000}, -1000},
	}
	for i, test := range tests {
		h := New()
		for _, n := range test.given {
			h.Insert(n)
		}
		result := h.Min()
		fmt.Println(result)
		if result != test.expected {
			t.Errorf("[%d] Expected: %v. Got: %v", i, test.expected, result)
		}
	}
}
