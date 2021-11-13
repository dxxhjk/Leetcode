package D2021_11

/*
我们定义，在以下情况时，单词的大写用法是正确的：

全部字母都是大写，比如 "USA" 。
单词中所有字母都不是大写，比如 "leetcode" 。
如果单词不只含有一个字母，只有首字母大写，比如 "Google" 。
给你一个字符串 word 。如果大写用法正确，返回 true ；否则，返回 false 。
 */

func Lc1113DetectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	flag := word[0] < 'a'
	if flag {
		f := word[1] < 'a'
		for i := 1; i < len(word); i++ {
			if f != (word[i] < 'a') {
				return false
			}
		}
	} else {
		for _, c := range word {
			if c < 'a' {
				return false
			}
		}
	}
	return true
}
