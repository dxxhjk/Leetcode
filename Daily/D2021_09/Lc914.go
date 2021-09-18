package D2021_09

import (
	"sort"
)

/*
524. 通过删除字母匹配到字典里最长单词
给你一个字符串 s 和一个字符串数组 dictionary ，找出并返回 dictionary 中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。
如果答案不止一个，返回长度最长且字典序最小的字符串。如果答案不存在，则返回空字符串。
*/

func Lc914FindLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) > len(dictionary[j]) ||
			len(dictionary[i]) == len(dictionary[j]) && dictionary[i] > dictionary[j]
	})
	for _, word := range dictionary {
		i, j := 0, 0
		for i != len(s) && j != len(word) {
			if s[i] == word[j] {
				j++
			}
			i++
		}
		if j == len(word) {
			return word
		}
	}
	return ""
}
