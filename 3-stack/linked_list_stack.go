package stack

type node struct {
	value int
	next  *node
}

type linkedListStack struct {
	head *node
	len  int
}

func (s *linkedListStack) Push(value int) {
	newNode := &node{value, s.head}

	s.head = newNode
	s.len++
}

func (s *linkedListStack) Pop() int {
	if s.len < 1 {
		panic("popping from empty stack")
	}

	headValue := s.head.value
	s.head = s.head.next
	s.len--

	return headValue
}

func (s *linkedListStack) Len() int {
	return s.len
}

func NewLinkedListStack() Stack {
	return &linkedListStack{}
}
