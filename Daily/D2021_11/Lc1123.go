package D2021_11

/*
859. 亲密字符串
给你两个字符串 s 和 goal ，只要我们可以通过交换 s 中的两个字母得到与 goal 相等的结果，
就返回 true ；否则返回 false 。
 */

func Lc1123buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	flag := false
	a, b, c, d := byte('a' - 1), byte('a' - 1), byte('a' - 1), byte('a' - 1)
	for i := range s {
		if !flag && s[i] != goal[i] {
			flag = true
			a = s[i]
			b = goal[i]
		} else if s[i] != goal[i] {
			if c != byte('a' - 1) {
				return false
			}
			c = s[i]
			d = goal[i]
		}
	}
	if flag {
		return a == d && b == c
	}
	m := make(map[rune]int)
	for _, r := range s {
		if m[r] == 1 {
			return true
		}
		m[r]++
	}
	return false
}
