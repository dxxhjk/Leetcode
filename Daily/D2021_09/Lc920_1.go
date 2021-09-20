package D2021_09

import "fmt"

/*
300. 最长递增子序列
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
1 <= nums.length <= 2500
-10^4 <= nums[i] <= 10^4
*/

func Lc920_1LengthOfLIS(nums []int) int {
	nums = append(nums, 10001)
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		tempMax := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j] > tempMax {
					tempMax = dp[j]
				}
			}
		}
		dp[i] = tempMax + 1
		fmt.Println(dp)
	}
	return dp[len(dp)-1] - 1
}
