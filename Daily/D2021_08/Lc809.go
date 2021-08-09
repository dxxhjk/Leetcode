package D2021_08

/*
313. 超级丑数
超级丑数 是一个正整数，并满足其所有质因数都出现在质数数组 primes 中。
给你一个整数 n 和一个整数数组 primes ，返回第 n 个 超级丑数 。
题目数据保证第 n 个 超级丑数 在 32-bit 带符号整数范围内。
 */

func getMinAndPos(nums []int) (int, int) {
	min, pos := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
			pos = i
		}
	}
	return min, pos
}

func Lc809NthSuperUglyNumber(n int, primes []int) int {
	poses := make([]int, len(primes))
	allNums := make([]int, 1)
	allNums[0] = 1
	for i := 0; i < n; {
		nums := make([]int, 0)
		for j := 0; j < len(primes); j++ {
			nums = append(nums, allNums[poses[j]] * primes[j])
		}
		min, pos := getMinAndPos(nums)
		poses[pos]++
		if min != allNums[len(allNums) - 1] {
			allNums = append(allNums, min)
			i++
		}
	}
	return allNums[len(allNums) - 2]
}