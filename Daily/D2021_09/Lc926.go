package D2021_09

/*
371. 两整数之和
给你两个整数 a 和 b ，不使用 运算符 + 和 -，计算并返回两整数之和。
*/

func Lc926GetSum(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}
