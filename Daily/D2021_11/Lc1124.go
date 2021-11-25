package D2021_11

import "strconv"

/*
423. 从英文中重建数字
给你一个字符串 s ，其中包含字母顺序打乱的用英文单词表示的若干数字（0-9）。
按 升序 返回原始的数字。
*/

func Lc1124OriginalDigits(s string) string {
	m := make(map[rune]int, 0)
	for _, r := range s {
		m[r]++
	}
	ans := make([]int, 10)
	ans[0] = m['z']
	ans[2] = m['w']
	ans[4] = m['u']
	ans[6] = m['x']
	ans[8] = m['g']
	ans[1] = m['o'] - ans[0] - ans[2] - ans[4]
	ans[3] = m['r'] - ans[0] - ans[4]
	ans[7] = m['s'] - ans[6]
	ans[5] = m['v'] - ans[7]
	ans[9] = m['i'] - ans[8] - ans[6] - ans[5]
	ret := ""
	for i := 0; i < 10; i++ {
		for j := 0; j < ans[i]; j++ {
			ret += strconv.Itoa(i)
		}
	}
	return ret
}
