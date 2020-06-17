package q7

import (
	"testing"
)

func TestWindowMove(t *testing.T) {
	arr := []int{4, 3, 5, 4, 3, 3, 6, 7}
	res1 := WindowMove(arr, 4)
	res2 := GetMaxWindow2(arr, 4)
	for i := 0; i < len(res1); i++ {
		if res2[i] != res1[i] {
			t.Errorf("Wrong window move \nres1: %d -VS- res2: %d", res1, res2)
			break
		}
	}
}
