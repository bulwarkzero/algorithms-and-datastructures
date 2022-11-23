package queue

type Queue interface {
	Enqueue(value interface{})
	Dequeue() interface{}
	Peek() interface{}
	Len() int
}
