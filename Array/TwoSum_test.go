package LeetCode

import "testing"

func TestTwoSumOne(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := TwoSum(nums, target)
	if res[0] != 0 || res[1] != 1 {
		t.Error("One fail, expected [0,1], result: ", res)
	}
}

func TestTwoSumTwo(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	res := TwoSum(nums, target)
	if res[0] != 1 || res[1] != 2 {
		t.Error("One fail, expected [1,2], result: ", res)
	}
}

func TestTwoSumThree(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	res := TwoSum(nums, target)
	if res[0] != 0 || res[1] != 1 {
		t.Error("One fail, expected [0,1], result: ", res)
	}
}
