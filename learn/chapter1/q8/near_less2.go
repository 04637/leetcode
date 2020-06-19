package q8

import "container/list"

func GetNearLessNoRepeat2(arr []int) (res [][]int) {
	decreaseStack := list.New()
	res = make([][]int, len(arr))
	for i := range res {
		res[i] = make([]int, 2)
	}

	for i, v := range arr {
		for decreaseStack.Len() != 0 && v < arr[decreaseStack.Back().Value.(int)] {
			popIndex := decreaseStack.Back().Value.(int)
			decreaseStack.Remove(decreaseStack.Back())
			if decreaseStack.Len() != 0 {
				res[popIndex][0] = decreaseStack.Back().Value.(int)
			} else {
				res[popIndex][0] = -1
			}
			res[popIndex][1] = i
		}
		decreaseStack.PushBack(i)
	}

	for decreaseStack.Len() != 0 {
		popIndex := decreaseStack.Back().Value.(int)
		decreaseStack.Remove(decreaseStack.Back())
		res[popIndex][1] = -1
		if decreaseStack.Len() != 0 {
			res[popIndex][0] = decreaseStack.Back().Value.(int)
		} else {
			res[popIndex][0] = -1
		}
	}

	return res
}
