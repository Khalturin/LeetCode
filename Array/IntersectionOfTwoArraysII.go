package main

/*
	Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must
	appear as many times as it shows in both arrays and you may return the result in any order.
*/

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	r := make([]int, 0, 0)
	for _, val := range nums1 {
		if _, inMap := m[val]; inMap {
			m[val]++
		} else {
			m[val] = 1
		}
	}

	for _, val := range nums2 {
		if _, inMap := m[val]; inMap {
			if m[val] != 0 {
				r = append(r, val)
				m[val]--
			}
		}
	}

	return r
}
