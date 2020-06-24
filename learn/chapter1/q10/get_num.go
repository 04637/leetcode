package q10

import (
	"leetcode/learn/lib"
)

// 最大值减去最小值小于或等于num的子数组数量
func GetNum(arr []int, num int) (n int) {
	if arr == nil || len(arr) == 0 || num < 0 {
		return 0
	}

	qmax := new(lib.Deque)
	qmin := new(lib.Deque)

	left := 0
	right := 0

	for left < len(arr) {
		for right < len(arr) {
			if qmin.IsEmpty() || qmin.PeekLast() != right {
				for !qmin.IsEmpty() && arr[qmin.PeekLast().(int)] >= arr[right] {
					qmin.PollLast()
				}
				qmin.AddLast(right)

				for !qmax.IsEmpty() && arr[qmax.PeekLast().(int)] <= arr[right] {
					qmax.PollLast()
				}
				qmax.AddLast(right)
			}
			if arr[qmax.PeekFirst().(int)]-arr[qmin.PeekFirst().(int)] > num {
				break
			}
			right++
		}
		n += right - left
		if qmin.PeekFirst() == left {
			qmin.PollFirst()
		}
		if qmax.PeekFirst() == left {
			qmax.PollFirst()
		}
		left++
	}

	return
}
