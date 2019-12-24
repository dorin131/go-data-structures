package stack

import (
	"fmt"
)

func ExampleStack() {
	stack := New()
	peekResult, _ := stack.Push(1).Push(2).Peek()
	fmt.Println(peekResult)
	popResult, _ := stack.Pop()
	fmt.Println(popResult)
	fmt.Println(stack.IsEmpty())
	popResult, _ = stack.Pop()
	fmt.Println(popResult)
	fmt.Println(stack.IsEmpty())
	_, err := stack.Pop()
	fmt.Println(err)
	// Output:
	// 2
	// 2
	// false
	// 1
	// true
	// Stack is empty
}
