/*
Package stack provides a slice-based Stack data structure
*/
package stack

import "fmt"

// Stack : data structure
type Stack struct {
	data []int
}

// New : returns a new instance of a stack
func New() *Stack {
	return &Stack{
		data: []int{},
	}
}

// IsEmpty : will return a boolean indicating whether there
// are any elements on the stack
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Peek : will return the element on the top of the stack
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

// Push : Adds an element on the stack
func (s *Stack) Push(n int) *Stack {
	s.data = append(s.data, n)
	return s
}

// Pop : removes an element from the stack and returns
// its value
func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("Stack is empty")
	}
	element := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return element, nil
}
