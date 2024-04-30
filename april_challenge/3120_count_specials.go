package aprilchallenge

import "unicode"

// Example 1:

// Input: word = "aaAbcBC"

// Output: 3

// Explanation:

// The special characters in word are 'a', 'b', and 'c'.

// Example 2:

// Input: word = "abc"

// Output: 0

// Explanation:

// No character in word appears in uppercase.

func NumberOfSpecialChars(word string) int {
	smallLetters := make(map[rune]bool)
	capitalLetters := make(map[rune]bool)

	for _, ch := range word {
		if string(ch) >= "a" && string(ch) <= "z" {
			smallLetters[ch] = true
		} else {
			capitalLetters[ch] = true
		}
	}

	// fmt.Println(smallLetters)
	// fmt.Println(capitalLetters)

	count := 0

	for k, _ := range smallLetters {
		_, ok := capitalLetters[unicode.ToUpper(k)]
		if ok {
			count++
		}
	}

	return count
}
