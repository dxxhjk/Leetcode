package Daily

/*
477. 汉明距离总和
两个整数的汉明距离指的是这两个数字的二进制数对应位不同的数量。
计算一个数组中，任意两个数之间汉明距离的总和。
 */

/*
思路：
建立一个32位的数组 bit，每次遍历一个数 num[i]，按位处理，
num[i][j] = 1 则 sum += i - bit[j],
num[i][j] = 0 则 sum += bit[j],
并且将它为 1 的位对应的 bit 数组元素加一
 */

func Lc528TotalHammingDistance(nums []int) int {
	sum := 0
	bit := make([]int, 32)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < 32; j++ {
			b := nums[i] & 1
			if b == 0 {
				sum += bit[j]
			} else {
				sum += i - bit[j]
				bit[j] += 1
			}
			nums[i] >>= 1
		}
	}
	return sum
}