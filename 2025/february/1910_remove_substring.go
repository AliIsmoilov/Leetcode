package february

import "strings"

// 1910. Remove All Occurrences of a Substring
// Medium

// Given two strings s and part, perform the following operation on s
// until all occurrences of the substring part are removed:

// Find the leftmost occurrence of the substring part and remove it from s.
// Return s after removing all occurrences of part.
// A substring is a contiguous sequence of characters in a string.

func RemoveOccurrences(s string, part string) string {

	for {
		index := strings.Index(s, part)
		if index == -1 {
			break
		}
		s = s[:index] + s[index+len(part):]
	}
	return s
}
