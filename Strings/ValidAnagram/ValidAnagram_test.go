package ValidAnagram

import "testing"

func TestValidAnagram1(test *testing.T) {
	s := "anagram"
	t := "nagaram"
	if !isAnagram(s, t) {
		test.Error("fail, expected true, result false")
	}
}
func TestValidAnagram2(test *testing.T) {
	s := "rat"
	t := "car"
	if isAnagram(s, t) {
		test.Error("fail, expected false, result true")
	}
}

func TestValidAnagram3(test *testing.T) {
	s := "cat myaw"
	t := "myaw tac"
	if !isAnagram(s, t) {
		test.Error("fail, expected true, result false")
	}
}

func TestValidAnagram4(test *testing.T) {
	s := "ab"
	t := "a"
	if isAnagram(s, t) {
		test.Error("fail, expected false, result true")
	}
}
