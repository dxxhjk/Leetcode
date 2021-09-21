package D2021_09

/*
58. 最后一个单词的长度
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。
返回字符串中最后一个单词的长度。
单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
*/

func Lc921LengthOfLastWord(s string) int {
	i := len(s) - 1
	for ; i >= 0; i-- {
		if s[i] != ' ' {
			break
		}
	}
	length := 0
	for ; i >= 0; i-- {
		if s[i] != ' ' {
			length++
		} else {
			break
		}
	}
	return length
}
