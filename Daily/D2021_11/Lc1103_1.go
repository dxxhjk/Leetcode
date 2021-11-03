package D2021_11

/*
接雨水1
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 */

func Lc1103_1Trap(height []int) int {
	ans := 0
	l, r := 0, len(height) - 1
	leftMax, rightMax := height[l], height[r]
	for l < r {
		if leftMax <= rightMax {
			ans += leftMax - height[l]
			l++
			if height[l] > leftMax {
				leftMax = height[l]
			}
		} else {
			ans += rightMax - height[r]
			r--
			if height[r] > rightMax {
				rightMax = height[r]
			}
		}
	}
	return ans
}
