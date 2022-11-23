package stack

type arrayStack struct {
	arr []int
}

func (s *arrayStack) Push(value int) {
	s.arr = append(s.arr, value)
}

func (s *arrayStack) Pop() int {
	if len(s.arr) < 1 {
		panic("popping from empty stack")
	}

	value := s.arr[len(s.arr)-1]
	s.arr = s.arr[0 : len(s.arr)-1]

	return value
}

func (s *arrayStack) Len() int {
	return len(s.arr)
}

func NewArrayStack() Stack {
	return &arrayStack{arr: make([]int, 0)}
}
