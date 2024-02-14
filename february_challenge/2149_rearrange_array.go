package februarychallenge

// You are given a 0-indexed integer array nums of even length consisting of an equal number of positive and negative integers.

// You should rearrange the elements of nums such that the modified array follows the given conditions:

// Every consecutive pair of integers have opposite signs.
// For all integers with the same sign, the order in which they were present in nums is preserved.
// The rearranged array begins with a positive integer.

// Example 1:
// Input: nums = [3,1,-2,-5,2,-4]
// Output: [3,-2,1,-5,2,-4]

func RearrangeArray(nums []int) []int {
	positiveNums := []int{}
	negativeNums := []int{}
	for _, num := range nums {
		if num > 0 {
			positiveNums = append(positiveNums, num)
		} else {
			negativeNums = append(negativeNums, num)
		}
	}

	result := []int{}
	for i := 0; i < len(nums)/2; i++ {
		result = append(result, positiveNums[i])
		result = append(result, negativeNums[i])
	}
	return result
}
