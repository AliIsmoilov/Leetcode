package aprilchallenge

// Example 1:

// Input: words = ["pay","attention","practice","attend"], pref = "at"
// Output: 2
// Explanation: The 2 strings that contain "at" as a prefix are: "attention" and "attend".
// Example 2:

// Input: words = ["leetcode","win","loops","success"], pref = "code"
// Output: 0
// Explanation: There are no strings that contain "code" as a prefix.

func PrefixCount(words []string, pref string) int {
	count := 0

	for _, word := range words {
		if len(word) < len(pref) {
			continue
		}
		if word[:len(pref)] == pref {
			count++
		}
	}

	return count
}
