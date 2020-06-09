package impl

func getAndRemoveLastEle(stack *Stack) int {
	v, ok := stack.Pop()
	if !ok {
		return v
	} else {
		last := getAndRemoveLastEle(stack)
		stack.Push(v)
		return last
	}
}

func Reverse(stack *Stack) {
	if stack.Empty() {
		return
	}
	v := getAndRemoveLastEle(stack)
	Reverse(stack)
	stack.Push(v)
}
