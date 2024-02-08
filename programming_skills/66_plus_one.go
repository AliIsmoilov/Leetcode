package programmingskills

import (
	"strconv"
)

// Input: digits = [1,2,3]
// Output: [1,2,4]
// Explanation: The array represents the integer 123.
// Incrementing by one gives 123 + 1 = 124.
// Thus, the result should be [1,2,4].

func PlusOne(digits []int) []int {
	num := ""
	for _, digit := range digits {
		num += strconv.Itoa(digit)
	}
	num2, _ := strconv.Atoi(num)
	num2 += 1
	str := strconv.Itoa(num2)
	newArr := []int{}
	for _, digit := range str {
		n := string(digit)
		num, _ := strconv.Atoi(n)
		newArr = append(newArr, num)
	}
	return newArr
}
