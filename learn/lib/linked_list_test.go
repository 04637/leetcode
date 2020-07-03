package lib

import (
	"fmt"
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	list := new(LinkedList)
	list.Append(NewNode("hello"))
	list.Append(NewNode("world"))
	list.Append(NewNode("this"))
	for p := list.Front; p != nil; p = p.Next {
		fmt.Println(p.Value)
	}
	fmt.Println(list.Size)
}

func TestDoubleLinkedList_Append(t *testing.T) {
	list := new(DoubleLinkedList)
	list.Append(NewDNode("hello"))
	list.Append(NewDNode("world"))
	list.Append(NewDNode("this"))

	for p := list.Front; p != nil; p = p.Next {
		fmt.Println(p.Value)
	}

	for p := list.Back; p != nil; p = p.Prev {
		fmt.Println(p.Value)
	}

	fmt.Println(list.Size)
}
