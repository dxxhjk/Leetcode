package D2021_11

import "fmt"

/*
391. 完美矩形
给你一个数组 rectangles ，其中 rectangles[i] = [xi, yi, ai, bi] 表示一个坐标轴平行的矩形。
这个矩形的左下顶点是 (xi, yi) ，右上顶点是 (ai, bi) 。
如果所有矩形一起精确覆盖了某个矩形区域，则返回 true ；否则，返回 false 。
 */

func Lc1116IsRectangleCover(rectangles [][]int) bool {
	points := make(map[[2]int]int, 0)
	totalArea := 0
	xMin, yMin, xMax, yMax := rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]
	for _, rectangle := range rectangles {
		points[[2]int{rectangle[0], rectangle[1]}]++
		if points[[2]int{rectangle[0], rectangle[1]}] == 2 {
			delete(points, [2]int{rectangle[0], rectangle[1]})
		}
		points[[2]int{rectangle[2], rectangle[3]}]++
		if points[[2]int{rectangle[2], rectangle[3]}] == 2 {
			delete(points, [2]int{rectangle[2], rectangle[3]})
		}
		points[[2]int{rectangle[0], rectangle[3]}]++
		if points[[2]int{rectangle[0], rectangle[3]}] == 2 {
			delete(points, [2]int{rectangle[0], rectangle[3]})
		}
		points[[2]int{rectangle[2], rectangle[1]}]++
		if points[[2]int{rectangle[2], rectangle[1]}] == 2 {
			delete(points, [2]int{rectangle[2], rectangle[1]})
		}
		totalArea += (rectangle[2] - rectangle[0]) * (rectangle[3] - rectangle[1])
		xMin = min(xMin, rectangle[0])
		yMin = min(yMin, rectangle[1])
		xMax = max(xMax, rectangle[2])
		yMax = max(yMax, rectangle[3])
	}
	fmt.Println(points, xMin, yMin, xMax, yMax)
	if len(points) == 4 && totalArea == (yMax - yMin) * (xMax - xMin) &&
		points[[2]int{xMin, yMin}] != 0 && points[[2]int{xMax, yMax}] != 0{
		return true
	}
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
