package D2021_05

/*
1035. 不相交的线
在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足满足：
	nums1[i] == nums2[j]
	且绘制的直线不与任何其他连线（非水平线）相交。
请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。
以这种方法绘制线条，并返回可以绘制的最大连线数。
 */

/*
思路：
最长重复字串，DP
时间O(n2)
空间O(n2)
 */

func max3(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Lc521MaxUncrossedLines(nums1 []int, nums2 []int) int {
	l1, l2 := len(nums1), len(nums2)
	if l1 == 0 || l2 == 0 {
		return 0
	}
	dp := make([][]int, l1)
	for i := 0; i < l1; i++ {
		dpLine := make([]int, l2)
		dp[i] = dpLine
	}
	if nums1[0] == nums2[0] {
		dp[0][0] = 1
	}
	for i := 1; i < l2; i++ {
		if dp[0][i - 1] == 1 || nums1[0] == nums2[i] {
			dp[0][i] = 1
		}
	}
	for i := 1; i < l1; i++ {
		if dp[i - 1][0] == 1 || nums1[i] == nums2[0] {
			dp[i][0] = 1
		}
	}
	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = max3(dp[i - 1][j], dp[i][j - 1], dp[i - 1][j - 1] + 1)
			} else {
				dp[i][j] = max2(dp[i - 1][j], dp[i][j - 1])
			}
		}
	}
	return dp[l1 - 1][l2 - 1]
}

