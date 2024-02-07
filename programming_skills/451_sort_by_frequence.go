package programmingskills

import (
	"sort"
)

// Input: s = "tree"
// Output: "eert"
// Explanation: 'e' appears twice while 'r' and 't' both appear once.
// So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.

func FrequencySort(s string) string {

	type pair struct {
		value rune
		freq  int
	}

	freq := make(map[rune]int)
	for _, v := range s {
		freq[v]++
	}
	freqArray := make([]pair, 0)
	for k, v := range freq {
		freqArray = append(freqArray, pair{k, v})
	}

	sort.Slice(freqArray, func(i, j int) bool {
		return freqArray[i].freq > freqArray[j].freq
	})

	result := make([]rune, 0)
	for _, v := range freqArray {
		for i := 0; i < v.freq; i++ {
			result = append(result, v.value)
		}
	}

	return string(result)
}
