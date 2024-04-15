package aprilchallenge

import (
	"fmt"
	"strconv"
)

// Example 1:

// Input: num = [1,2,0,0], k = 34
// Output: [1,2,3,4]
// Explanation: 1200 + 34 = 1234
// Example 2:

// Input: num = [2,7,4], k = 181
// Output: [4,5,5]
// Explanation: 274 + 181 = 455
// Example 3:

// Input: num = [2,1,5], k = 806
// Output: [1,0,2,1]
// Explanation: 215 + 806 = 1021

func AddToArrayForm(num []int, k int) []int {
	res := []int{}
	numStr := ""
	for _, n := range num {
		numStr += fmt.Sprint(n)
	}
	fmt.Println(numStr)
	number, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(number)
	number += k
	fmt.Println(number)
	numStr = fmt.Sprint(number)

	for _, ch := range numStr {
		temp, _ := strconv.Atoi(string(ch))
		res = append(res, temp)
	}

	return res
}
