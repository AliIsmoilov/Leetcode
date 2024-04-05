package aprilchallenge

// Example 1:

// Input: nums = [1,3,5,6], target = 5
// Output: 2
// Example 2:

// Input: nums = [1,3,5,6], target = 2
// Output: 1
// Example 3:

// Input: nums = [1,3,5,6], target = 7
// Output: 4
func SearchInsert(nums []int, target int) int {
	for ind, val := range nums {
		if val == target {
			return ind
		} else if val > target {
			return ind
		}
	}

	return len(nums)
}
