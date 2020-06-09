package main

import (
	"fmt"
	"leetcode/learn/chapter1/impl"
	"math/rand"
)

func main() {
	queue := new(impl.Queue)
	for i := 0; i < 10; i++ {
		toAdd := rand.Intn(100)
		queue.Add(toAdd)
		fmt.Printf("add: %d\n", toAdd)
		if peek, ok := queue.Peek(); ok {
			fmt.Printf("peek: %d\n", peek)
		}
		if poll, ok := queue.Poll(); ok {
			fmt.Printf("poll: %d\n", poll)
		}
	}
}
