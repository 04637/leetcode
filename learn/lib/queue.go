package lib

import "sync"

type Queue struct {
	items []interface{}
	lock  sync.RWMutex
}

func (q *Queue) Add(item interface{}) *interface{} {
	q.lock.Lock()
	q.items = append(q.items, item)
	q.lock.Unlock()
	return &item
}

func (q *Queue) Poll() interface{} {
	if q.IsEmpty() {
		return nil
	}
	result := q.items[0]
	q.items = q.items[1:]
	return result
}

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.items[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) checkEmpty() {
	if len(q.items) == 0 {
		panic("queue is empty!")
	}
}

func (q *Queue) Size() int {
	return len(q.items)
}
