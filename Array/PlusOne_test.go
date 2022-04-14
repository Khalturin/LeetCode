package LeetCode

import "testing"

func TestPlusOne1(t *testing.T) {
	num := []int{1, 2, 3}
	res := plusOne(num)
	if res[0] != 1 || res[1] != 2 || res[2] != 4 {
		t.Error("TestPlusOne1 fail, expected 1,2,4 result: ", res)
	}
}

func TestPlusOne2(t *testing.T) {
	num := []int{4, 3, 2, 1}
	res := plusOne(num)
	if res[0] != 4 || res[1] != 3 || res[2] != 2 || res[3] != 2 {
		t.Error("TestPlusOne2 fail, expected 4,4,2,2 result: ", res)
	}
}

func TestPlusOne3(t *testing.T) {
	num := []int{9}
	res := plusOne(num)
	if res[0] != 1 || res[1] != 0 {
		t.Error("TestPlusOne3 fail, expected 1,0 result: ", res)
	}
}

func TestPlusOne4(t *testing.T) {
	num := []int{9, 9}
	res := plusOne(num)
	if res[0] != 1 || res[1] != 0 || res[2] != 0 {
		t.Error("TestPlusOne3 fail, expected 1,0,0 result: ", res)
	}
}

func TestPlusOne5(t *testing.T) {
	num := []int{9, 9, 9}
	res := plusOne(num)
	if res[0] != 1 || res[1] != 0 || res[2] != 0 || res[3] != 0 {
		t.Error("TestPlusOne3 fail, expected 1,0,0,0 result: ", res)
	}
}

func TestPlusOne6(t *testing.T) {
	num := []int{9, 8, 7}
	res := plusOne(num)
	if res[0] != 9 || res[1] != 8 || res[2] != 8 {
		t.Error("TestPlusOne3 fail, expected 9,8,8 result: ", res)
	}
}

func TestPlusOne7(t *testing.T) {
	num := []int{7, 8, 9}
	res := plusOne(num)
	if res[0] != 7 || res[1] != 9 || res[2] != 0 {
		t.Error("TestPlusOne3 fail, expected 7,9,0 result: ", res)
	}
}
