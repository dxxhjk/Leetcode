package D2021_06

import "fmt"

/*
523. 连续的子数组和
给你一个整数数组 nums 和一个整数 k ，编写一个函数来判断该数组是否含有同时满足下述条件的连续子数组：
子数组大小至少为 2，且子数组元素总和为 k 的倍数。
如果存在，返回 true ；否则，返回 false 。
如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。
 */

/*
思路：
前缀和加哈希表
若 sum[i - 1] % k == sum[j] % k，则子数组元素和为k的倍数。
哈希存余数和索引之间的映射。
*/

func Lc602CheckSubarraySum(nums []int, k int) bool {
	h := make(map[int]int)
	pre := 0
	h[0] = 0
	for i, num := range nums {
		pre = (pre + num) % k
		index, ok := h[pre]
		if ok {
			if i - index > 0 {
				return true
			}
		} else {
			h[pre] = i + 1
		}
		fmt.Println(pre, h)
	}
	return false
}
