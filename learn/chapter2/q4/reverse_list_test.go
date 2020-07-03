package q4

import (
	"fmt"
	"leetcode/learn/lib"
	"testing"
)

func TestReverseList(t *testing.T) {
	list := new(lib.LinkedList)
	list.Append(lib.NewNode(1))
	list.Append(lib.NewNode(2))
	list.Append(lib.NewNode(3))
	list.Append(lib.NewNode(4))
	list.Append(lib.NewNode(5))
	fmt.Println("=== before reverse")
	list.Print()
	ReverseList(list)
	fmt.Println("=== after reverse")
	list.Print()
}

func TestReverseDeList(t *testing.T) {
	list := new(lib.DoubleLinkedList)
	list.Append(lib.NewDNode(1))
	list.Append(lib.NewDNode(2))
	list.Append(lib.NewDNode(3))
	list.Append(lib.NewDNode(4))
	list.Append(lib.NewDNode(5))
	fmt.Println("=== before reverse")
	list.Print()
	ReverseDeList(list)
	fmt.Println("=== after reverse")
	list.Print()
	fmt.Println("=== back")
	list.PrintBack()
}
