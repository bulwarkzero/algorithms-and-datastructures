package union

import (
	graph "new-way/4-1-1-graph"
	"testing"
)

func Test_Union(t *testing.T) {
	// g := graph.NewGraph()
	n0 := graph.NewNode(0, "A") // A
	n1 := graph.NewNode(1, "B") // B
	n2 := graph.NewNode(2, "C") // C
	n3 := graph.NewNode(3, "D") // D
	n4 := graph.NewNode(3, "E") // E

	u := New([]*graph.Node{n0, n1, n2, n3, n4})

	u.Union(n0, n1) // union A-B
	u.Union(n1, n3) // union B-D
	u.Union(n2, n4) // union C-E

	if u.Find(n0) != u.Find(n1) {
		t.Fatal("A and B must be in same component")
	}

	if u.Find(n0) != u.Find(n3) {
		t.Fatal("A and D must be in same component")
	}

	if u.Find(n2) != u.Find(n4) {
		t.Fatal("C and E must be in same component")
	}

	if u.Find(n0) == u.Find(n2) {
		t.Fatal("A and C must not be in same component")
	}
}
