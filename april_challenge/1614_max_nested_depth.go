package aprilchallenge

// Example 1:

// Input: s = "(1+(2*3)+((8)/4))+1"
// Output: 3
// Explanation: Digit 8 is inside of 3 nested parentheses in the string.
// Example 2:

// Input: s = "(1)+((2))+(((3)))"
// Output: 3
func MaxDepth(s string) int {

	count := 0
	maxDepth := 0

	for _, ch := range s {
		if string(ch) == "(" {
			count++
			if count > maxDepth {
				maxDepth = count
			}
		} else if string(ch) == ")" {
			count--
		}
	}

	return maxDepth
}
