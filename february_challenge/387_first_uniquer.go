package februarychallenge

// Input: s = "leetcode"
// Output: 0
func isUnique(s string, ch string) bool {
	count := 0
	for _, v := range s {
		if string(v) == ch {
			count++
			if count >= 2 {
				return false
			}
		}
	}
	if count >= 2 {
		return false
	} else {
		return true
	}
}

func FirstUniqChar(s string) int {
	for i, ch := range s {
		is_unique := isUnique(s, string(ch))
		if is_unique {
			return i
		}
	}

	return -1
}
