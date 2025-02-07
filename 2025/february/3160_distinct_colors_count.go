package february

// 3160. Find the Number of Distinct Colors Among the Balls

// You are given an integer limit and a 2D array queries of size n x 2.

// There are limit + 1 balls with distinct labels in the range [0, limit].
// Initially, all balls are uncolored. For every query in queries that is of the form [x, y],
// you mark ball x with the color y.
// After each query, you need to find the number of distinct colors among the balls.

// Return an array result of length n, where result[i] denotes the
// number of distinct colors after ith query.

// Note that when answering a query, lack of a color will not be considered as a color.

// Example 1:

// Input: limit = 4, queries = [[1,4],[2,5],[1,3],[3,4]]

// Output: [1,2,2,3]

func QueryResults(limit int, queries [][]int) []int {

	res := make([]int, len(queries))
	ballColors := make(map[int]int)
	colorCount := make(map[int]int)

	for i := 0; i < len(queries); i++ {
		ball := queries[i][0]
		color := queries[i][1]

		if prevColor, exists := ballColors[ball]; exists {
			colorCount[prevColor]--
			if colorCount[prevColor] == 0 {
				delete(colorCount, prevColor)
			}
		}

		ballColors[ball] = color
		colorCount[color]++

		res[i] = len(colorCount)
	}

	// fmt.Println(ballColors)
	// fmt.Println(colorCount)

	return res
}
