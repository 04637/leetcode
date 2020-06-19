package q9

import (
	"leetcode/learn/lib"
	"leetcode/learn/tools"
)

// 求最大子矩阵大小
func MaxRecSize(arr [][]int) (maxArea int) {
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
		maxArea = MaxRecFromBottom(height)
	}
	return
}

func MaxRecFromBottom(height []int) (maxArea int) {
	if height == nil || len(height) == 0 {
		return
	}
	stack := lib.NewStack()
	for i := 0; i < len(height); i++ {
		for !stack.IsEmpty() && height[i] <= height[stack.Peek().(int)] {
			j := stack.Pop().(int)
			var k int
			if stack.IsEmpty() {
				k = -1
			} else {
				k = stack.Peek().(int)
			}
			curArea := (j - k - 1) * height[j]
			maxArea = tools.MaxInt(curArea, maxArea)
		}
		stack.Push(i)
	}
	for !stack.IsEmpty() {
		j := stack.Pop().(int)
		var k int
		if stack.IsEmpty() {
			k = -1
		} else {
			k = stack.Peek().(int)
			curArea := (len(height) - k - 1) * height[j]
			maxArea = tools.MaxInt(curArea, maxArea)
		}
	}
	return maxArea
}
