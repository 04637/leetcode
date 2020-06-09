package main

import (
	main2 "leetcode/learn/chapter1/impl"
	"math/rand"
)

func main() {
	stack := new(main2.Stack)
	for i := 0; i < 10; i++ {
		stack.Push(rand.Intn(10))
		stack.GetMin()
	}

	for i := 0; i < 10; i++ {
		stack.GetMin()
		stack.Pop()
	}
}
