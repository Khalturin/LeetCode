package main

/*
	Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
	You must implement a solution with a linear runtime complexity and use only constant extra space.
*/

func singleNumber(nums []int) int {
	mapa := map[int]struct{}{}
	for _, val := range nums {
		if _, inMap := mapa[val]; inMap {
			delete(mapa, val)
		} else {
			mapa[val] = struct{}{}
		}
	}

	for key := range mapa {
		return key
	}
	return 0
}
