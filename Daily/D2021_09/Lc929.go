package D2021_09

/*
517.超级洗衣机
假设有 n 台超级洗衣机放在同一排上。
开始的时候，每台洗衣机内可能有一定量的衣服，也可能是空的。
在每一步操作中，你可以选择任意 m (1 <= m <= n) 台洗衣机，
与此同时将每台洗衣机的一件衣服送到相邻的一台洗衣机。
给定一个整数数组 machines 代表从左至右每台洗衣机中的衣物数量，
请给出能让所有洗衣机中剩下的衣物的数量相等的最少的操作步数。
如果不能使每台洗衣机中衣物的数量相等，则返回 -1 。
*/

func Lc929FindMinMoves(machines []int) int {
	avg := 0
	for _, machine := range machines {
		avg += machine
	}
	if avg%len(machines) != 0 {
		return -1
	}
	avg /= len(machines)
	l, r := 0, machines[0]-avg
	num := max(0, r)
	ans := num
	for i := 1; i < len(machines); i++ {
		l, r = l+avg-machines[i-1], r-avg+machines[i]
		num = max(0, l) + max(0, r)
		ans = max(num, ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
