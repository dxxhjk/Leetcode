package D2021_09

import "Leetcode/Structs"

/*
给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words，找出所有同时在二维网格和字典中出现的单词。
单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母在一个单词中不允许被重复使用。
*/

func Lc916FindWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 {
		return []string{}
	}
	trie := Structs.MakeTrie()
	for _, word := range words {
		trie.Insert(word)
	}
	ansMap := make(map[string]int)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, len(board[0])-1, len(board)-1, 0, "", board, trie, ansMap)
		}
	}
	ans := make([]string, 0)
	for k, _ := range ansMap {
		ans = append(ans, k)
	}
	return ans
}

func dfs(x, y, w, h, i int, temp string, board [][]byte, trie Structs.Trie, ans map[string]int) {
	if board[x][y] == '#' {
		return
	}
	temp = temp + string(board[x][y])
	if !trie.StartsWith(temp) {
		return
	}
	if trie.Search(temp) {
		ans[temp]++
	}
	tempCh := board[x][y]
	board[x][y] = '#'
	if x > 0 {
		dfs(x-1, y, w, h, i+1, temp, board, trie, ans)
	}
	if x < h {
		dfs(x+1, y, w, h, i+1, temp, board, trie, ans)
	}
	if y > 0 {
		dfs(x, y-1, w, h, i+1, temp, board, trie, ans)
	}
	if y < w {
		dfs(x, y+1, w, h, i+1, temp, board, trie, ans)
	}
	board[x][y] = tempCh
	return
}
