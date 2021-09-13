package D2021_09

/*
447. 回旋镖的数量
给定平面上 n 对 互不相同 的点 points ，其中 points[i] = [xi, yi] 。
回旋镖 是由点 (i, j, k) 表示的元组 ，其中 i 和 j 之间的距离和 i 和 k 之间的距离相等（需要考虑元组的顺序）。
返回平面上所有回旋镖的数量。
 */

func Lc913NumberOfBoomerangs(points [][]int) int {
	ans := 0
	m := make(map[int]int)
	for _, p := range points {
		for _, ap := range points {
			m[getDis(p[0], ap[0], p[1], ap[1])]++
		}
		ans += getBoomerangs(m)
		m = make(map[int]int)
	}
	return ans * 2
}

func getDis(x1, x2, y1, y2 int) int {
	return (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2)
}

func getBoomerangs(m map[int]int) int {
	num := 0
	for _, n := range m {
		num += n * (n - 1) / 2
	}
	return num
}