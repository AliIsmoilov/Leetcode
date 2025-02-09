package february

// 2364. Count Number of Bad Pairs

// You are given a 0-indexed integer array nums.
// A pair of indices (i, j) is a bad pair if i < j and j - i != nums[j] - nums[i].

// Return the total number of bad pairs in nums.

func CountBadPairs(nums []int) int64 {
	badPairdCount := 0
	freqMap := make(map[int]int)

	for i, num := range nums {
		diff := num - i

		// Add the number of good pairs (those that already occurred) to totalBadPairs
		badPairdCount += i - freqMap[diff]

		// Increment the frequency of the difference
		freqMap[diff]++
	}

	return int64(badPairdCount)
}
