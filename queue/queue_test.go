package queue

import (
	"fmt"
)

func ExampleQueue() {
	queue := New()
	result, _ := queue.Add(1).Add(2).Add(3).Peek()
	fmt.Println(result)
	fmt.Println(queue.IsEmpty())
	result, _ = queue.Remove()
	fmt.Println(result)
	result, _ = queue.Remove()
	fmt.Println(result)
	queue.Remove()
	fmt.Println(queue.IsEmpty())
	_, err := queue.Peek()
	fmt.Println(err)
	// Output:
	// 1
	// false
	// 1
	// 2
	// true
	// Queue is empty
}
