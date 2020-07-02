package tools

// 两数最大值
func MaxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 整数绝对值
func AbsInt(n int) int {
	if n > 0 {
		return n
	} else {
		return -1 * n
	}
}

// 阶乘
func Fac(n int) (result int) {
	if n == 1 {
		return 1
	} else {
		return n * Fac(n-1)
	}
}
