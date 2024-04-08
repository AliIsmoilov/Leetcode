package aprilchallenge

// Example 1:

// Input: nums = [2,2,1]
// Output: 1
// Example 2:

// Input: nums = [4,1,2,1,2]
// Output: 4
// Example 3:

// Input: nums = [1]
// Output: 1

func SingleNumber(nums []int) int {
	numsMap := make(map[int]int)
	for _, num := range nums {
		numsMap[num]++
	}

	//fmt.Println(numsMap)
	for _, num := range nums {
		v := numsMap[num]
		if v == 1 {
			return num
		}
	}

	return -1
}
