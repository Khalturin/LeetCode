package LeetCode

import "testing"

func TestIntersectionOne(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	rnums := intersect(nums1, nums2)
	if rnums[0] != 2 && rnums[1] != 2 {
		t.Error("One fail, expected [2, 2], res: ", rnums)
	}
}

func TestIntersectionTwo(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	rnums := intersect(nums1, nums2)
	if rnums[0] != 9 && rnums[1] != 4 {
		t.Error("One fail, expected [9, 4], res: ", rnums)
	}
}
