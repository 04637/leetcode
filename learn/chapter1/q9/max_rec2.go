package q9

import (
	"leetcode/learn/lib"
	"math"
)

func MaxRecSize2(arr [][]int) (maxArea int) {
	if arr == nil || len(arr) == 0 || len(arr[0]) == 0 {
		return
	}

	height := make([]int, len(arr[0]))
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == 0 {
				height[j] = 0
			} else {
				height[j] += arr[i][j]
			}
		}
		nearLess := GetNearLess(height)
		for i2, less := range nearLess {
			left := less[0] + 1
			right := less[1] - 1
			_size := ((i2 - left) + (right - i2) + 1) * height[i2]
			_size = int(math.Abs(float64(_size)))
			if _size > maxArea {
				maxArea = _size
			}
		}
	}
	return
}

func GetNearLess(arr []int) (res [][]int) {

	res = make([][]int, len(arr))
	for i := range res {
		res[i] = make([]int, 2)
	}

	stack := lib.NewStack()
	for i := 0; i < len(arr); i++ {
		for !stack.IsEmpty() && arr[stack.Peek().([]int)[0]] > arr[i] {
			popIndexs := stack.Pop().([]int)
			var leftIndex int
			if stack.IsEmpty() {
				leftIndex = -1
			} else {
				leftIndex = stack.Peek().([]int)[len(stack.Peek().([]int))-1]
			}
			for _, v := range popIndexs {
				res[v][0] = leftIndex
				res[v][1] = i
			}
		}
		if !stack.IsEmpty() && arr[stack.Peek().([]int)[0]] == arr[i] {
			oldArr := stack.Pop().([]int)
			stack.Push(append(oldArr, i))
		} else {
			stack.Push([]int{i})
		}
	}

	for !stack.IsEmpty() {
		popIndexs := stack.Pop().([]int)
		var leftIndex int
		if stack.IsEmpty() {
			leftIndex = -1
		} else {
			leftIndex = stack.Peek().([]int)[len(stack.Peek().([]int))-1]
		}
		for _, index := range popIndexs {
			res[index][0] = leftIndex
			res[index][1] = -1
		}
	}

	return
}
