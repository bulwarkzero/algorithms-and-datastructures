package avltree

import (
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
	tree.Remove(3)
	tree.Remove(1)

	if tree.Len() != 7 {
		t.Fatalf("expected tree length is 7, but it's: %d", tree.Len())
	}
}

func TestTreeReverseOrder(t *testing.T) {
	tree := New()

	for i := 10; i >= 1; i-- {
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
}

func TestTreeDuplicate(t *testing.T) {
	tree := New()

	tree.Insert(10)

	if tree.Insert(10) {
		t.Fatal("inserting duplicate values must return false, got true")
	}

	if tree.Len() != 1 {
		t.Fatal("expected to tree length didn't change after duplicate inserts")
	}
}
