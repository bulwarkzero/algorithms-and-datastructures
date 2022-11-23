package stack

import "testing"

func Test_LinkedListStackPush(t *testing.T) {
	stack := NewLinkedListStack()

	stack.Push(10)

	if stack.Len() != 1 {
		t.Fatal("Stack length expected to be 1 but it's", stack.Len())
	}
}

func Test_LinkedListStackPop(t *testing.T) {
	stack := NewLinkedListStack()

	stack.Push(10)

	if stack.Pop() != 10 {
		t.Fatal("Stack popped value expected to be 10 but it's", stack.Pop())
	}
}
