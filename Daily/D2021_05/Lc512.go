package D2021_05

/*
1310. 子数组异或查询
有一个正整数数组 arr，现给你一个对应的查询数组 queries，其中 queries[i] = [Li, Ri]。
对于每个查询 i，请你计算从 Li 到 Ri 的 XOR 值（即 arr[Li] xor arr[Li+1] xor ... xor arr[Ri]）作为本次查询的结果。
并返回一个包含给定查询queries所有结果的数组。
 */

/*
思路：
一遍 arr 记录前缀到 map，一遍 queries，直接用前缀做异或
时间O(n)
空间O(n)
 */

func Lc512XorQueries(arr []int, queries [][]int) []int {
	xors := make([]int, 0)
	xors = append(xors, 0)
	for i := 0; i < len(arr); i++ {
		xor := arr[i] ^ xors[i]
		xors = append(xors, xor)
	}
	ans := make([]int, 0)
	for _, query := range queries {
		ans = append(ans, xors[query[0]] ^ xors[query[1] + 1])
	}
	return ans
}