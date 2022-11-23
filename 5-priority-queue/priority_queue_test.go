package pq

import (
	"testing"
)

func Test_PriorityQueue(t *testing.T) {
	q := New()

	q.Insert(1)
	q.Insert(10)
	q.Insert(0)
	q.Insert(2)

	if len := q.Len(); len != 4 {
		t.Fatalf("queue length expected to be 4, but it's %d", len)
	}

	if leastValue := q.Poll(); leastValue != 0 {
		t.Fatalf("least value expected to be 0, but it's %d", leastValue)
	}

	if len := q.Len(); len != 3 {
		t.Fatalf("queue length after poll expected to be 3, but it's %d", len)
	}

	if leastValue := q.Poll(); leastValue != 1 {
		t.Fatalf("least value after two poll expected to be 1, but it's %d", leastValue)
	}
}
