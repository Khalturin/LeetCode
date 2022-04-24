package ReversInteger

import "testing"

func TestReversIntegerOne(t *testing.T) {
	i := reverse(123)
	if i != 321 {
		t.Error("fail, expected 321, result: ", i)
	}
}

func TestReversIntegerTwo(t *testing.T) {
	i := reverse(-123)
	if i != -321 {
		t.Error("fail, expected -321, result: ", i)
	}
}

func TestReversIntegerThree(t *testing.T) {
	i := reverse(120)
	if i != 21 {
		t.Error("fail, expected 21, result: ", i)
	}
}

func TestReversIntegerFore(t *testing.T) {
	i := reverse(1534236469)
	if i != 0 {
		t.Error("fail, expected 0, result: ", i)
	}
}

func TestReversIntegerFive(t *testing.T) {
	i := reverse(-2147483648)
	if i != 0 {
		t.Error("fail, expected 0, result: ", i)
	}
}
func TestReversIntegerSix(t *testing.T) {
	i := reverse(1563847412)
	if i != 0 {
		t.Error("fail, expected 0, result: ", i)
	}
}
