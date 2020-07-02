package tools

import "testing"

func TestMaxInt(t *testing.T) {
	a := 3
	b := 5
	if max := MaxInt(a, b); max != 5 {
		t.Errorf("wrong maxInt, expected 5 but got %d", max)
	}
}

func TestFac(t *testing.T) {
	if Fac(3) != 6 {
		t.Errorf("wrong fac, expected 6 but got %d", Fac(3))
	}
	if Fac(2) != 2 {
		t.Errorf("wrong fac, expected 2 but got %d", Fac(2))
	}
	if Fac(4) != 24 {
		t.Errorf("wrong fac, expected 2 but got %d", Fac(4))
	}
}
