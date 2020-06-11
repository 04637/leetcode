package q3

import (
	"fmt"
	"leetcode/learn/chapter1/q1"
	"math/rand"
	"testing"
)

func TestReverse(t *testing.T) {
	stack := new(q1.Stack)
	for i := 0; i < 10; i++ {
		stack.Push(rand.Intn(100))
	}
	fmt.Println("倒置前: ")
	fmt.Println(stack)
	Reverse(stack)
	fmt.Println("倒置后: ")
	fmt.Println(stack)
}
