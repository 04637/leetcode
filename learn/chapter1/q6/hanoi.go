package q6

import "fmt"

func hanoiP1(num int, left string, mid string, right string) int {
	if num < 1 {
		return 0
	}
	return process(num, left, mid, right, left, right)
}

func process(num int, left string, mid string, right string, from string, to string) int {
	if num == 1 {
		if from == mid || to == mid {
			fmt.Println("Move 1 from", from, "to", to)
			return 1
		} else {
			fmt.Println("Move 1 from", from, "to", mid)
			fmt.Println("Move 1 from", mid, "to", to)
			return 2
		}
	}
	if from == mid || to == mid {
		var another string
		if from == left || to == left {
			another = right
		} else {
			another = left
		}

		part1 := process(num-1, left, mid, right, from, another)
		part2 := 1

		fmt.Println("Move", num, "from", from, "to", to)
		part3 := process(num-1, left, mid, right, another, to)
		return part1 + part2 + part3
	} else {
		part1 := process(num-1, left, mid, right, from, to)
		part2 := 1
		fmt.Println("Move", num, "from", from, "to", mid)
		part3 := process(num-1, left, mid, right, to, from)
		part4 := 1
		fmt.Println("Move", num, "from", mid, "to", to)
		part5 := process(num-1, left, mid, right, from, to)
		return part1 + part2 + part3 + part4 + part5
	}
}
