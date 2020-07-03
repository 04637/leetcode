package q1

import (
	"container/list"
	"fmt"
)

// 打印两个有序链表的公共部分
func PrintCommonList(list1 list.List, list2 list.List) {

	p1 := list1.Front()
	p2 := list2.Front()

	for p1 != nil && p2 != nil {
		if p1.Value.(int) < p2.Value.(int) {
			p2 = p2.Next()
		} else if p1.Value.(int) > p2.Value.(int) {
			p1 = p1.Next()
		} else {
			fmt.Println(p1.Value)
			p1 = p1.Next()
			p2 = p2.Next()
		}
	}
}
