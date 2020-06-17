package q7

import (
	"container/list"
)

func GetMaxWindow2(arr []int, w int) (res []int) {
	if arr == nil || len(arr) < w {
		return nil
	}
	maxList := list.New()

	for i := 0; i < len(arr); i++ {
		for !(maxList.Len() == 0) && arr[maxList.Back().Value.(int)] < arr[i] {
			maxList.Remove(maxList.Back())
		}
		maxList.PushBack(i)
		if maxList.Front().Value.(int) == i-w {
			maxList.Remove(maxList.Front())
		}
		if i >= w-1 {
			res = append(res, arr[maxList.Front().Value.(int)])
		}
	}
	return res
}
