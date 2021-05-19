package Daily

import (
	"Leetcode/Structs"
	"fmt"
)

/*
1738. 找出第 K 大的异或坐标值
给你一个二维矩阵 matrix 和一个整数 k ，矩阵大小为 m x n 由非负整数组成。
矩阵中坐标 (a, b) 的值可由对所有满足 0 <= i <= a < m 且 0 <= j <= b < n 的元素 matrix[i][j]（下标从 0 开始计数）执行异或运算得到。
请你找出 matrix 的所有坐标中第 k 大的值（k 的值从 1 开始计数）。

eg1:
输入：matrix = [[5,2],[1,6]], k = 1
输出：7
解释：坐标 (0,1) 的值是 5 XOR 2 = 7 ，为最大的值

eg2:
输入：matrix = [[5,2],[1,6]], k = 3
输出：4
解释：坐标 (1,0) 的值是 5 XOR 1 = 4 ，为第 3 大的值
 */

/*
思路：
看到最大 k 直接想到用堆来解
递推加堆
v(i, j) = v(i - 1, j) XOR v(i, j - 1) XOR (i - 1, j - 1) XOR m(i, j)
使用二维数组对 v 进行记录，同时维护大小为 k 的小根堆
 */

func Lc519KthLargestValue(matrix [][]int, k int) int {
	v := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		vLine := make([]int, len(matrix[0]))
		v = append(v, vLine)
	}
	h := Structs.MakeHeap(k)
	v[0][0] = matrix[0][0]
	h.Push(v[0][0])
	for i := 1; i < len(matrix); i++ {
		v[i][0] = matrix[i][0] ^ v[i - 1][0]
		h.Push(v[i][0])
	}
	for i := 1; i < len(matrix[0]); i++ {
		v[0][i] = matrix[0][i] ^ v[0][i - 1]
		h.Push(v[0][i])
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			v[i][j] = v[i - 1][j] ^ v[i][j - 1] ^ v[i - 1][j - 1] ^ matrix[i][j]
			h.Push(v[i][j])
		}
	}
	fmt.Println(v)
	fmt.Println(h)
	return h.GetMin()
}