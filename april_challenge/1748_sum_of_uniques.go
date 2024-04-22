package aprilchallenge

// Example 1:

// Input: nums = [1,2,3,2]
// Output: 4
// Explanation: The unique elements are [1,3], and the sum is 4.
// Example 2:

// Input: nums = [1,1,1,1,1]
// Output: 0
// Explanation: There are no unique elements, and the sum is 0.
// Example 3:

// Input: nums = [1,2,3,4,5]
// Output: 15
// Explanation: The unique elements are [1,2,3,4,5], and the sum is 15.

func SumOfUnique(nums []int) int {
	m := make(map[int]int)
	sum := 0

	for _, num := range nums {
		m[num]++
	}
	// fmt.Println(m)

	for _, num := range nums {
		v, _ := m[num]
		if v == 1 {
			sum += num
		}
	}

	return sum
}
