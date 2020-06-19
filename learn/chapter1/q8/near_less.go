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
			if stack.Len() == 0 {
				res[popIndex][0] = -1
			} else {
				res[popIndex][0] = stack.Back().Value.(int)
			}
			res[popIndex][1] = i
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

func GetNearLess(arr []int) (res [][]int) {
	res = make([][]int, len(arr))
	for i := range res {
		res[i] = make([]int, 2)
	}

	stack := list.New()

	for i, v := range arr {
		for stack.Len() != 0 && arr[stack.Back().Value.([]int)[0]] > v {
			popIs := stack.Back().Value.([]int)
			stack.Remove(stack.Back())
			var leftLessIndex int
			if stack.Len() == 0 {
				leftLessIndex = -1
			} else {
				leftLessIndex = stack.Back().Value.([]int)[len(stack.Back().Value.([]int))-1]
			}
			for _, v := range popIs {
				res[v][0] = leftLessIndex
				res[v][1] = i
			}
		}
		if stack.Len() != 0 && arr[stack.Back().Value.([]int)[0]] == v {
			stack.Back().Value = append(stack.Back().Value.([]int), i)
		} else {
			stackList := []int{i}
			stack.PushBack(stackList)
		}
	}

	for stack.Len() != 0 {
		popIs := stack.Back().Value.([]int)
		stack.Remove(stack.Back())
		var leftLessIndex int
		if stack.Len() == 0 {
			leftLessIndex = -1
		} else {
			leftLessIndex = stack.Back().Value.([]int)[len(stack.Back().Value.([]int))-1]
		}
		for _, v := range popIs {
			res[v][0] = leftLessIndex
			res[v][1] = -1
		}
	}
	return
}
