package stack

import "testing"

func Test_ArrayStackPush(t *testing.T) {
	stack := NewArrayStack()

	stack.Push(10)

	if stack.Len() != 1 {
		t.Fatal("Stack length expected to be 1 but it's", stack.Len())
	}
}

func Test_ArrayStackPop(t *testing.T) {
	stack := NewArrayStack()

	stack.Push(10)

	if stack.Pop() != 10 {
		t.Fatal("Stack popped value expected to be 10 but it's", stack.Pop())
	}
}
