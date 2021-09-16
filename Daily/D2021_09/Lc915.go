package D2021_09

/*
峰值元素是指其值严格大于左右相邻值的元素。
给你一个整数数组 nums，找到峰值元素并返回其索引。
数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
你可以假设 nums[-1] = nums[n] = -∞ 。
你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
*/

func Lc915FindPeakElement(nums []int) int {
	if len(nums) == 1 || nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-2] < nums[len(nums)-1] {
		return len(nums) - 1
	}
	l, r := 1, len(nums)-2
	mid := (l + r) / 2
	for !(nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1]) {
		if nums[mid-1] > nums[mid] {
			r = mid - 1
		} else if nums[mid] < nums[mid+1] {
			l = mid + 1
		}
		mid = (l + r) / 2
	}
	return mid
}
