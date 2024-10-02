package programmingskills

import "sort"

// Given an array of integers arr, replace each element with its rank.

// The rank represents how large the element is. The rank has the following rules:

// Rank is an integer starting from 1.
// The larger the element, the larger the rank. If two elements are equal, their rank must be the same.
// Rank should be as small as possible.

// Example 1:

// Input: arr = [40,10,20,30]
// Output: [4,1,2,3]
// Explanation: 40 is the largest element. 10 is the smallest. 20 is the second smallest. 30 is the third smallest.
// Example 2:

// Input: arr = [100,100,100]
// Output: [1,1,1]
// Explanation: Same elements share the same rank.

func arrayRankTransform(arr []int) []int {
	// Create a copy of the original array to sort and assign ranks
	sortedArr := append([]int{}, arr...)
	sort.Ints(sortedArr)

	// Create a map to store the rank of each unique element
	rankMap := make(map[int]int)
	rank := 1

	for _, num := range sortedArr {
		_, ok := rankMap[num]
		if !ok {
			rankMap[num] = rank
			rank++
		}
	}

	for i, num := range arr {
		arr[i] = rankMap[num]
	}

	return arr
}
