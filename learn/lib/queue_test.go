package lib

import "testing"

var queue *Queue

func init() {
	queue = new(Queue)
}

func TestAdd(t *testing.T) {
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	if size := queue.Size(); size != 3 {
		t.Errorf("Wrong count, expected 3 and got %d", size)
	}
}

func TestPoll(t *testing.T) {
	queue.Poll()
	if size := len(queue.items); size != 2 {
		t.Errorf("Wrong count, expected 2 and got %d", size)
	}

	v1 := queue.Poll().(int)
	if v1 != 2 {
		t.Errorf("Wrong poll value, expected 2 and got %d", v1)
	}
	v2 := queue.Poll().(int)
	if v2 != 3 {
		t.Errorf("Wrong poll value, expected 3 and got %d", v2)
	}
	if size := len(queue.items); size != 0 {
		t.Errorf("Wrong count, expected 0 and got %d", size)
	}

	if !queue.IsEmpty() {
		t.Errorf("IsEmpty should return true")
	}
}
