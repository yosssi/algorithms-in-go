package fifoqueues

type Queue []interface{}

func (q *Queue) Enqueue(item interface{}) {
	*q = append(*q, item)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	item := []interface{}(*q)[0]
	*q = []interface{}(*q)[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return len(*q)
}

func NewQueue() *Queue {
	return &Queue{}
}
