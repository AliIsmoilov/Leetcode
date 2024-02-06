package programmingskills

import "sort"

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

func GroupAnagrams(strs []string) [][]string {
	keyToAnagrams := make(map[string][]string)

	for _, str := range strs {
		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		key := string(chars)
		keyToAnagrams[key] = append(keyToAnagrams[key], str)
	}

	var result [][]string
	for _, anagrams := range keyToAnagrams {
		result = append(result, anagrams)
	}

	return result
}
