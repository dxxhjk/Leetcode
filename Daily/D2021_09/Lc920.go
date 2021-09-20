package D2021_09

/*
673. 最长递增子序列的个数
给定一个未排序的整数数组，找到最长递增子序列的个数。
注意: 给定的数组长度不超过 2000 并且结果一定是32位有符号整数。
*/

func Lc920FindNumberOfLIS(nums []int) int {
	nums = append(nums, 10000001)
	dp := make([]int, len(nums))
	dp[0] = 1
	dpNum := make([]int, len(nums))
	dpNum[0] = 1
	for i := 1; i < len(nums); i++ {
		tempMax := 0
		tempNum := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j] > tempMax {
					tempMax = dp[j]
					tempNum = dpNum[j]
				} else if dp[j] == tempMax {
					tempNum += dpNum[j]
				}
			}
		}
		dp[i] = tempMax + 1
		dpNum[i] = tempNum
	}
	return dpNum[len(dpNum)-1]
}
