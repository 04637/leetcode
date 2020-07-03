package lib

import "fmt"

type DNode struct {
	Value interface{}
	Next  *DNode
	Prev  *DNode
}

func NewDNode(value interface{}) *DNode {
	return &DNode{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
}

type DoubleLinkedList struct {
	Front *DNode
	Back  *DNode
	Size  int
}

func (l *DoubleLinkedList) Append(node *DNode) {
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
	node.Prev = p
	l.Back = node
	l.Size++
}

func (l *DoubleLinkedList) Print() {
	for p := l.Front; p != nil; p = p.Next {
		fmt.Println(p.Value)
	}
}

func (l *DoubleLinkedList) PrintBack() {
	for p := l.Back; p != nil; p = p.Prev {
		fmt.Println(p.Value)
	}
}
