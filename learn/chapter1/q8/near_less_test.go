package q8

import (
	"testing"
)

func TestGetNearLessNoRepeat(t *testing.T) {
	arr := []int{3, 4, 1, 5, 6, 2, 7}
	res := [][]int{
		{-1, 2},
		{0, 2},
		{-1, -1},
		{2, 5},
		{3, 5},
		{2, -1},
		{5, -1},
	}
	realRes := GetNearLessNoRepeat(arr)
	for i, v := range res {
		for j, v := range v {
			if v != realRes[i][j] {
				t.Errorf("Wrong\n%d --VS-- %d\n", res, realRes)
				return
			}
		}
	}
}

func TestGetNearLess(t *testing.T) {
	arr := []int{3, 1, 3, 4, 3, 5, 3, 2, 2}
	res := [][]int{
		{-1, 1},
		{-1, -1},
		{1, 7},
		{2, 4},
		{1, 7},
		{4, 6},
		{1, 7},
		{1, -1},
		{1, -1},
	}
	realRes := GetNearLess(arr)
	for i, v := range res {
		for j, v := range v {
			if v != realRes[i][j] {
				t.Errorf("Wrong\n%d --VS-- %d\n", res, realRes)
				return
			}
		}
	}
}
