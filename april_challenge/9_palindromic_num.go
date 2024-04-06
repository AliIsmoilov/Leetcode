package aprilchallenge

func IsPalindrome(x int) bool {
	var after = 0
	var before = x

	for x > 0 {
		lastNum := x % 10
		after = after*10 + lastNum
		x /= 10
	}

	return after == before
}
