package aprilchallenge

// Input: s = "codeleet", indices = [4,5,6,7,0,2,1,3]
// Output: "leetcode"
// Explanation: As shown, "codeleet" becomes "leetcode" after shuffling.
// Example 2:

// Input: s = "abc", indices = [0,1,2]
// Output: "abc"
// Explanation: After shuffling, each character remains in its position.

func RestoreString(s string, indices []int) string {
	res := make([]byte, len(indices))
	//    fmt.Println(res)

	for i, v := range indices {
		res[v] = s[i]
	}

	return string(res)
}
