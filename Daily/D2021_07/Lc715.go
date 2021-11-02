package D2021_07

import "sort"

/*
1846. 减小和重新排列数组后的最大元素
给你一个正整数数组 arr 。请你对 arr 执行一些操作（也可以不进行任何操作），使得数组满足以下条件：
	arr 中第一个元素必须为 1 。
	任意相邻两个元素的差的绝对值小于等于 1。
你可以执行以下 2 种操作任意次：
	减小 arr 中任意元素的值，使其变为一个 更小的正整数 。
	重新排列 arr 中的元素，你可以以任意顺序重新排列。
请你返回执行以上操作后，在满足前文所述的条件下，arr 中可能的 最大值 。
 */

/*
思路：
排序加贪心
 */

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Lc75MaximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	arr[0] = 1
	for i := 1; i < len(arr); i++ {
		arr[i] = min(arr[i], arr[i - 1] + 1)
	}
	return arr[len(arr) - 1]
}