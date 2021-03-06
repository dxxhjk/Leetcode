package D2021_05

/*
461. 汉明距离
两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
给出两个整数 x 和 y，计算它们之间的汉明距离。
0 ≤ x, y < 2^31.
 */

func Lc527HammingDistance(x int, y int) int {
	xor := x ^ y
	ans := 0
	for xor != 0 {
		ans += xor & 1
		xor >>= 1
	}
	return ans
}