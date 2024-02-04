package programmingskills

// Input: word1 = "abc", word2 = "pqr"
// Output: "apbqcr"
// Explanation: The merged string will be merged as so:
// word1:  a   b   c
// word2:    p   q   r
// merged: a p b q c r

func MergeAlternately(word1 string, word2 string) string {

	result := ""
	i := 0
	for i = range word1 {
		if i == len(word2) {
			return result + word1[i:]
		}
		result += string(word1[i]) + string(word2[i])
	}
	return result + word2[i+1:]
}
