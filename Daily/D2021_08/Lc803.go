package D2021_08

/*
581. 最短无序连续子数组
给你一个整数数组 nums ，你需要找出一个连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
请你找出符合题意的最短子数组，并输出它的长度。
 */

func Lc803FindUnsortedSubarray(nums []int) int {
	n := len(nums)
	l, r, ans := 0, n - 1, n
	for l < n - 1 && nums[l] <= nums[l + 1] {
		l++
	}
	for r >= 1 && nums[r] >= nums[r - 1] {
		r--
	}
	if l == n - 1 || r == 0 {
		return 0
	}
	min, max := nums[l], nums[r]
	for i := l; i < r + 1; i++ {
		if max < nums[i] {
	 		max = nums[i]
		}
		if min > nums[i] {
			min = nums[i]
		}
	}
	for i := 0; i < n && nums[i] <= min; i++ {
		ans--
	}
	for i := n - 1; i >= 0 && nums[i] >= max; i-- {
		ans--
	}
	return ans
}