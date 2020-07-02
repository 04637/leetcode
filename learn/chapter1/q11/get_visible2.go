package q11

import (
	"fmt"
	"leetcode/learn/lib"
)

// 可见山峰对的数量
func GetVisibleNum2(arr []int) (result int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	size := len(arr)
	maxIndex := 0
	// 先在环中找到其中一个最大值的位置, 哪一个都行
	for i := range arr {
		if arr[maxIndex] < arr[i] {
			maxIndex = i
		}
	}

	stack := lib.NewStack()
	// 先把(最大值,1)这个记录放入stack中
	stack.Push(newRecord(arr[maxIndex]))
	stack.Print()
	// 从最大值位置的下一个位置开始沿next方向遍历
	index := nextIndex(maxIndex, size)
	// 用小找大的方式统计所有可见山峰对
	// 遍历阶段开始, 当index再次回到maxIndex时, 说明转了一圈, 遍历结束
	for index != maxIndex {
		// 当前数字 arr[index]要进栈, 判断会不会破坏第一维的数字从顶到底依次变大
		// 如果破坏了, 就依次弹出栈顶记录, 并计算山峰对的数量
		for stack.Peek().(*Record).value < arr[index] {
			k := stack.Pop().(*Record).times
			// 弹出记录为 (X,K) 如果 K==1, 产生2对
			// 如果 K>1, 产生 2*K+C(2,K)对
			result += getInternalSum(k) + 2*k
		}
		// 当前数字arr[index]要进入栈了, 如果和当前栈顶数字一样就合并
		// 不一样就把记录(arr[index], 1)放入栈中
		if stack.Peek().(*Record).value == arr[index] {
			stack.Peek().(*Record).times++
		} else {
			stack.Push(newRecord(arr[index]))
		}
		index = nextIndex(index, size)
		stack.Print()
	}
	fmt.Println("清算前: ", result)
	// 清算阶段
	// 清算阶段的第 1 小阶段
	for stack.Size() > 2 {
		times := stack.Pop().(*Record).times
		result += getInternalSum(times) + 2*times
	}
	stack.Print()
	fmt.Print("清算第二小阶段", result)
	if stack.Size() == 2 {
		times := stack.Pop().(*Record).times
		result += getInternalSum(times)
		if stack.Peek().(*Record).times == 1 {
			result += times
		} else {
			result += times * 2
		}
	}
	stack.Print()
	fmt.Print("清算第三小阶段", result)
	// 清算阶段的第 3 小阶段
	result += getInternalSum(stack.Pop().(*Record).times)
	return
}

func nextIndex(i int, size int) int {
	if i < size-1 {
		return i + 1
	} else {
		return 0
	}
}

func getInternalSum(k int) int {
	if k == 1 {
		return 0
	} else {
		return k * (k - 1) / 2
	}
}
