package impl

import "strconv"

// getMin O(1) 栈
type Stack struct {
	data []int
	min  []int
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	result := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	if result == s.min[len(s.min)-1] {
		s.min = s.min[:len(s.min)-1]
	}
	return result, true
}

func (s *Stack) Push(n int) {
	if len(s.data) == 0 {
		s.min = append(s.min, n)
	} else if n <= s.min[len(s.min)-1] {
		s.min = append(s.min, n)
	}
	s.data = append(s.data, n)
}

func (s Stack) Peek() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	} else {
		return s.data[len(s.data)-1], true
	}
}

func (s Stack) GetMin() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	result := s.min[len(s.min)-1]
	return result, true
}

func (s Stack) Empty() bool {
	return len(s.data) == 0
}

func (s Stack) String() string {
	var str string
	for i := len(s.data) - 1; i >= 0; i-- {
		str += strconv.Itoa(s.data[i])+", "
	}
	return str
}
