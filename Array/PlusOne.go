package main

/*
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the
integer.
The digits are ordered from most significant to least significant in left-to-right order.
The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.

Constraints:

1 <= digits.length <= 100
0 <= digits[i] <= 9
digits does not contain any leading 0's.
*/

func plusOne(digits []int) []int {
	lastIdx := len(digits) - 1
	for i := lastIdx; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			break
		} else {
			if i == 0 {
				newSlice := make([]int, lastIdx+2, lastIdx+2)
				newSlice[0] = 1
				digits[i] = 0
				for j, _ := range digits {
					newSlice[j+1] = digits[i]
					return newSlice
				}
			} else {
				digits[i] = 0
			}
		}
	}
	return digits
}

// [0] 9
// [1] 9
// [2] 9

// [0] 1
// [1] 0
// [2] 0
// [3] 0

//func plusOne(digits []int) []int {
//
//	return digits
//}
