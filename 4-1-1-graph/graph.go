package graph

import (
	"fmt"
	queue "new-way/4-queue"
)

type Node struct {
	value int
	label string
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) Label() string {
	return n.label
}

func NewNode(value int, label string) *Node {
	return &Node{value, label}
}

type Graph struct {
	nodes []*Node
	edges map[*Node][]*Node
}

func (g *Graph) Nodes() []*Node {
	return g.nodes
}

func (g *Graph) Edges() map[*Node][]*Node {
	return g.edges
}

func (g *Graph) AddNode(node *Node) {
	g.nodes = append(g.nodes, node)
}

func (g *Graph) AddEdge(source, target *Node) {
	if g.edges == nil {
		g.edges = make(map[*Node][]*Node)
	}

	if sourceEdges, ok := g.edges[source]; ok {
		g.edges[source] = append(sourceEdges, target)
	} else {
		g.edges[source] = []*Node{target}
	}

	if targetEdges, ok := g.edges[target]; ok {
		g.edges[target] = append(targetEdges, source)
	} else {
		g.edges[target] = []*Node{source}
	}
}

func (g *Graph) IsEmpty() bool {
	return len(g.nodes) < 1
}

func (g *Graph) String() string {
	str := ""

	for _, node := range g.nodes {
		str += fmt.Sprintf("Node %d -> ", node.value)
		if nodeEdges, ok := g.edges[node]; ok {
			for _, edge := range nodeEdges {
				str += fmt.Sprintf("%d ", edge.value)
			}
		}
		str += "\n"
	}

	return str
}

// bfs
func (g *Graph) Traverse(startNode *Node, f func(*Node)) {
	if g.IsEmpty() {
		return
	}

	q := queue.NewArrayBasedQueue()

	node := g.nodes[0]
	if startNode != nil {
		node = startNode
	}

	q.Enqueue(node)

	visited := make(map[*Node]bool)

	for q.Len() > 0 {
		node = q.Dequeue().(*Node)
		visited[node] = true

		if edges, ok := g.edges[node]; ok {
			for _, edge := range edges {
				if _, ok := visited[edge]; !ok {
					q.Enqueue(edge)
					visited[edge] = true
				}
			}
		}

		if f != nil {
			f(node)
		}
	}
}

func NewGraph() *Graph {
	return &Graph{}
}
