package queue

import "testing"

func Test_ArrayBasedQueue(t *testing.T) {
	q := NewArrayBasedQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if len := q.Len(); len != 3 {
		t.Fatal("Length of queue expected to be 3, but its", len)
	}

	if v := q.Peek(); v != 1 {
		t.Fatal("value 1 expected after peek, but we got", v)
	}

	if v := q.Dequeue(); v != 1 {
		t.Fatal("value 1 expected after dequeue, but we got", v)
	}

	if len := q.Len(); len != 2 {
		t.Fatal("Length of queue after dequeue expected to be 2, but its", len)
	}
}
