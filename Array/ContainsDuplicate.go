package main

/*
	Given an integer array nums, return true if any value appears at least twice in the array, and return false if every
	 element is distinct.
*/

//func containsDuplicate(nums []int) bool {
//	for i, _ := range nums {
//		for j, _ := range nums {
//			if i == j {
//				continue
//			}
//			if nums[i] == nums[j] {
//				return true
//			}
//		}
//	}
//	return false
//}

func containsDuplicate(nums []int) bool {
	mapa := make(map[int]struct{})
	for _, val := range nums {
		if _, inMap := mapa[val]; inMap {
			return true
		}
		mapa[val] = struct{}{}
	}
	return false
}
