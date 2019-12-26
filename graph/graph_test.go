package graph

import (
	"fmt"
)

func ExampleGraph() {
	graph := New()
	node0 := graph.AddNode()
	node1 := graph.AddNode()
	node2 := graph.AddNode()

	graph.AddEdge(node0, node1, 1)

	fmt.Println(node0)
	fmt.Println(node1)
	fmt.Println(node2)
	fmt.Println(graph.Nodes())
	graph.AddEdge(node0, node2, 5)
	graph.Neighbors(node0)
	// Output:
	// 0
	// 1
	// 2
	// [0 1 2]
}
