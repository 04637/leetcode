package lib

import "container/list"

type Stack struct {
	data *list.List
}

func NewStack() *Stack {
	return new(Stack).Init()
}

func (s *Stack) Init() *Stack {
	s.data = list.New()
	return s
}

func (s *Stack) Push(ele interface{}) interface{} {
	s.data.PushBack(ele)
	return ele
}

func (s *Stack) Pop() interface{} {
	if s.data.Len() == 0 {
		return nil
	}
	back := s.data.Back()
	s.data.Remove(back)
	return back.Value
}

func (s *Stack) Peek() interface{} {
	return s.data.Back().Value
}

func (s *Stack) IsEmpty() bool {
	return s.data.Len() == 0
}
