package D2022_01

import "fmt"

/*
给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。
如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；
否则，返回 false 。
 */

func Lc0112IncreasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first, second := nums[0], 0x7fffffff
	for _, num := range nums {
		fmt.Println(first, second, num)
		if num > second {
			return true
		}
		if num > first {
			second = num
			continue
		}
		first = num
	}
	return false
}
