package FirstUniqInString

/*
Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
*/
func firstUniqChar(s string) int {
	m := make(map[byte]int)

	for _, val := range []byte(s) {
		m[val]++
	}

	for i, val := range []byte(s) {
		if m[val] == 1 {
			return i
		}
	}

	return -1
}
