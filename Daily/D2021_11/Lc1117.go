package D2021_11

/*
给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，
并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。
如果不存在这样的两个单词，返回 0。
 */

func Lc1117MaxProduct(words []string) int {
	m := make(map[int]int, 0)
	for _, word := range words{
		bit := 0
		for _, c := range word {
			bit = bit | (1 << (c - 'a'))
		}
		m[bit] = max(m[bit], len(word))
	}
	if len(m) < 2 {
		return 0
	}
	mArr := make([]int, 0)
	for k := range m {
		mArr = append(mArr, k)
	}
	ans := 0
	for i := 0; i < len(mArr); i++ {
		for j := i + 1; j < len(mArr); j++ {
			if mArr[i] & mArr[j] == 0 {
				ans = max(ans, m[mArr[i]] * m[mArr[j]])
			}
		}
	}
	return ans
}
