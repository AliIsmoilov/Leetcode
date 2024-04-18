package aprilchallenge

import "strings"

// Example 1:

// Input: words = ["leet","code"], x = "e"
// Output: [0,1]
// Explanation: "e" occurs in both words: "leet", and "code". Hence, we return indices 0 and 1.
// Example 2:

// Input: words = ["abc","bcd","aaaa","cbc"], x = "a"
// Output: [0,2]
// Explanation: "a" occurs in "abc", and "aaaa". Hence, we return indices 0 and 2.

func FindWordsContaining(words []string, x byte) []int {
	resp := []int{}

	for i, word := range words {
		if strings.Contains(word, string(x)) {
			resp = append(resp, i)
		}
	}

	return resp
}
