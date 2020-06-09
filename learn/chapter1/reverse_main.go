package main

import (
	"fmt"
	"leetcode/learn/chapter1/impl"
	"math/rand"
)

func main() {
	stack := new(impl.Stack)
	for i:=0;i<10;i++ {
		stack.Push(rand.Intn(100))
	}
	fmt.Println("倒置前: ")
	fmt.Println(stack)
	impl.Reverse(stack)
	fmt.Println("倒置后: ")
	fmt.Println(stack)

}
