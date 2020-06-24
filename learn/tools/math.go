package tools

func MaxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func AbsInt(n int) int {
	if n > 0 {
		return n
	} else {
		return -1 * n
	}
}
