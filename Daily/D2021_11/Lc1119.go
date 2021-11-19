package D2021_11

/*
397. 整数替换
给定一个正整数 n ，你可以做如下操作：
	如果 n 是偶数，则用 n / 2替换 n 。
	如果 n 是奇数，则可以用 n + 1或n - 1替换 n 。
n 变为 1 所需的最小替换次数是多少？
 */

func Lc1119IntegerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if n % 2 == 0 {
		return Lc1119IntegerReplacement(n / 2) + 1
	}
	if n % 4 == 1 {
		return Lc1119IntegerReplacement(n - 1) + 1
	}
	if n % 4 == 3 {
		return Lc1119IntegerReplacement(n + 1) + 1
	}
	return -1
}
