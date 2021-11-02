package D2021_06

import "fmt"

/*
474. 一和零
给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。
如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
 */

/*
思路：
dp：f(strs, m, n) = max(f(strs[:len(strs) - 1], m, n),
						f(strs[:len(strs) - 1], m - z, n - o) + 1)
即 dp[i][m][n] = max(dp[i - 1][m][n], dp[i - 1][m - zi][n - oi] + 1)
采用滚动数组优化掉一维:
逆向更新 dp 数组，dp[m][n] = max(dp[m][n], dp[m - zi][n - oi] + 1)
)
 */

func calculate(str string) (int, int) {
	zero, one := 0, 0
	for _, b := range str {
		if b == '0' {
			zero++
		} else {
			one++
		}
	}
	return zero, one
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Lc606FindMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m + 1)
	for i := 0; i <= m; i++ {
		dpLine := make([]int, n + 1)
		dp[i] = dpLine
	}
	for _, str := range strs {
		z, o := calculate(str)
		for i := m; i >= z; i-- {
			for j := n; j >= o; j-- {
				dp[i][j] = max(dp[i][j], dp[i - z][j - o] + 1)
			}
		}
		fmt.Println(dp)
	}
	return dp[m][n]
}