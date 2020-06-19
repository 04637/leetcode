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
	maxSize := MaxRecSize2(arr)
	if maxSize != 6 {
		t.Errorf("Wrong rec size, expected 6 but got %d", maxSize)
	}
}
