package LeetCode

import "testing"

func TestMaxProfitOne(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(nums)
	expected := 7
	if res != expected {
		t.Error("fail TestMaxProfitOne, expected: ", expected, "result: ", res)
	}
}

func TestMaxProfitTwo(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	res := maxProfit(nums)
	expected := 4
	if res != expected {
		t.Error("fail TestMaxProfitTwo, expected: ", expected, "result: ", res)
	}
}

func TestMaxProfitThree(t *testing.T) {
	nums := []int{7, 6, 4, 3, 1}
	res := maxProfit(nums)
	expected := 0
	if res != expected {
		t.Error("fail TestMaxProfitThree, expected: ", expected, "result: ", res)
	}
}
