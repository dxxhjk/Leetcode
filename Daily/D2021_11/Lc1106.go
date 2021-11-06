package D2021_11

/*
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
 */

func missingNumber(nums []int) int {
	ans := 0
	for i, num := range nums {
		ans += i - num
	}
	return ans + len(nums)
}
