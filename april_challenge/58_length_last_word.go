package aprilchallenge

// Example 1:

// Input: s = "Hello World"
// Output: 5
// Explanation: The last word is "World" with length 5.
// Example 2:

// Input: s = "   fly me   to   the moon  "
// Output: 4
// Explanation: The last word is "moon" with length 4.
func LengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " && count == 0 {
			continue
		} else {
			if string(s[i]) == " " {
				break
			}
			count++
		}
	}
	return count
}
