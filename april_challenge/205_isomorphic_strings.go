package aprilchallenge

// Example 1:

// Input: s = "egg", t = "add"
// Output: true
// Example 2:

// Input: s = "foo", t = "bar"
// Output: false

func IsIsomorphic(s string, t string) bool {
	mapS := make(map[string]int)
	mapT := make(map[string]int)

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {

		if mapS[string(s[i])] != mapT[string(t[i])] {
			return false
		}

		mapS[string(s[i])] = i + 1
		mapT[string(t[i])] = i + 1
	}
	return true
}
