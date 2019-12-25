package queue

import (
	"fmt"
)

func ExampleQueue() {
	queue := New()
	result, _ := queue.Enqueue(1).Enqueue(2).Enqueue(3).Peek()
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
