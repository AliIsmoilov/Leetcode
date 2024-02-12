package februarychallenge

// The majority element is the element that appears
// more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

// Input: nums = [3,2,3]
// Output: 3

func MajorityElement(nums []int) int {
	for _, num := range nums {
		count := countInList(nums, num)
		if count > len(nums)/2 {
			return num
		}
	}
	return -1
}

func countInList(nums []int, n int) int {
	count := 0
	for _, v := range nums {
		if v == n {
			count++
		}
	}
	return count
}
