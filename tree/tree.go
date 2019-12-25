/*
Package tree provides a Tree data structure (Binary Search Tree)
We assume that every value is of type "int" and is unique
*/
package tree

import (
	"fmt"
)

// Node : represents a single node of a Tree data structure
type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

// New : constructor function for a Tree node
func New(data int) *Node {
	return &Node{
		Left:  nil,
		Right: nil,
		Data:  data,
	}
}

// Insert : adds a node to the tree
func (t *Node) Insert(data int) *Node {
	if data < t.Data {
		if t.Left != nil {
			t.Left.Insert(data)
		} else {
			t.Left = New(data)
		}
	} else {
		if t.Right != nil {
			t.Right.Insert(data)
		} else {
			t.Right = New(data)
		}
	}
	return t
}

// Contains : checks if the tree contains a node with a given value
func (t *Node) Contains(data int) bool {
	if t.Data == data {
		return true
	}
	if data < t.Data {
		if t.Left != nil {
			t.Left.Contains(data)
		} else {
			return false
		}
	} else {
		if t.Right != nil {
			t.Right.Contains(data)
		} else {
			return false
		}
	}
	return false
}

// Traverse : prints all elements of the tree
func (t *Node) Traverse() {
	if t.Left != nil {
		t.Left.Traverse()
	}
	fmt.Println(t.Data)
	if t.Right != nil {
		t.Right.Traverse()
	}
}
