package main

import "testing"

func TestSimpleROTATE(t *testing.T) {
	// ### TEST 1
	nums := []int{1, 2, 3}
	k := 1
	res := rotateRight(nums, k)
	if res[0] != 3 ||
		res[1] != 1 ||
		res[2] != 2 {
		t.Error("ROTATE 1 ERROR expected: ", nums, "result: ", res)
	}

	// ### TEST 2
	nums = []int{1, 2, 3}
	k = 2
	res = rotateRight(nums, k)

	if res[0] != 2 ||
		res[1] != 3 ||
		res[2] != 1 {
		t.Error("ROTATE 1 ERROR expected: ", nums, "result: ", res)
	}

	// ### TEST 3
	nums = []int{1, 2, 3}
	k = 3
	res = rotateRight(nums, k)

	if res[0] != 1 ||
		res[1] != 2 ||
		res[2] != 3 {
		t.Error("ROTATE 1 ERROR expected: ", nums, "result: ", res)
	}
}

func TestROTATE_ONE(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	out := make([]int, len(nums), len(nums))
	copy(out, nums)
	res := rotateRight(out, k)

	if res[0] != 5 ||
		res[1] != 6 ||
		res[2] != 7 ||
		res[3] != 1 ||
		res[4] != 2 ||
		res[5] != 3 ||
		res[6] != 4 {
		t.Error("ROTATE ONE ERROR expected: ", nums, "result: ", res)
	}
}

func TestROTATE_TWO(t *testing.T) {
	nums := []int{-1, -100, 3, 99}
	k := 2

	out := make([]int, len(nums), len(nums))
	copy(out, nums)
	res := rotateRight(out, k)

	if res[0] != 3 ||
		res[1] != 99 ||
		res[2] != -1 ||
		res[3] != -100 {
		t.Error("ROTATE TWO ERROR expected: ", nums, "result: ", res)
	}
}
