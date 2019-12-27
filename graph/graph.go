/*
Package graph provides a Graph data structure (directed & weighted)
*/
package graph

// Graph : represents a Graph
type Graph struct {
	nodes []*GraphNode
}

// GraphNode : represents a Graph node
type GraphNode struct {
	id    int
	edges map[int]int
}

// New : returns a new instance of a Graph
func New() *Graph {
	return &Graph{
		nodes: []*GraphNode{},
	}
}

// AddNode : adds a new node to the Graph
func (g *Graph) AddNode() (id int) {
	id = len(g.nodes)
	g.nodes = append(g.nodes, &GraphNode{
		id:    id,
		edges: make(map[int]int),
	})
	return
}

// AddEdge : adds a directional edge together with a weight
func (g *Graph) AddEdge(n1, n2 int, w int) {
	g.nodes[n1].edges[n2] = w
}

// Neighbors : returns a list of node IDs that are linked to this node
func (g *Graph) Neighbors(id int) []int {
	neighbors := []int{}
	for _, node := range g.nodes {
		for edge := range node.edges {
			if node.id == id {
				neighbors = append(neighbors, edge)
			}
			if edge == id {
				neighbors = append(neighbors, node.id)
			}
		}
	}
	return neighbors
}

// Nodes : returns a list of node IDs
func (g *Graph) Nodes() []int {
	nodes := make([]int, len(g.nodes))
	for i := range g.nodes {
		nodes[i] = i
	}
	return nodes
}

// Edges : returns a list of edges with weights
func (g *Graph) Edges() [][3]int {
	edges := make([][3]int, 0, len(g.nodes))
	for i := range g.nodes {
		for k, v := range g.nodes[i].edges {
			edges = append(edges, [3]int{i, k, int(v)})
		}
	}
	return edges
}
