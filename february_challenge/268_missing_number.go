package februarychallenge

import "fmt"

// Given an array nums containing n distinct numbers in the range [0, n],
// return the only number in the range that is missing from the array.

func MissingNumber(nums []int) int {
	orderedNums := []int{}
	for i := 0; i <= len(nums); i++ {
		orderedNums = append(orderedNums, i)
	}

	fmt.Println(orderedNums)
	for _, v := range orderedNums {
		if !isIncludes(nums, v) {
			return v
		}
	}
	return -1
}

func isIncludes(nums []int, n int) bool {
	for _, num := range nums {
		if num == n {
			return true
		}
	}
	return false
}
