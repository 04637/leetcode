package q11

import (
	"fmt"
	"leetcode/learn/lib"
	"strconv"
)

// 可见山峰对的数量
func GetVisibleNum(arr []int) (result int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	max := -1
	maxIndex := -1
	// 找到最大值
	for i, v := range arr {
		if v > max {
			max = v
			maxIndex = i
		}
	}
	stack := lib.NewStack()

	for i := maxIndex; i < len(arr)+maxIndex; i++ {
		var current int
		if i < len(arr) {
			current = arr[i]
		} else {
			current = arr[i-len(arr)]
		}
		for !stack.IsEmpty() && stack.Peek().(*Record).value < current {
			record := stack.Pop().(*Record)
			if record.times == 1 {
				result += 2
			} else {
				result += 2*record.times + getInternalSum(record.times)
			}
		}
		if !stack.IsEmpty() && stack.Peek().(*Record).value == current {
			stack.Peek().(*Record).times += 1
		} else {
			stack.Push(newRecord(current))
		}
		stack.Print()
	}
	fmt.Println("清算前: ", result)
	for !stack.IsEmpty() {
		record := stack.Pop().(*Record)
		if stack.Size() >= 2 {
			if record.times == 1 {
				result += 2
			} else {
				result += 2*record.times + getInternalSum(record.times)
			}
		} else if stack.Size() == 1 {
			fmt.Print("清算第二小阶段", result)
			result += getInternalSum(record.times)
			if stack.Peek().(*Record).times >= 2 {
				result += 2 * record.times
			} else {
				result += record.times
			}
		} else if stack.IsEmpty() {
			fmt.Print("清算第三小阶段", result)
			result += getInternalSum(record.times)
		}
		stack.Print()
	}

	return
}

type Record struct {
	value int
	times int
}

func newRecord(v int) *Record {
	return &Record{
		v,
		1,
	}
}

func (r *Record) String() string {
	return strconv.Itoa(r.value) + " - " + strconv.Itoa(r.times)
}
