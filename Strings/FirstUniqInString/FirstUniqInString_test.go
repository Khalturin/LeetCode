package FirstUniqInString

import "testing"

func TestFirstUniqInString1(t *testing.T) {
	s := "leetcode"
	expected := 0

	res := firstUniqChar(s)
	if res != expected {
		t.Errorf("fail 1, expected %d, result: %d", expected, res)
	}
}

func TestFirstUniqInString2(t *testing.T) {
	s := "loveleetcode"
	expected := 2

	res := firstUniqChar(s)
	if res != expected {
		t.Errorf("fail 2, expected %d, result: %d", expected, res)
	}
}

func TestFirstUniqInString3(t *testing.T) {
	s := "aabb"
	expected := -1

	res := firstUniqChar(s)
	if res != expected {
		t.Errorf("fail 3, expected %d, result: %d", expected, res)
	}
}
