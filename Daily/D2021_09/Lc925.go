package D2021_09

/*
583. 两个字符串的删除操作
给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。
*/

func Lc925MinDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	dp := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = max3(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]+1)
			} else {
				dp[i][j] = max3(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])
			}
		}
	}
	return l1 + l2 - 2*dp[l1][l2]
}

func max3(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}
