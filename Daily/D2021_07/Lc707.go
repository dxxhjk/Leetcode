package D2021_07

/*
1711. 大餐计数
大餐是指恰好包含两道不同餐品的一餐，其美味程度之和等于 2 的幂。
你可以搭配任意两道餐品做一顿大餐。
给你一个整数数组 deliciousness ，其中 deliciousness[i] 是第 i 道餐品的美味程度，返回你可以用数组中的餐品做出的不同大餐的数量。
结果需要对 10^9 + 7 取余。注意，只要餐品下标不同，就可以认为是不同的餐品，即便它们的美味程度相同。
 */

/*
思路：
哈希表
 */

func Lc707CountPairs(deliciousness []int) int {
	m := make(map[int]int)
	ans := 0
	for _, d := range deliciousness {
		temp := 1
		for i := 0; i < 22; i++ {
			if temp >= d {
				v, ok := m[temp - d]
				if ok {
					ans += v
				}
			}
			temp <<= 1
		}
		m[d]++
	}
	return ans % (1e9 + 7)
}
