package ValidAnagram

import (
	"fmt"
	"sort"
	"strings"
)

/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the
original letters exactly once.
*/
func isAnagram(s string, t string) bool {
	m := make(map[byte]int)

	first := strings.Split(s, "")
	second := strings.Split(t, "")
	sort.Strings(first)
	sort.Strings(second)
	fmt.Println(strings.Join(first, ""))
	//fmt.Println(first, second)

	if len(s) != len(t) {
		return false
	}

	for _, val := range []byte(s) {
		m[val]++
	}

	for _, val := range []byte(t) {
		if _, ok := m[val]; ok {
			if m[val] <= 0 {
				return false
			}
			m[val]--
		} else {
			return false
		}
	}

	return true
}
