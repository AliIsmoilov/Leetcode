package february

// 2342. Max Sum of a Pair With Equal Sum of Digits
// Medium

// You are given a 0-indexed array nums consisting of positive integers.
// You can choose two indices i and j, such that i != j, and the
// sum of digits of the number nums[i] is equal to that of nums[j].

// Return the maximum value of nums[i] + nums[j] that you can obtain
// over all possible indices i and j that satisfy the conditions.

func MaximumSum(nums []int) int {
	var prevSum [82]int
	maxSum := -1
	var sum, n int
	for _, num := range nums {
		for sum, n = 0, num; n > 0; n = n / 10 {
			sum += n % 10
		}
		if prevSum[sum] != 0 {
			maxSum = max(maxSum, prevSum[sum]+num)
		}
		prevSum[sum] = max(prevSum[sum], num)
	}
	return maxSum
}
