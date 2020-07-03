package q2

import (
	"fmt"
	"leetcode/learn/lib"
	"testing"
)

func TestRemoveLastK(t *testing.T) {
	list := new(lib.LinkedList)
	list.Append(lib.NewNode(3))
	list.Append(lib.NewNode(2))
	list.Append(lib.NewNode(1))
	RemoveLastK(list, 3)
	list.Print()
}

func TestRemoveDeLastK(t *testing.T) {
	list := new(lib.DoubleLinkedList)
	list.Append(lib.NewDNode(3))
	list.Append(lib.NewDNode(2))
	list.Append(lib.NewDNode(1))
	RemoveDeLastK(list, 2)
	list.Print()
	fmt.Println("----------------------")
	list.PrintBack()
}
