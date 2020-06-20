package q9

import (
	"testing"
)

func TestMaxRecSize(t *testing.T) {
	arr := [][]int{
		{1, 0, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 0},
	}
	maxSize := MaxRecSize(arr)
	maxSize2 := MaxRecSize2(arr)
	if maxSize != maxSize2 {
		t.Errorf("Wrong rec size, expected %d but got %d", maxSize, maxSize2)
	}
}
