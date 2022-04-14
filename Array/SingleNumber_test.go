package LeetCode

import "testing"

func TestSingleNumberOne(t *testing.T) {
	nums := []int{2, 2, 1}
	res := singleNumber(nums)
	if res != 1 {
		t.Error("One fail expected 1, result: ", res)
	}
}

func TestSingleNumberTwo(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	res := singleNumber(nums)
	if res != 4 {
		t.Error("One fail expected 4, result: ", res)
	}
}

func TestSingleNumberThree(t *testing.T) {
	nums := []int{1}
	res := singleNumber(nums)
	if res != 1 {
		t.Error("One fail expected 1, result: ", res)
	}
}
