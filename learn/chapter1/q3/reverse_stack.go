package q3

import "leetcode/learn/chapter1/q1"

// 使用递归函数和栈操作逆序一个栈

func getAndRemoveLastEle(stack *q1.Stack) int {
	v := stack.Pop()
	if stack.Empty() {
		return v
	}
	last := getAndRemoveLastEle(stack)
	stack.Push(v)
	return last
}

func Reverse(stack *q1.Stack) {
	if stack.Empty() {
		return
	}
	v := getAndRemoveLastEle(stack)
	Reverse(stack)
	// 利用递归也是栈的特性, 最先push的为原栈顶元素
	stack.Push(v)
}
