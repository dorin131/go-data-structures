package tree

import "fmt"

func ExampleTreeInsert() {
	tree := New(5)

	tree.Insert(3).Insert(8).Insert(-5).Insert(200).Traverse()
	// Output:
	// -5
	// 3
	// 5
	// 8
	// 200
}

func ExampleTreeContains() {
	tree := New(5)

	tree.Insert(3).Insert(8).Insert(-5).Insert(200)

	fmt.Println(tree.Contains(5))
	// Output:
	// true
}

func ExampleTreeDiesntContain() {
	tree := New(5)

	tree.Insert(3).Insert(8).Insert(-5).Insert(200)

	fmt.Println(tree.Contains(0))
	// Output:
	// false
}
