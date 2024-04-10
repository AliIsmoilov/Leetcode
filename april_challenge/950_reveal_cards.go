package aprilchallenge

import (
	"sort"
)

// Example 1:

// Input: deck = [17,13,11,2,3,5,7]
// Output: [2,13,3,11,5,17,7]
func DeckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	// fmt.Println(deck)

	res := make([]int, len(deck))
	queue := make([]int, len(deck))

	for i := range queue {
		queue[i] = i
	}

	for _, card := range deck {
		idx := queue[0]
		res[idx] = card
		queue = queue[1:]

		if len(queue) > 0 {
			queue = append(queue, queue[0])
			queue = queue[1:]
		}
	}
	// fmt.Println(queue)

	return res
}
