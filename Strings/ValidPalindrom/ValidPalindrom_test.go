package ValidPalindrom

import "testing"

func TestValidPalindrom(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	if !isPalindrome(s) {
		t.Error("fail 1")
	}

	s = "race a car"
	if isPalindrome(s) {
		t.Error("fail 2")
	}

	s = " "
	if !isPalindrome(s) {
		t.Error("fail 3")
	}
}
