package D2021_05

/*
1074. 元素和为目标值的子矩阵数量
给出矩阵 matrix 和目标值 target，返回元素总和等于目标值的非空子矩阵的数量。
子矩阵 x1, y1, x2, y2 是满足 x1 <= x <= x2 且 y1 <= y <= y2 的所有单元 matrix[x][y] 的集合。
如果 (x1, y1, x2, y2) 和 (x1', y1', x2', y2') 两个子矩阵中部分坐标不同（如：x1 != x1'），那么这两个子矩阵也不同。
 */

/*
思路：
遍历子数组的上下边界，对每一对边界 i, j 内的每一列求和，
将问题转换成一维数组的子数组固定和问题，用前缀和加哈希表解决。
子问题见Lc529Sub
 */

func Lc529NumSubMatrixSumTarget(matrix [][]int, target int) int {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		temp := make([]int, len(matrix[0]))
		for j := i; j < len(matrix); j++ {
			pre := 0
			h := make(map[int]int)
			h[0] = 1
			for k := 0; k < len(matrix[0]); k++ {
				temp[k] += matrix[j][k]
				pre += temp[k]
				sum += h[pre - target]
				h[pre]++
			}
		}
	}
	return sum
}

