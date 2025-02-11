package february

import "strings"

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
