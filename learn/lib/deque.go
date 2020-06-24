package lib

import (
	"fmt"
	"sync"
)

// 双端队列
type Deque struct {
	items []interface{}
	lock  sync.RWMutex
}

func (q *Deque) AddLast(item interface{}) *interface{} {
	q.lock.Lock()
	q.items = append(q.items, item)
	q.lock.Unlock()
	return &item
}

func (q *Deque) AddFirst(item interface{}) *interface{} {
	q.lock.Lock()
	_tmp := []interface{}{item}
	_tmp = append(_tmp, q.items...)
	q.items = _tmp
	q.lock.Unlock()
	return &item
}

func (q *Deque) PollFirst() interface{} {
	if q.IsEmpty() {
		return nil
	}
	result := q.items[0]
	q.items = q.items[1:]
	return result
}

func (q *Deque) PollLast() interface{} {
	if q.IsEmpty() {
		return nil
	}
	result := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return result
}

func (q *Deque) PeekFirst() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.items[0]
}

func (q *Deque) PeekLast() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.items[len(q.items)-1]
}

func (q *Deque) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Deque) checkEmpty() {
	if len(q.items) == 0 {
		panic("queue is empty!")
	}
}

func (q *Deque) Size() int {
	return len(q.items)
}

func (q *Deque) Print(prefix ...string) {
	fmt.Println(prefix, q.items)
}
