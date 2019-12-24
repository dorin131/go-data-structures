package queue

import (
	"fmt"
)

func ExampleQueue() {
	queue := New()
	result, _ := queue.Queue(1).Queue(2).Queue(3).Peek()
	fmt.Println(result)
	fmt.Println(queue.IsEmpty())
	result, _ = queue.Dequeue()
	fmt.Println(result)
	result, _ = queue.Dequeue()
	fmt.Println(result)
	queue.Dequeue()
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
