package D2021_11

/*
分糖果
给定一个偶数长度的数组，其中不同的数字代表着不同种类的糖果，每一个数字代表一个糖果。
你需要把这些糖果平均分给一个弟弟和一个妹妹。返回妹妹可以获得的最大糖果的种类数。
 */

func Lc1101DistributeCandies(candyType []int) int {
	tNum := 0
	m := make(map[int]int, 0)
	for _, t := range candyType {
		if m[t] == 0 {
			tNum++
			m[t]++
		}
	}
	if len(candyType) / 2 > tNum {
		return tNum
	}
	return len(candyType) / 2
}
