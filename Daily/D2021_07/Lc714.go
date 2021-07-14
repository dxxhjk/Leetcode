package D2021_07

import "sort"

/*
1818. 绝对差值和
给你两个正整数数组 nums1 和 nums2 ，数组的长度都是 n 。
数组 nums1 和 nums2 的 绝对差值和 定义为所有 |nums1[i] - nums2[i]|（0 <= i < n）的 总和（下标从 0 开始）。
你可以选用 nums1 中的 任意一个 元素来替换 nums1 中的 至多 一个元素，以 最小化 绝对差值和。
在替换数组 nums1 中最多一个元素 之后 ，返回最小绝对差值和。
因为答案可能很大，所以需要对 109 + 7 取余 后返回。、
 */

/*
思路：
二分查找，将 nums1 拷贝并排序，用来进行二分查找。
遍历 nums1 和 nums2，在计算据对差值的同时，
在拷贝有序 nums1 中找到能够使当前数字对差值最小的替换方案，
并在遍历过程中维护能够减小的最大值 max，
到最后 sum - max 即为答案。
 */

func Lc714MinAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	rec := append(sort.IntSlice(nil), nums1...)
	rec.Sort()
	sum, maxn, n := 0, 0, len(nums1)
	for i, v := range nums2 {
		diff := abs(nums1[i] - v)
		sum += diff
		j := rec.Search(v)
		if j < n {
			maxn = max(maxn, diff-(rec[j]-v))
		}
		if j > 0 {
			maxn = max(maxn, diff-(v-rec[j-1]))
		}
	}
	return (sum - maxn) % (1e9 + 7)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}