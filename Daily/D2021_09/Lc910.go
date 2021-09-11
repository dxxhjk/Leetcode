package D2021_09

/*
1894. 找到需要补充粉笔的学生编号
一个班级里有 n 个学生，编号为 0 到 n - 1。
每个学生会依次回答问题，编号为 0 的学生先回答，然后是编号为 1 的学生，
以此类推，直到编号为 n - 1 的学生，然后老师会重复这个过程，重新从编号为 0 的学生开始回答问题。

给你一个长度为 n 且下标从 0 开始的整数数组 chalk 和一个整数 k。
一开始粉笔盒里总共有 k 支粉笔。
当编号为 i 的学生回答问题时，他会消耗 chalk[i] 支粉笔。
如果剩余粉笔数量严格小于 chalk[i] ，那么学生 i 需要补充粉笔。
请你返回需要补充粉笔的学生编号。
 */

func Lc910ChalkReplacer(chalk []int, k int) int {
	sums := make([]int, 0)
	sums = append(sums, 0)
	sum := 0
	for _, c := range chalk {
		sum += c
		sums = append(sums, sum)
	}

	k = k % sum
	l, r := 0, len(chalk)
	mid := (l + r) / 2
	for !(sums[mid] <= k && sums[mid + 1] > k) {
		mid = (l + r) / 2
		if sums[mid + 1] < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return mid
}
