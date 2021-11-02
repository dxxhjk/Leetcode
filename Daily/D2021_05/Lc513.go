package D2021_05

/*
1269. 停在原地的方案数
有一个长度为 arrLen 的数组，开始有一个指针在索引 0 处。
每一步操作中，你可以将指针向左或向右移动 1 步，或者停在原地（指针不能被移动到数组范围外）。
给你两个整数 steps 和 arrLen ，请你计算并返回：在恰好执行 steps 次操作以后，指针仍然指向索引 0 处的方案数。
由于答案可能会很大，请返回方案数 模 10^9 + 7 后的结果。
 */

/*
思路：
DP，搞一个二维数组，i: step，j: pos
DP(i, j) = DP(i - 1, j - 1) + DP(i - 1, j) + DP(i - i, j + 1)
 */

func Lc513NumWays(steps int, arrLen int) int {
	if steps == 1 {
		return 1
	}
	if arrLen > steps {
		arrLen = steps
	}
	dp := make([][]int, steps)
	for i := 0; i < steps; i++ {
		dpLine := make([]int, arrLen)
		dp[i] = dpLine
	}
	dp[0][0], dp[0][1] = 1, 1
	for i := 1; i < steps; i++ {
		for j := 0; j < arrLen; j++ {
			dp[i][j] = dp[i - 1][j]
			if j != 0 {
				dp[i][j] += dp[i - 1][j - 1]
				dp[i][j] %= 1000000007
			}
			if j != arrLen - 1 {
				dp[i][j] += dp[i - 1][j + 1]
				dp[i][j] %= 1000000007
			}
		}
	}
	return dp[steps - 1][0]
}