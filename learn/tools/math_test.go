package tools

import "testing"

func TestMaxInt(t *testing.T) {
	a := 3
	b := 5
	if max := MaxInt(a, b); max != 5 {
		t.Errorf("wrong maxInt, expected 5 but got %d", max)
	}
}
