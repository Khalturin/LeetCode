package LeetCode

import "fmt"

/*
	Given an array, rotate the array to the right by k steps, where k is non-negative.
*/

// Поворот массива влево
func rotateLeft(nums []int, k int) []int {
	for i := 0; i < k; i++ {
		f := nums[0]
		for j := 0; j < len(nums)-1; j++ {
			s := nums[j+1]
			nums[j] = s
		}
		nums[len(nums)-1] = f
	}
	return nums
}

// Поворот вправо простым перебором
//func rotateRight(nums []int, k int) []int {
//	for i := 0; i < k; i++ {
//		l := nums[len(nums)-1]
//		for j := len(nums) - 1; j > 0; j-- {
//			s := nums[j-1]
//			nums[j] = s
//		}
//		nums[0] = l
//	}
//	return nums
//}

// Повотор вправо с разделением массива на 2 части и заменой их местами
//func rotateRight(nums []int, k int) []int {
//	lenNums := len(nums)
//	lenNumsIdx := lenNums - 1
//
//	if lenNums < 2 {
//		fmt.Println(nums)
//		return nums
//	}
//	k = k % lenNums
//
//	left := make([]int, lenNums-k, lenNums-k)
//	copy(left, nums[:lenNums-k])
//
//	right := make([]int, k, k)
//	copy(right, nums[lenNums-k:])
//
//	for i := 0; i < k; i++ {
//		nums[i] = right[i]
//	}
//
//	j := 0
//	for i := k; i <= lenNumsIdx; i++ {
//		nums[i] = left[j]
//		j++
//	}
//
//	fmt.Println(nums)
//
//	return nums
//}

func reversArr(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// Самый дешевый вариант - с переворотом массива и каждой его половины
func rotateRight(nums []int, k int) []int {
	n := len(nums)
	if n < 2 {
		fmt.Println(nums)
		return nums
	}
	k = k % n

	reversArr(nums, 0, n-1)
	reversArr(nums, 0, k-1)
	reversArr(nums, k, n-1)

	fmt.Println(nums)

	return nums
}
