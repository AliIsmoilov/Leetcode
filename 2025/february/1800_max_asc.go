package february

// Given an array of positive integers nums,
// return the maximum possible sum of an ascending subarray in nums.

// Example 1:

// Input: nums = [10,20,30,5,10,50]
// Output: 65
// Explanation: [5,10,50] is the ascending subarray with the maximum sum of 65.
// Example 2:

// Input: nums = [10,20,30,40,50]
// Output: 150
// Explanation: [10,20,30,40,50] is the ascending subarray with the maximum sum of 150.

func MaxAscendingSum(nums []int) int {
	maxSum := 0
	sum := 0
	arrLen := len(nums)
	for i, v := range nums {
		sum += v
		if i < arrLen-1 {
			if nums[i+1] <= v {
				if sum > maxSum {
					maxSum = sum
				}
				sum = 0
				// fmt.Println(v)
			}
		}
		if i == arrLen-1 {
			if sum > maxSum {
				maxSum = sum
			}
			sum = 0
		}
	}

	return maxSum
}
