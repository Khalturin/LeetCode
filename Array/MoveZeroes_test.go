package LeetCode

import "testing"

func TestMoveZeroesOne(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	res := MoveZeroesForTest(nums)
	if res[0] != 1 ||
		res[1] != 3 ||
		res[2] != 12 ||
		res[3] != 0 ||
		res[4] != 0 {
		t.Error("TestMoveZeroesOne fail, expected = [1, 3, 12, 0, 0], res: ", res)
	}
}

func TestMoveZeroesTwo(t *testing.T) {
	nums := []int{0}
	res := MoveZeroesForTest(nums)
	if res[0] != 0 {
		t.Error("TestMoveZeroesOne fail, expected = [0], res: ", res)
	}
}

func TestMoveZeroesThree(t *testing.T) {
	nums := []int{0, 0}
	res := MoveZeroesForTest(nums)
	if res[0] != 0 ||
		res[1] != 0 {
		t.Error("TestMoveZeroesOne fail, expected = [0, 0], res: ", res)
	}
}

func TestMoveZeroesFore(t *testing.T) {
	nums := []int{1, 0}
	res := MoveZeroesForTest(nums)
	if res[0] != 1 ||
		res[1] != 0 {
		t.Error("TestMoveZeroesOne fail, expected = [1, 0], res: ", res)
	}
}

func TestMoveZeroesFive(t *testing.T) {
	nums := []int{0, 1}
	res := MoveZeroesForTest(nums)
	if res[0] != 1 ||
		res[1] != 0 {
		t.Error("TestMoveZeroesOne fail, expected = [1, 0], res: ", res)
	}
}

func TestMoveZeroesSix(t *testing.T) {
	nums := []int{0, 1, 0, 0, 0, 1}
	res := MoveZeroesForTest(nums)
	if res[0] != 1 ||
		res[1] != 1 ||
		res[2] != 0 ||
		res[3] != 0 ||
		res[4] != 0 ||
		res[5] != 0 {
		t.Error("TestMoveZeroesOne fail, expected = [1, 1, 0, 0, 0, 0], res: ", res)
	}
}
