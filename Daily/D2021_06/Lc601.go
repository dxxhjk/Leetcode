package D2021_06

/*
1744. 你能在你最喜欢的那天吃到你最喜欢的糖果吗？
给你一个下标从 0 开始的正整数数组 candiesCount ，其中 candiesCount[i] 表示你拥有的第 i 类糖果的数目。
同时给你一个二维数组 queries，其中 queries[i] = [favoriteTypei, favoriteDayi, dailyCapi]。
你按照如下规则进行一场游戏：
你从第 0 天开始吃糖果。
你在吃完 所有第 i - 1 类糖果之前，不能吃任何一颗第 i 类糖果。
在吃完所有糖果之前，你必须每天至少吃一颗糖果。
请你构建一个布尔型数组 answer，满足 answer.length == queries.length 。
answer[i] 为 true 的条件是：在每天吃 不超过 dailyCapi 颗糖果的前提下，
你可以在第 favoriteDayi 天吃到第 favoriteTypei 类糖果；否则 answer[i] 为 false 。
注意，只要满足上面 3 条规则中的第二条规则，你就可以在同一天吃不同类型的糖果。
请你返回得到的数组 answer。
 */

/*
思路：
前缀和，算一下最爱糖的编号范围，是否落在最少糖到最多糖区间内
 */

func Lc601CanEat(candiesCount []int, queries [][]int) []bool {
	sum := []int{0}
	for _, c := range candiesCount {
		sum = append(sum, c + sum[len(sum) - 1])
	}
	ans := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		min, max := queries[i][1] + 1, (queries[i][1] + 1) * queries[i][2]
		candyMin, candyMax := sum[queries[i][0]] + 1, sum[queries[i][0] + 1]
		if max < candyMin || min > candyMax {
			ans[i] = false
		} else {
			ans[i] = true
		}
	}
	return ans
}