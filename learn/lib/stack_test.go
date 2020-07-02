package lib

import (
	"testing"
)

func TestStack_Print(t *testing.T) {
	stack := NewStack()
	stack.Push("hello")
	stack.Push("world")
	stack.Push("this")
	stack.Push("current")
	stack.Print()
}
