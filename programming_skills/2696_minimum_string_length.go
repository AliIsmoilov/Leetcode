package programmingskills

func MinLength(s string) int {
	stack := make([]byte, 0, len(s))
	// ABFCACDB
	for _, char := range s {
		if len(stack) > 0 {
			last := stack[len(stack)-1]
			switch {
			case char == 'B' && last == 'A':
				stack = stack[:len(stack)-1]
			case char == 'D' && last == 'C':
				stack = stack[:len(stack)-1]
			default:
				stack = append(stack, byte(char))
			}
		} else {
			stack = append(stack, byte(char))
		}
	}

	return len(stack)
}
