package q5

import (
	"fmt"
	"leetcode/learn/chapter1/q1"
	"testing"
)

func TestStackSort(t *testing.T) {
	stack := new(q1.Stack)
	//for i := 0; i < 10; i++ {
	//	stack.Push(rand.Intn(100))
	//}
	stack.Push(2)
	stack.Push(4)
	stack.Push(1)
	stack.Push(3)

	fmt.Println("排序前: ")
	fmt.Println(stack)
	fmt.Println("排序后: ")
	StackSort(stack)
	fmt.Println(stack)
}
