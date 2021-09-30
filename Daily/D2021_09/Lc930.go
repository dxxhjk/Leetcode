package D2021_09

/*
223. 矩形面积
给你 二维 平面上两个 由直线构成的 矩形，请你计算并返回两个矩形覆盖的总面积。
每个矩形由其 左下 顶点和 右上 顶点坐标表示：
第一个矩形由其左下顶点 (ax1, ay1) 和右上顶点 (ax2, ay2) 定义。
第二个矩形由其左下顶点 (bx1, by1) 和右上顶点 (bx2, by2) 定义。
*/

func Lc930ComputeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	x, y := getCoincideLength(ax1, ax2, bx1, bx2), getCoincideLength(ay1, ay2, by1, by2)
	return (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1) - x*y

}

func getCoincideLength(l1l, l1r, l2l, l2r int) int {
	if l1r <= l2l {
		return 0
	} else if l1l <= l2l && l1r >= l2l && l1r <= l2r {
		return l1r - l2l
	} else if l1l >= l2l && l1r <= l2r {
		return l1r - l1l
	} else if l1l >= l2l && l1l <= l2r && l1r >= l2r {
		return l2r - l1l
	} else if l1l >= l2r {
		return 0
	} else {
		return l2r - l2l
	}
}
