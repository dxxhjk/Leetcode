package D2021_07

/*
面试题 17.10. 主要元素
数组中占比超过一半的元素称之为主要元素。给你一个 整数 数组，找出其中的主要元素。
若没有，返回 -1 。请设计时间复杂度为 O(N) 、空间复杂度为 O(1) 的解决方案。
 */

func Lc709MajorityElement(nums []int) int {
	tempElem, tempNum := 0, 0
	for _, num := range nums {
		if tempNum == 0 {
			tempElem = num
			tempNum = 1
		} else {
			if num == tempElem {
				tempNum++
			} else {
				tempNum--
			}
		}
	}
	if tempNum != 0 {
		cnt := 0
		for _, num := range nums {
			if num == tempElem {
				cnt++
			}
		}
		if cnt > len(nums) / 2 {
			return tempElem
		}
	}
	return -1
}