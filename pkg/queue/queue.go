package queue

import (
	"order_export_go/pkg/providers"
)

type Queue struct {
	data []*providers.Order
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]*providers.Order, 0),
	}
}

func (q *Queue) Enqueue(item *providers.Order) {
	q.data = append(q.data, item)
}

func (q *Queue) Dequeue() *providers.Order {
	if q.isEmpty() {
		return nil
	}

	item := q.data[0]
	q.data = q.data[1:]

	return item
}

func (q *Queue) Peek() *providers.Order {
	if q.isEmpty() {
		return nil
	}

	return q.data[0]
}

func (q *Queue) isEmpty() bool {
	return len(q.data) == 0
}
