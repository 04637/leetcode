package q1

import "strconv"

// getMin O(1) 栈
type Stack struct {
	data []int
	min  []int
}

func (s *Stack) Pop() int {
	s.checkEmpty()
	result := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	if result == s.min[len(s.min)-1] {
		s.min = s.min[:len(s.min)-1]
	}
	return result
}

func (s *Stack) Push(n int) {
	if len(s.data) == 0 {
		s.min = append(s.min, n)
	} else if n <= s.min[len(s.min)-1] {
		s.min = append(s.min, n)
	}
	s.data = append(s.data, n)
}

func (s Stack) Peek() int {
	s.checkEmpty()
	return s.data[len(s.data)-1]
}

func (s Stack) GetMin() int {
	s.checkEmpty()
	result := s.min[len(s.min)-1]
	return result
}

func (s Stack) Empty() bool {
	return len(s.data) == 0
}

func (s Stack) checkEmpty() {
	if len(s.data) == 0 {
		panic("stack is empty")
	}
}

func (s Stack) String() string {
	var str string
	for i := len(s.data) - 1; i >= 0; i-- {
		str += strconv.Itoa(s.data[i]) + ", "
	}
	return str
}
