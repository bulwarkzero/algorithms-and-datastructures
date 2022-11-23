package main

import (
	"fmt"
	graph "new-way/4-1-1-graph"
)

/*
0___1___4___5
|          |
2___3_______/
*/
func initializeGraph() *graph.Graph {
	g := graph.NewGraph()
	n0 := graph.NewNode(0, "A") //A
	n1 := graph.NewNode(1, "B") //B
	n2 := graph.NewNode(2, "C") //C
	n3 := graph.NewNode(3, "D") //D

	g.AddNode(n0)
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)

	g.AddEdge(n0, n1)
	g.AddEdge(n0, n2)
	g.AddEdge(n1, n2)
	g.AddEdge(n2, n0)
	g.AddEdge(n2, n3)
	g.AddEdge(n3, n3)

	return g
}

func main() {
	g := initializeGraph()

	fmt.Println(g)

	g.Traverse(g.Nodes()[2], func(n *graph.Node) {
		fmt.Println(n.Value())
	})
}
