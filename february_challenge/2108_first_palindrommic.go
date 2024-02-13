package februarychallenge

// Example 1:

// Input: words = ["abc","car","ada","racecar","cool"]
// Output: "ada"
// Explanation: The first string that is palindromic is "ada".
// Note that "racecar" is also palindromic, but it is not the first.

func FirstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindromic(word) {
			return word
		}
	}
	return ""
}

func isPalindromic(word string) bool {
	left := 0
	right := len(word) - 1
	for left < right {
		if word[left] != word[right] {
			return false
		}
		left++
		right--
	}
	return true
}
