package D2021_11

/*
接雨水
给你一个 m x n 的矩阵，其中的值均为非负整数，代表二维高度图每个单元的高度，
请计算图中形状最多能接多少体积的雨水。
 */

func Lc1103TrapRainWater(heightMap [][]int) int {
	flag := true
	volume := 0
	l, w, h := len(heightMap), len(heightMap[0]), 1
	for flag {
		flag = false
		layerMap := make([][]bool, l)
		for i := 0; i < l; i++ {
			layerMap[i] = make([]bool, w)
			for j := 0; j < w; j++ {
				if heightMap[i][j] >= h {
					layerMap[i][j] = true
					if heightMap[i][j] > h {
						flag = true
					}
				}
			}
		}
		volume += getLayerVolume(layerMap)
		h++
	}
	return volume
}

func getLayerVolume(layerMap [][]bool) int {
	for i := 1; i < len(layerMap[0]); i++ {
		if !layerMap[0][i] {
			dfs(layerMap, 0, i)
		}
	}
	for i := 1; i < len(layerMap); i++ {
		if !layerMap[i][len(layerMap[0]) - 1] {
			dfs(layerMap, i, len(layerMap[0]) - 1)
		}
	}
	for i := 1; i < len(layerMap[0]); i++ {
		if !layerMap[len(layerMap) - 1][len(layerMap[0]) - i - 1] {
			dfs(layerMap, len(layerMap) - 1, len(layerMap[0]) - i - 1)
		}
	}
	for i := 1; i < len(layerMap); i++ {
		if !layerMap[len(layerMap) - i - 1][0] {
			dfs(layerMap, len(layerMap) - i - 1, 0)
		}
	}
	num := 0
	for i := 0; i < len(layerMap); i++ {
		for j := 0; j < len(layerMap[0]); j++ {
			if !layerMap[i][j] {
				num++
			}
		}
	}
	return num
}

func dfs(layerMap [][]bool, x, y int) {
	layerMap[x][y] = true
	if y > 0 && !layerMap[x][y - 1] {
		dfs(layerMap, x, y - 1)
	}
	if x > 0 && !layerMap[x - 1][y] {
		dfs(layerMap, x - 1, y)
	}
	if y < len(layerMap[0]) - 1 && !layerMap[x][y + 1] {
		dfs(layerMap, x, y + 1)
	}
	if x < len(layerMap) - 1 && !layerMap[x + 1][y] {
		dfs(layerMap, x + 1, y)
	}
}
