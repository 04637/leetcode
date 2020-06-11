package q2

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := new(Queue)
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
