package february

// 3174. Clear Digits
// Easy

// You are given a string s.
// Your task is to remove all digits by doing this operation repeatedly:
// Delete the first digit and the closest non-digit character to its left.
// Return the resulting string after removing all digits.

// Example 1:
// Input: s = "abc"
// Output: "abc"

// Explanation:
// There is no digit in the string.

// Example 2:
// Input: s = "cb34"
// Output: ""

// Explanation:
// First, we apply the operation on s[2], and s becomes "c4".
// Then we apply the operation on s[1], and s becomes "".

func ClearDigits(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			result = result[:len(result)-1]
		} else {
			result = result + string(ch)
		}
	}
	return result
}
