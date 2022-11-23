package main

import (
	"fmt"
	graph "new-way/4-1-1-graph"
	union "new-way/6-union-find"
)

/*
created graph

A ------- B
|         |
|         |
|         |
C ------- D
|
|
|
E
*/
func initiateGraph() *graph.Graph {
	g := graph.NewGraph()
	n0 := graph.NewNode(0, "A") // A
	n1 := graph.NewNode(1, "B") // B
	n2 := graph.NewNode(2, "C") // C
	n3 := graph.NewNode(3, "D") // D
	n4 := graph.NewNode(3, "E") // E

	g.AddNode(n0)
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)

	g.AddEdge(n0, n1)
	g.AddEdge(n0, n2)
	g.AddEdge(n1, n3)
	g.AddEdge(n2, n3)
	g.AddEdge(n2, n4)

	return g
}

func KruskalMST(g *graph.Graph) (results [][]*graph.Node) {
	allEdges := g.Edges()
	graphNodes := g.Nodes()

	union := union.New(graphNodes)

	// our graph is not weighted, so we simply create a loop over
	for node, edges := range allEdges {
		// in minimum spanning tree, maximum results are less that equal to graph nodes
		if len(results) >= len(graphNodes) {
			break
		}

		for _, edge := range edges {
			// not in a same union
			if union.Find(node) != union.Find(edge) {
				results = append(results, []*graph.Node{node, edge})
				union.Union(node, edge)
			}
		}
	}

	return results
}

func main() {
	g := initiateGraph()

	results := KruskalMST(g)

	fmt.Println("The graph is:")
	fmt.Println(`
	A ------- B
	|         |
	|         |
	|         |
	C ------- D
	|
	|
	|
	E
	`)

	fmt.Println("The MST is:")

	for _, vertex := range results {
		fmt.Printf("%s ------ %s\r\n", vertex[0].Label(), vertex[1].Label())
	}
}
