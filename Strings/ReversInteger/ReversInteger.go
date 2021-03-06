package ReversInteger

import (
	"math"
)

/*
	Given a signed 32-bit integer x, return x with its digits reversed.
	If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
*/
func reverse(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return res
}
