package D2021_09

/*
请你判断一个 9x9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
数独部分空格内已填入了数字，空白格用 '.' 表示。
*/

func Lc917IsValidSudoku(board [][]byte) bool {
	map1, map2, map3 := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			ch := board[i][j]
			if ch != '.' {
				ch -= '1'
				if map1[i][ch] {
					return false
				} else {
					map1[i][ch] = true
				}
				if map2[j][ch] {
					return false
				} else {
					map2[j][ch] = true
				}
				if map3[(i/3)*3+j/3][ch] {
					return false
				} else {
					map3[(i/3)*3+j/3][ch] = true
				}
			}
		}
	}
	return true
}
