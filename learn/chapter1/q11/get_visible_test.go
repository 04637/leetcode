package q11

import (
	"fmt"
	"testing"
)

func TestGetVisibleNum(t *testing.T) {
	arr := []int{
		4, 2, 4, 5, 3, 4, 5, 2, 3, 5, 4,
	}
	num := GetVisibleNum(arr)
	fmt.Println("#####################")
	num2 := GetVisibleNum2(arr)
	if num2 != num {
		t.Errorf("wrong visible num, expected %d, but got %d", num2, num)
	}
}
