package q7

import (
	"leetcode/learn/lib"
	"math"
)

func WindowMove(arr []int, w int) []int {
	// 所有最大值数组
	var res []int
	// 临时存储当前窗口数
	var maxW []int
	for i := 0; i < len(arr); i++ {
		maxW = append(maxW, arr[i])
		if v, _arr, ok := maxAndPoll(maxW, w); ok {
			maxW = _arr
			res = append(res, v)
		}
	}
	return res
}

// 获取数组最大值并且将第一个元素去除
func maxAndPoll(arr []int, w int) (int, []int, bool) {
	if len(arr) != w {
		return 0, arr, false
	}

	max := math.MinInt64
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	arr = arr[1:]
	return max, arr, true
}

// 生成窗口最大值数组 O(N)算法
func GetMaxWindow(arr []int, w int) []int {
	if arr == nil || w < 1 || len(arr) < w {
		return nil
	}
	qmax := new(lib.Deque)
	res := make([]int, len(arr)-w+1)
	index := 0
	for i := 0; i < len(arr); i++ {
		for !qmax.IsEmpty() && arr[qmax.PeekLast().(int)] <= arr[i] {
			qmax.PollLast()
		}
		qmax.AddLast(i)
		if qmax.PeekFirst() == i-w {
			qmax.PollFirst()
		}
		if i >= w-1 {
			res[index] = arr[qmax.PeekFirst().(int)]
			index++
		}
	}
	return res
}
