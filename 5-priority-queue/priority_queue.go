package pq

type PriorityQueue struct {
	items []int
}

func (q *PriorityQueue) heapifyUp(idx int) {
	for idx != 0 {
		parentIdx := q.parentIndex(idx)
		parent := q.parent(idx)
		value := q.value(idx)

		if parent <= value {
			break
		}

		q.swap(idx, parentIdx)

		idx = parentIdx
	}
}

func (q *PriorityQueue) heapifyDown(idx int) {
	for q.hasLeftNode(idx) {
		leastNodeIndex := q.leftNodeIndex(idx)

		if q.hasRightNode(idx) && q.rightNode(idx) < q.leftNode(idx) {
			leastNodeIndex = q.rightNodeIndex(idx)
		}

		if q.value(idx) < q.value(leastNodeIndex) {
			break
		}

		q.swap(idx, leastNodeIndex)

		idx = leastNodeIndex
	}
}

func (q *PriorityQueue) value(idx int) int {
	return q.items[idx]
}

func (q *PriorityQueue) hasLeftNode(idx int) bool {
	return q.leftNodeIndex(idx) < q.Len()
}

func (q *PriorityQueue) hasRightNode(idx int) bool {
	return q.rightNodeIndex(idx) < q.Len()
}

func (q *PriorityQueue) parent(idx int) int {
	return q.value(q.parentIndex(idx))
}

func (q *PriorityQueue) leftNode(idx int) int {
	return q.value(q.leftNodeIndex(idx))
}

func (q *PriorityQueue) rightNode(idx int) int {
	return q.value(q.rightNodeIndex(idx))
}

func (q *PriorityQueue) parentIndex(idx int) int {
	return (idx - 1) / 2
}

func (q *PriorityQueue) leftNodeIndex(idx int) int {
	return (idx * 2) + 1
}

func (q *PriorityQueue) rightNodeIndex(idx int) int {
	return (idx * 2) + 2
}

func (q *PriorityQueue) swap(idx1, idx2 int) {
	q.items[idx1], q.items[idx2] = q.items[idx2], q.items[idx1]
}

// Insert insert into queue
func (q *PriorityQueue) Insert(value int) {
	if q.IsEmpty() {
		q.items = []int{value}

		return
	}

	q.items = append(q.items, value)

	q.heapifyUp(q.Len() - 1)
}

// Peek get first node value
func (q *PriorityQueue) Peek() int {
	if q.Len() < 1 {
		panic("queue is empty")
	}

	return q.items[0]
}

// Poll remove first node
func (q *PriorityQueue) Poll() int {
	v := q.Peek()

	// swap head with last item
	q.swap(0, q.Len()-1)

	// remove last item that is now holding old head value
	q.items = q.items[0 : q.Len()-1]

	q.heapifyDown(0)

	return v
}

// Len
func (q *PriorityQueue) Len() int {
	return len(q.items)
}

// IsEmpty
func (q *PriorityQueue) IsEmpty() bool {
	return q.Len() < 1
}

func New() *PriorityQueue {
	return &PriorityQueue{}
}
