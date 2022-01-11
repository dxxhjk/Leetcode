package D2022_01

/*
在一个 106 x 106 的网格中，每个网格上方格的坐标为 (x, y) 。
现在从源方格 source = [sx, sy] 开始出发，意图赶往目标方格 target = [tx, ty]。
数组 blocked 是封锁的方格列表，其中每个 blocked[i] = [xi, yi] 表示坐标为 (xi, yi) 的方格是禁止通行的。
每次移动，都可以走到网格中在四个方向上相邻的方格，只要该方格不在给出的封锁列表 blocked 上。
同时，不允许走出网格。
只有在可以通过一系列的移动从源方格 source 到达目标方格 target 时才返回 true.
否则，返回 false。
 */

func Lc0111IsEscapePossible(blocked [][]int, source []int, target []int) bool {
	maxSize := len(blocked) * (len(blocked) - 1) / 2
	bm := make(map[int]int)
	for _, pos := range blocked {
		bm[pos[0] * 1e6 + pos[1]] = 1
	}
	return bfsFunc(bm, source, maxSize) && bfsFunc(bm, target, maxSize)

}

func bfsFunc(bm map[int]int, source []int, maxSize int) bool {
	bfs := make([]int, 0)
	bfs = append(bfs, source[0] * 1e6 + source[1])
	vis := bm
	for len(vis) - len(bm) < maxSize && len(bfs) != 0 {
		tmp := bfs[0]
		bfs = bfs[1:]
		if vis[tmp] == 1 {
			continue
		}
		vis[tmp] = 1
		//fmt.Println(tmp, bfs,vis)
		if tmp + 1e6 < 1e12 && vis[tmp + 1e6] == 0 {
			bfs = append(bfs, tmp + 1e6)
		}
		if tmp - 1e6 >= 0 && vis[tmp - 1e6] == 0 {
			bfs = append(bfs, tmp - 1e6)
		}
		if tmp % 1e6 + 1 < 1e6 && vis[tmp + 1] == 0 {
			bfs = append(bfs, tmp + 1)
		}
		if tmp % 1e6 - 1 >= 0 && vis[tmp - 1] == 0 {
			bfs = append(bfs, tmp - 1)
		}
	}
	return len(vis) - len(bm) >= maxSize
}