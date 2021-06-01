package D2021_05

/*
560. 和为K的子数组
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
 */

/*
思路：
前缀和+哈希表
由 pre[i] - pre[j - 1] = k 可得
遍历到 num[i] 的时候操作为
	sum += h[pre[i] - k]
	h[pre[i]]++
*/

func Lc529SubSubarraySum(nums []int, k int) int {
	h := make(map[int]int, 0)
	h[0] = 1
	pre := 0
	sum := 0
	for _, num := range nums {
		pre += num
		sum += h[pre - k]
		h[pre]++
	}
	return sum
}