package LeetCode

import "testing"

// Equal проверяет, что a и b содержат одинаковые элементы.
// nil аргумент эквивалентен пустому срезу.
func EqualMass(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestRemoveDuplicatesOne(t *testing.T) {
	nums := []int{1, 1, 2}
	etalon := []int{1, 2}
	res := removeDuplicates(nums)

	if res != len(etalon) {
		t.Error("fail TestRemoveDuplicatesOne, expected: ", etalon, "result: ", res)
	}
}

func TestRemoveDuplicatesTwo(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	etalon := []int{0, 1, 2, 3, 4}
	res := removeDuplicates(nums)

	if res != len(etalon) {
		t.Error("fail TestRemoveDuplicatesTwo, expected: ", etalon, "result: ", res)
	}
}
