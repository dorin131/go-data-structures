package linkedlist

func ExampleNew() {
	linkedList := New()
	linkedList.Append(1).Append(2).Append(3).PrintAll()
	// Output:
	// 1
	// 2
	// 3
}

func ExampleNewWithDeletion() {
	linkedList := New()
	linkedList.Append(1).Append(2).Append(3).DeleteWithValue(2).PrintAll()
	// Output:
	// 1
	// 3
}

func ExampleNewWithDeletionOfHead() {
	linkedList := New()
	linkedList.Append(1).Append(2).Append(3).DeleteWithValue(1).PrintAll()
	// Output:
	// 2
	// 3
}

func ExampleNewWithDeletionOfTail() {
	linkedList := New()
	linkedList.Append(1).Append(2).Append(3).DeleteWithValue(3).PrintAll()
	// Output:
	// 1
	// 2
}
