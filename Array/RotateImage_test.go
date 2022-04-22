package LeetCode

import "testing"

// Equal проверяет, что a и b содержат одинаковые элементы.
// nil аргумент эквивалентен пустому срезу.
func Equal2D(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestRotateImageOne(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	etalon := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	res := Rotate(matrix)
	if !Equal2D(etalon, res) {
		t.Error("fail TestRotateImageOne, \nexpected: \n", etalon, "\nres: \n", res)
	}
}

func TestRotateImageTwo(t *testing.T) {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	etalon := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}
	res := Rotate(matrix)

	if !Equal2D(etalon, res) {
		t.Error("fail TestRotateImageTwo, \nexpected: \n", etalon, "\nres: \n", res)
	}
}
