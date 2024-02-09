package februarychallenge

// Example 1:

// Input: nums = [1,3,4,2,2]
// Output: 2

func FindDuplicate(nums []int) int {
	numsMap := make(map[int]int)
	for _, num := range nums {
		_, ok := numsMap[num]
		if ok {
			return num
		} else {
			numsMap[num] += 1
		}
	}
	return -1
}
