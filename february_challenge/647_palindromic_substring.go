package februarychallenge

func CountSubstrings(s string) int {
	var res int

	m := make(map[byte][]int)

	for i := 0; i < len(s); i++ {
		m[s[i]] = append(m[s[i]], i)

		if len(m[s[i]]) > 1 {
			for _, val := range m[s[i]] {
				if val != i && IsPalindrome(s[val:i+1]) {
					res++
				}
			}
		}
	}

	return res + len(s)
}

func IsPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
