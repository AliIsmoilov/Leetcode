package februarychallenge

// 201. Bitwise AND of Numbers Range

// Given two integers left and right that represent
// the range [left, right], return the bitwise AND of all numbers in this range, inclusive.
func RangeBitwiseAnd(left int, right int) int {
	for left < right {
		right &= right - 1
	}
	return right
}
