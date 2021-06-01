package D2021_05

import (
	"fmt"
	"math"
)

/*
664. 奇怪的打印机
有台奇怪的打印机有以下两个特殊要求：
打印机每次只能打印由同一个字符组成的序列。
每次可以在任意起始和结束位置打印新字符，并且会覆盖掉原来已有的字符。
给你一个字符串 s ，你的任务是计算这个打印机打印它需要的最少打印次数。
 */

/*
思路：
动态规划，记录从 i 到 j 的最少操作数。
若 s[i] == s[j] 则 f[i][j] = f[i][j - 1]
否则 f[i][j] = f[i][k] + f[k + 1][j](k: i -> j)中的最小者
 */

func show(nums [][]int) {
	for _, i := range nums {
		fmt.Println(i)
	}
}

func Lc524StrangePrinter(s string) int {
	n := len(s)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		f[i][i] = 1
		for j := i + 1; j < n; j++ {
			fmt.Println(i,j)
			if s[i] == s[j] {
				f[i][j] = f[i][j-1]
			} else {
				f[i][j] = math.MaxInt64
				for k := i; k < j; k++ {
					f[i][j] = min(f[i][j], f[i][k]+f[k+1][j])
				}
			}
			show(f)
		}
	}
	return f[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
