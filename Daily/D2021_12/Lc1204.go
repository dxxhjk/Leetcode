package D2021_12

/*
383. 赎金信
为了不在赎金信中暴露字迹，从杂志上搜索各个需要的字母，组成单词来表达意思。
给你一个赎金信 (ransomNote) 字符串和一个杂志(magazine)字符串，
判断 ransomNote 能不能由 magazines 里面的字符构成。
如果可以构成，返回 true ；否则返回 false 。
magazine 中的每个字符只能在 ransomNote 中使用一次。
 */

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int, 0)
	for _, r := range ransomNote {
		m[r]++
	}
	for _, r := range magazine {
		m[r]--
	}
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}
