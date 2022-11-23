package binarytree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	tree := New()

	for i := 1; i <= 10; i++ {
		tree.Insert(i)
	}

	if tree.Len() != 10 {
		t.Fatalf("expected tree length is 10, but it's: %d", tree.Len())
	}

	for i := 1; i <= 10; i++ {
		if tree.Find(i) == nil {
			t.Fatalf("node not found: %d", i)
		}
	}

	tree.Remove(10)

	if tree.Len() != 9 {
		t.Fatalf("expected tree length is 9, but it's: %d", tree.Len())
	}

	tree.TraverseInOrder(func(n *Node) {
		fmt.Println(n.value)
	})
}
