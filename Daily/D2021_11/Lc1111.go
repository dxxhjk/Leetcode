package D2021_11

func Lc1111KInversePairs(n int, k int) int {
	if n == 0 {
		if k == 0 {
			return 1
		} else {
			return 0
		}
	}
	if k == 0 {
		return 1
	}
	dp := make([][]int, n + 1)
	for i := 0; i < n + 1; i++ {
		dp[i] = make([]int, k + 1)
		dp[i][0] = 1
	}
	for i := 1; i < n + 1; i++ {
		for j := 1; j < k + 1; j++ {
			if j < i {
				dp[i][j] = (dp[i][j - 1] + dp[i - 1][j]) % 1000000007
			} else {
				dp[i][j] = (dp[i][j - 1] - dp[i - 1][j - i] + dp[i - 1][j] + 10000000007) % 1000000007
			}
		}
	}
	return dp[n][k]
}
