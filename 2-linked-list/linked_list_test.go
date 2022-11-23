package linkedlist

import "testing"

func TestAdd(t *testing.T) {
	list := New()
	for i := 1; i <= 10; i++ {
		list.Add(i)
	}

	if list.Size() != 10 {
		t.Fatalf("size expected to be 10, got %d", list.Size())
	}
}

func TestRemoveFirst(t *testing.T) {
	list := New()

	list.Add(12)
	list.Add(13)

	value := list.RemoveFirst()
	if value != 12 {
		t.Fatalf("first item expected to be 12, got %d", value)
	}

	if list.Size() != 1 {
		t.Fatalf("size expected to be 1, got %d", list.Size())
	}
}

func TestRemoveLast(t *testing.T) {
	list := New()

	list.Add(12)
	list.Add(13)

	value := list.RemoveLast()
	if value != 13 {
		t.Fatalf("last item expected to be 13, got %d", value)
	}

	if list.Size() != 1 {
		t.Fatalf("size expected to be 1, got %d", list.Size())
	}
}

func TestRemoveNode(t *testing.T) {
	list := New()

	list.Add(12)
	list.Add(13)
	list.Add(14)

	head := list.Head()

	value := list.RemoveNode(head)
	if value != 12 {
		t.Fatalf("first item expected to be 12, got %d", value)
	}

	if list.Size() != 2 {
		t.Fatalf("size expected to be 2, got %d", list.Size())
	}
}
