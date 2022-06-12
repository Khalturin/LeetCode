package StringToInt

import (
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	// 1
	val := myAtoi("42")
	if val != 42 {
		t.Error("fail 1, expected 42, result = ", val)
	}

	// 2
	val = myAtoi("   -42")
	if val != -42 {
		t.Error("fail 2, expected -42, result = ", val)
	}

	// 3
	val = myAtoi("4193 with words")
	if val != 4193 {
		t.Error("fail 3, expected 4193, result = ", val)
	}

	// 4
	val = myAtoi(" with words 4193")
	if val != 0 {
		t.Error("fail 4, expected 0, result = ", val)
	}

	// 5
	val = myAtoi("-91283472332")
	if val != -2147483648 {
		t.Error("fail 4, expected -2147483648, result = ", val)
	}

	// 5
	val = myAtoi("91283472332")
	if val != 2147483647 {
		t.Error("fail 5, expected 2147483647, result = ", val)
	}

	// 6
	val = myAtoi("-000000000000000000000000000000000000000000000000001")
	if val != -1 {
		t.Error("fail 6, expected -1, result = ", val)
	}

	// 6
	val = myAtoi("+000000000000000000000000000000000000000000000000001")
	if val != 1 {
		t.Error("fail 6, expected 1, result = ", val)
	}

	// 7
	val = myAtoi("21474836460")
	if val != 2147483647 {
		t.Error("fail 7, expected 2147483647, result = ", val)
	}

	// 8
	val = myAtoi("9223372036854775808")
	if val != 2147483647 {
		t.Error("fail 8, expected 2147483647, result = ", val)
	}

	// 9
	val = myAtoi("2147483646")
	if val != 2147483646 {
		t.Error("fail 9, expected 2147483646, result = ", val)
	}

	// 10
	val = myAtoi("-2147483647")
	if val != -2147483647 {
		t.Error("fail 10, expected 2147483647, result = ", val)
	}

	// 11
	val = myAtoi("+100000000000000000000000000000000000000000000000000000000000000000001")
	if val != math.MaxInt32 {
		t.Error("fail 11, expected 2147483647, result = ", val)
	}

	// 12
	val = myAtoi("-100000000000000000000000000000000000000000000000000000000000000000001")
	if val != math.MinInt32 {
		t.Error("fail 12, expected -2147483647, result = ", val)
	}
}
