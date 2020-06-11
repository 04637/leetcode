package q2

import "leetcode/learn/chapter1/q1"

// 由两个栈组成的队列
type Queue struct {
	stackPush q1.Stack
	stackPop  q1.Stack
}

func (q *Queue) Add(n int) {
	q.stackPush.Push(n)
	q.pushToPop()
}

func (q *Queue) Poll() int {
	q.pushToPop()
	return q.stackPop.Pop()
}

func (q *Queue) Peek() int {
	q.pushToPop()
	return q.stackPop.Peek()
}

// 队列只能查看或取最先进入队列的元素,
// 所以只用在pop栈为空时, 进行一次转换即可
// 不为空时无需关心它是否与push栈同步
func (q *Queue) pushToPop() {
	if q.stackPop.Empty() {
		for !q.stackPush.Empty() {
			v := q.stackPush.Pop()
			q.stackPop.Push(v)
		}
	}
}
