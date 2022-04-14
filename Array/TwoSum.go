package LeetCode

// С помощью мапы
func TwoSum(nums []int, target int) []int {
	res := make([]int, 0, 2)
	mapa := make(map[int]int)

	for i, val := range nums {
		if _, inMap := mapa[target-val]; inMap {
			res = append(res, mapa[target-val], i)
			return res
		} else {
			mapa[val] = i
		}
	}
	return res
}

func twoSum(nums []int, target int) []int {
	return TwoSum(nums, target)
}
