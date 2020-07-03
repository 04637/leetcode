package q1

import (
	"container/list"
	"testing"
)

func TestPrintCommonList(t *testing.T) {
	list1 := list.New()
	list1.PushBack(8)
	list1.PushBack(7)
	list1.PushBack(5)
	list1.PushBack(4)
	list1.PushBack(1)

	list2 := list.New()
	list2.PushBack(6)
	list2.PushBack(4)
	list2.PushBack(3)
	list2.PushBack(1)
	PrintCommonList(*list1, *list2)
}
