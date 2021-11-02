package D2021_05

import "fmt"

/*
342. 4的幂
给定一个整数，写一个函数来判断它是否是 4 的幂次方。
 */

func Lc531IsPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	for i := 31; i >= 0; i-- {
		fmt.Print((n & (1 << i))>>i)
	}
	if n & 0xAAAAAAAA != 0 {
		return false
	}
	sum := 0
	for n != 0 {
		sum += n & 1
		n >>= 1
	}
	return sum == 1
}
