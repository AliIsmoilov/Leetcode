package februarychallenge

func NumSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = n
	}

	for i := 1; i <= n; i++ {
		for j := 0; j*j <= i; j++ {
			if dp[i] > dp[i-j*j]+1 {
				dp[i] = dp[i-j*j] + 1
			}
		}
	}

	return dp[n]
}
