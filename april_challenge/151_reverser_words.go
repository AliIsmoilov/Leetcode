package aprilchallenge

import (
	"fmt"
	"strings"
)

// Example 1:

// Input: s = "the sky is blue"
// Output: "blue is sky the"
// Example 2:

// Input: s = "  hello world  "
// Output: "world hello"
// Explanation: Your reversed string should not contain leading or trailing spaces.

func ReverseWords(s string) string {
	strList := strings.Split(s, " ")
	fmt.Println(strList)
	res := ""
	for _, ch := range strList {
		fmt.Println(">", ch)
	}

	for i := len(strList) - 1; i >= 0; i-- {
		if strList[i] != "" {
			res += strList[i]
			if i != len(strList)-1 {
				res += " "
			}
		}
	}

	return res
}

func ReverseWords2(s string) string {
	words := strings.Fields(s)
	fmt.Println(words)

	left, right := 0, len(words)-1

	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.Join(words, " ")
}
