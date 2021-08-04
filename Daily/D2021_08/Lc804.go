package D2021_08

import "sort"

/*
611. 有效三角形的个数
给定一个包含非负整数的数组，你的任务是统计其中可以组成三角形三条边的三元组个数。
 */

/*
思路：
排序+双指针
 */

func Lc804TriangleNumber(nums []int) int {
	ans := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		k := i
		for j := i + 1; j < len(nums); j++ {
			for k < len(nums) && nums[k] < nums[i] + nums[j] {
				k++
			}
			if k - j > 1 {
				ans += k - j
			}
		}
	}
	return ans
}