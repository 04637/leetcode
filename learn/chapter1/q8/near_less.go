package q8

import "container/list"

// 单调栈结构
func GetNearLessNoRepeat(arr []int) (res [][]int) {
	stack := list.New()
	res = make([][]int, len(arr))
	for i := range res {
		res[i] = make([]int, 2)
	}
	for i := 0; i < len(arr); i++ {
		for stack.Len() != 0 && arr[stack.Back().Value.(int)] > arr[i] {
			popIndex := stack.Back().Value.(int)
			stack.Remove(stack.Back())
			var leftLessIndex int
			if stack.Len() == 0 {
				leftLessIndex = -1
			} else {
				leftLessIndex = stack.Back().Value.(int)
				res[popIndex][0] = leftLessIndex
				res[popIndex][1] = i
			}
		}
		stack.PushBack(i)
	}
	for stack.Len() != 0 {
		popIndex := stack.Back().Value.(int)
		stack.Remove(stack.Back())
		var leftLessIndex int
		if stack.Len() == 0 {
			leftLessIndex = -1
		} else {
			leftLessIndex = stack.Back().Value.(int)
		}

		res[popIndex][0] = leftLessIndex
		res[popIndex][1] = -1
	}
	return res
}
