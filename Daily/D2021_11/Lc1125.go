package D2021_11

/*
458. 可怜的小猪
有 buckets 桶液体，其中 正好 有一桶含有毒药，其余装的都是水。它们从外观看起来都一样。
为了弄清楚哪只水桶含有毒药，你可以喂一些猪喝，通过观察猪是否会死进行判断。
不幸的是，你只有 minutesToTest 分钟时间来确定哪桶液体是有毒的。
喂猪的规则如下：
	1. 选择若干活猪进行喂养
	2. 可以允许小猪同时饮用任意数量的桶中的水，并且该过程不需要时间。
	3. 小猪喝完水后，必须有 minutesToDie 分钟的冷却时间。在这段时间里，你只能观察，而不允许继续喂猪。
	4. 过了 minutesToDie 分钟后，所有喝到毒药的猪都会死去，其他所有猪都会活下来。
	5. 重复这一过程，直到时间用完。
给你桶的数目 buckets ，minutesToDie 和 minutesToTest ，返回在规定时间内判断哪个桶有毒所需的 最小 猪数。
 */

func Lc1125PoorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	times := minutesToTest / minutesToDie
	dpf := make([][]int, buckets)
	for i := range dpf {
		dpf[i] = make([]int, times)
		dpf[i][0] = 1
	}
	for i := range dpf[0] {
		dpf[0][i] = 1
	}
	return 0
}

func getC(a, b int) int {
	if b > a / 2 {
		b = a - b
	}
	ans := 1
	for i := 0; i < b; i++ {
		ans *= a - i
	}
	for i := 0; i < b; i++ {
		ans /= i + 1
	}
	return ans
}
