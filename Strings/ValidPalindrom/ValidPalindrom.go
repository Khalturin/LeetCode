package ValidPalindrom

import (
	"strings"
	"unicode"
)

func formatString(s string) string {
	s = strings.ToLower(s)
	r := []rune(s)
	var res []rune
	for _, val := range r {
		if unicode.IsDigit(val) || unicode.IsLetter(val) {
			res = append(res, val)
		}
	}
	return string(res)
}

func isPalindromeString(s string) bool {
	lidx := len(s) - 1
	for i := 0; i <= lidx; i++ {
		if s[i] != s[lidx] {
			return false
		}
		lidx--
	}
	return true
}

func isPalindrome(s string) bool {
	res := formatString(s)

	return isPalindromeString(res)
}
