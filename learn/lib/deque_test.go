package lib

import "testing"

var deque *Deque

func init() {
	deque = new(Deque)
}

func TestDeque_AddFirst(t *testing.T) {
	deque.AddFirst(3)
	deque.AddFirst(4)
	deque.AddFirst(5)
	if deque.PeekLast() != 3 {
		t.Errorf("Wrong peek last, expected 3 but got %d", deque.PeekLast())
	}
	if deque.PeekFirst() != 5 {
		t.Errorf("Wrong peek first, expected 5 but got %d", deque.PeekFirst())
	}
}

func TestDeque_AddLast(t *testing.T) {
	deque.AddLast(7)
	deque.AddLast(8)
	deque.AddLast(9)
	if deque.PeekLast() != 9 {
		t.Errorf("Wrong peek last, expected 9 but got %d", deque.PeekLast())
	}
	if deque.PeekFirst() != 5 {
		t.Errorf("Wrong peek first, expected 5 but got %d", deque.PeekFirst())
	}
}

func TestDeque_PollFirst(t *testing.T) {
	res := deque.PollFirst()
	if res != 5 {
		t.Errorf("Wrong poll first, expected 5 but got %d", res)
	}

	res2 := deque.PollFirst()
	if res2 != 4 {
		t.Errorf("Wrong poll first, expected 4 but got %d", res2)
	}
}

func TestDeque_PollLast(t *testing.T) {
	res := deque.PollLast()
	if res != 9 {
		t.Errorf("Wrong poll last, expected 9 but got %d", res)
	}

	res2 := deque.PollLast()
	if res2 != 8 {
		t.Errorf("Wrong poll first, expected 8 but got %d", res2)
	}
}
