package q2

import (
	"leetcode/learn/lib"
)

// 在单链表或双链表中删除倒数第K个节点, 如果链表长度为 N 时间复杂度要求达到 O(N),
// 额外空间复杂度达到 O(1)

// 从单链表删除
func RemoveLastK(list *lib.LinkedList, k int) interface{} {
	if k <= 0 || k > list.Size {
		return nil
	}
	_size := list.Size
	_k := _size - k - 1
	if _size == 0 || _k >= _size {
		return nil
	}
	index := 0

	if _k == -1 {
		// 表示删除第一个元素
		tmp := list.Front
		list.Front = tmp.Next
		list.Size--
		return tmp.Value
	}

	for p := list.Front; p != nil; p = p.Next {
		tmp := p.Next
		if index == _k {
			p.Next = tmp.Next
			list.Size--
			if p.Next == nil {
				list.Back = p
			}
			return tmp.Value
		}
		index++
	}
	return nil
}

// 双链表删除
func RemoveDeLastK(list *lib.DoubleLinkedList, k int) interface{} {
	if k < 1 || k > list.Size {
		return nil
	}
	// 删除倒数第一个元素
	if k == 1 {
		temp := list.Back
		temp.Prev.Next = nil
		list.Back = temp.Prev
		list.Size--
		return list.Back
	}
	// 删除第一个元素
	if k == list.Size {
		temp := list.Front
		temp.Next.Prev = nil
		list.Front = temp.Next
		list.Size--
		return temp
	}
	// 删除中间元素
	index := 1
	for p := list.Back; p != nil; p = p.Prev {
		if index == k {
			temp := p.Prev
			temp.Next = p.Next
			temp.Next.Prev = temp
			list.Size--
			return p
		}
		index++
	}

	return nil
}
