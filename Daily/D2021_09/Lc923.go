package D2021_09

/*
326. 3的幂
给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。
整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x
*/

func Lc923IsPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n != 1 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return true
}
