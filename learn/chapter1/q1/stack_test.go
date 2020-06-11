package q1

import (
	"math/rand"
	"testing"
)

func TestStack(t *testing.T) {
	stack := new(Stack)
	for i := 0; i < 10; i++ {
		stack.Push(rand.Intn(10))
		stack.GetMin()
	}

	for i := 0; i < 10; i++ {
		stack.GetMin()
		stack.Pop()
	}
}
