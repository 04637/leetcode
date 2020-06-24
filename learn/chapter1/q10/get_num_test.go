package q10

import (
	"fmt"
	"testing"
)

func TestGetNum(t *testing.T) {
	arr := []int{3, 2, 1, 4, 5}
	n := GetNum(arr, 2)
	fmt.Println(n)
}
