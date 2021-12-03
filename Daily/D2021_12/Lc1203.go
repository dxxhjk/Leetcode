package D2021_12

import (
	"fmt"
	"sort"
)

/*
给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：
选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
重复这个过程恰好 k 次。可以多次选择同一个下标 i 。
以这种方式修改数组后，返回数组 可能的最大和 。
 */

func Lc1203LargestSumAfterKNegations(nums []int, k int) int {
	minus := make([]int, 0)
	minPlus := 0x3f3f3f3f
	sum := 0
	for _, num :=  range nums {
		if num >= 0 && num < minPlus {
			minPlus = num
		}
		if num < 0 {
			minus = append(minus, num)
			if -num < minPlus {
				minPlus = -num
			}
		}
		sum += num
	}
	sort.Ints(minus)
	fmt.Println(minus)
	if k >= len(minus) {
		for _, num := range minus {
			sum -= 2 * num
		}
		if (k - len(minus)) % 2 == 1 {
			sum -= 2 * minPlus
		}
	} else {
		for i := 0; i < k; i++ {
			sum -= 2 * minus[i]
		}
	}
	return sum
}
