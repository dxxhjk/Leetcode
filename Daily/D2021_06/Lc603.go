package D2021_06

/*
525. 连续数组
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
 */

/*
思路：
前缀和加哈希表
前缀和是 1 的个数减 0 的个数，哈希存前缀和对应的最早出现的索引。
 */

func Lc603FindMaxLength(nums []int) int {
	h := make(map[int]int)
	h[0] = -1
	pre := 0
	ans := 0
	for i, num := range nums {
		if num == 1 {
			pre++
		} else {
			pre--
		}
		index, ok := h[pre]
		if ok {
			if i - index > ans {
				ans = i - index
			}
		} else {
			h[pre] = i
		}
	}
	return ans
}
