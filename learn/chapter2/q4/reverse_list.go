package q4

import "leetcode/learn/lib"

func ReverseList(list *lib.LinkedList) {
	list.Back = list.Front
	var curPrev *lib.Node
	curNext := list.Front.Next
	for cur := list.Front; cur != nil; cur = curNext {
		curNext = cur.Next
		cur.Next = curPrev
		curPrev = cur
	}
	list.Front = curPrev
}

func ReverseDeList(list *lib.DoubleLinkedList) {
	list.Back = list.Front
	var curPrev *lib.DNode
	curNext := list.Front.Next
	for cur := list.Front; cur != nil; cur = curNext {
		curNext = cur.Next
		cur.Next = curPrev
		cur.Prev = curNext
		curPrev = cur
	}
	list.Front = curPrev
}
