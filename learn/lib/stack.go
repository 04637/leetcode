package lib

import (
	"container/list"
	"fmt"
)

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

func (s *Stack) Size() int {
	return s.data.Len()
}

func (s Stack) Print() {
	index := 0
	fmt.Println("-----------------------------")
	for e := s.data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value, "\t: ", index)
		index++
	}
	fmt.Println()
}
