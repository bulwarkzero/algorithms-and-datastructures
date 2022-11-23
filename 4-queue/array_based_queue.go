package queue

type arrayQueue struct {
	arr []interface{}
}

func (q *arrayQueue) Enqueue(value interface{}) {
	q.arr = append(q.arr, value)
}

func (q *arrayQueue) Dequeue() interface{} {
	value := q.Peek()

	q.arr = q.arr[1:]

	return value
}

func (q *arrayQueue) Peek() interface{} {
	if len(q.arr) < 1 {
		panic("peek from empty queue")
	}

	return q.arr[0]
}

func (q *arrayQueue) Len() int {
	return len(q.arr)
}

func NewArrayBasedQueue() Queue {
	return &arrayQueue{arr: make([]interface{}, 0)}
}
