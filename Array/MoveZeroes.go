package LeetCode

import "fmt"

// [0,1,0,3,12]
// 12, 3, 0
// 3, 12, 0
// 0, 1, 3, 12, 0
// 12, 3, 1, 0

// Лучшее решение
func MoveZeroesForTest(nums []int) []int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	return nums
}

// Моё решение
func moveZeroes(nums []int) {
	idxLast := len(nums) - 1
	if idxLast == 0 { // если элемент всего один
		fmt.Println(nums)
		return
	}

	idxLastZero := idxLast
	for i := idxLast; i >= 0; i-- {
		if nums[i] == 0 {
			reversArr(nums, i, idxLastZero)
			idxLastZero--
			reversArr(nums, i, idxLastZero)
		}
	}
	fmt.Println(nums)
}
