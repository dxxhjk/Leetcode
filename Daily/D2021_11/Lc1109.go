package D2021_11

/*
你正在参与祖玛游戏的一个变种。
在这个祖玛游戏变体中，桌面上有 一排 彩球，每个球的颜色可能是：红色 'R'、黄色 'Y'、蓝色 'B'、绿色 'G' 或白色 'W' 。你的手中也有一些彩球。
你的目标是 清空 桌面上所有的球。每一回合：
	从你手上的彩球中选出 任意一颗 ，然后将其插入桌面上那一排球中：两球之间或这一排球的任一端。
	接着，如果有出现 三个或者三个以上 且 颜色相同 的球相连的话，就把它们移除掉。
	如果这种移除操作同样导致出现三个或者三个以上且颜色相同的球相连，则可以继续移除这些球，直到不再满足移除条件。
	如果桌面上所有球都被移除，则认为你赢得本场游戏。
	重复这个过程，直到你赢了游戏或者手中没有更多的球。
给你一个字符串 board ，表示桌面上最开始的那排球。
另给你一个字符串 hand ，表示手里的彩球。
请你按上述操作步骤移除掉桌上所有球，计算并返回所需的 最少 球数。
如果不能移除桌上所有的球，返回 -1 。
 */

type state struct {
	board string
	hand [5]int
}

func Lc1109FindMinStep(board string, hand string) int {
	cache := map[string]string{}
	COLORS := "RYBGW"

	var clean func(b string) string
	clean = func(board string) string {
		if v, ok := cache[board]; ok {
			return v
		}
		res := board
		for i, j := 0, 0; i < len(board); {
			for j < len(board) && board[i] == board[j] {
				j += 1
			}
			if j - i > 2 {
				res = clean(board[:i] + board[j:])
				cache[board] = res
				return res
			}
			i = j
		}
		cache[board] = res
		return res
	}

	cnts := func(hand string) [5]int {
		res := [5]int{}
		for i := 0; i < len(hand); i++ {
			for j, c := range COLORS {
				if hand[i] == byte(c) {
					res[j]++
					break
				}
			}
		}
		return res
	}

	queue := make([]state, 0, 6)
	init := state{board, cnts(hand)}
	queue = append(queue, init)
	visited := map[state]int{}
	visited[init] = 0
	for len(queue) > 0 {
		curState := queue[0]
		cur_board, cur_hand := curState.board, curState.hand
		if len(cur_board) == 0 {
			return visited[curState]
		}
		queue = queue[1:]
		for i := 0; i <= len(cur_board) ; i++ {
			for j, r := range COLORS {
				if cur_hand[j] > 0 {
					c := byte(r)
					// 第 1 个剪枝条件: 只在连续相同颜色的球的开头位置插入新球(在它前面插入过了，不需要再插入，意义相同)
					if i > 0 && cur_board[i - 1] == c{
						continue
					}

					/**
					 *  第 2 个剪枝条件: 只在以下两种情况放置新球
					 *  - 第 1 种情况 : 当前后颜色相同且与当前颜色不同时候放置球
					 *  - 第 2 种情况 : 当前球颜色与后面的球的颜色相同
					 */
					choose := false
					if 0 < i && i < len(cur_board) && cur_board[i - 1] == cur_board[i] && cur_board[i - 1] != c{
						choose = true
					}
					if i < len(cur_board) && cur_board[i] == c{
						choose = true
					}

					if choose {
						nxt := [5]int{}
						for k,_ := range COLORS{
							nxt[k] = cur_hand[k]
						}
						nxt[j] -= 1

						nextState := state{clean(cur_board[:i] + string(c) + cur_board[i:]), nxt}
						if _,ok := visited[nextState]; !ok {
							queue = append(queue, nextState)
							visited[nextState] = visited[curState] + 1
						}
					}
				}
			}
		}
	}
	return -1
}
