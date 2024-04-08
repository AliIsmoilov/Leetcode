package aprilchallenge

func RemoveDuplicates2(nums []int) int {

	i := 1
	for _, val := range nums {
		if val != nums[i-1] {
			nums[i] = val
			i++
		}
	}
	return i
}

func RemoveDuplicates(nums []int) int {

	numsMap := make(map[int]int)
	newList := []int{}
	for _, num := range nums {
		_, ok := numsMap[num]
		if !ok {
			newList = append(newList, num)
		}
		numsMap[num]++
	}

	copy(nums, newList)
	return len(newList)
}
