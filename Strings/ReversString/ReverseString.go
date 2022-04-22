package ReversString

import "fmt"

/*
ReversString
	Write a function that reverses a string. The input string is given as an array of characters s.
	You must do this by modifying the input array in-place with O(1) extra memory.
*/
func ReversString(s []byte) []byte {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}

func reverseString(s []byte) {
	fmt.Println(ReversString(s))
}
