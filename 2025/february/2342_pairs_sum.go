package february

func maximumSum(nums []int) int {
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
