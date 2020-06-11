package q5

import (
	"leetcode/learn/chapter1/q1"
)

// 用一个栈实现另一个栈的排序 使栈顶->栈底: 大->小
func StackSort(stack *q1.Stack) {
	help := new(q1.Stack)
	for !stack.Empty() {
		cur := stack.Pop()
		// 要使原栈从大到小排序, 辅助栈需要从小到大, 栈底为最大值
		// 当准备压入辅助栈的值比栈顶大时, 需要清空辅助栈, 将该值压入栈底
		for !help.Empty() && help.Peek() < cur {
			stack.Push(help.Pop())
		}
		help.Push(cur)
	}
	// 此时辅助栈已按从小到大排序, pop至原栈, 即为从大到小排序
	for !help.Empty() {
		stack.Push(help.Pop())
	}
}
