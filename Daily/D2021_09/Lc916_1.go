package D2021_09

/*
给定一个m x n 二维字符网格board 和一个字符串单词word 。
如果 word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母不允许被重复使用。
*/

func Lc916_1Exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs1(i, j, len(board[0])-1, len(board)-1, 0, board, word) {
				return true
			}
		}
	}
	return false
}

func dfs1(x, y, w, h, i int, board [][]byte, word string) bool {
	if i == len(word)-1 {
		return board[x][y] == word[i]
	}
	if board[x][y] != word[i] {
		return false
	}
	temp := board[x][y]
	board[x][y] = '#'
	result := false
	if x > 0 {
		result = result || dfs1(x-1, y, w, h, i+1, board, word)
	}
	if x < h {
		result = result || dfs1(x+1, y, w, h, i+1, board, word)
	}
	if y > 0 {
		result = result || dfs1(x, y-1, w, h, i+1, board, word)
	}
	if y < w {
		result = result || dfs1(x, y+1, w, h, i+1, board, word)
	}
	board[x][y] = temp
	return result
}
