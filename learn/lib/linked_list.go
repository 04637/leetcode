package lib

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{
		Value: value,
		Next:  nil,
	}
}

type LinkedList struct {
	Front *Node
	Back  *Node
	Size  int
}

func (l *LinkedList) Append(node *Node) {
	if l.Front == nil {
		l.Front = node
		l.Back = node
		l.Size++
		return
	}

	p := l.Front
	for ; p.Next != nil; p = p.Next {
	}
	p.Next = node
	l.Back = node
	l.Size++
}

func (l *LinkedList) Print() {
	for p := l.Front; p != nil; p = p.Next {
		fmt.Println(p.Value)
	}
}
