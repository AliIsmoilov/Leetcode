package aprilchallenge

import "math"

// Example 1:

// Input: x = 123
// Output: 321
// Example 2:

// Input: x = -123
// Output: -321

// var maxInt32 int32
// maxInt32 = 2147483648

func Reverse(x int) int {
	result := 0

	for x != 0 {
		lastNum := x % 10
		x /= 10
		result = result*10 + lastNum

		if result > math.MaxInt32 && result < math.MinInt32 {
			return 0
		}
	}

	return result
}
