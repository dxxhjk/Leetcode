package D2021_05

import "fmt"

/*
231. 2 的幂
给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。
 */

func Lc530IsPowerOfTwo(n int) bool {
	if n < 0 {
		return false
	}
	sum := 0
	for n != 0 {
		fmt.Print(n & 1)
		sum += n & 1
		n >>= 1
	}
	return sum == 1
}
