package D2021_11

/*
438. 找到字符串中所有字母异位词
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。
不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
 */

func Lc1128FindAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	var diffNums [26]int
	for i, c := range p {
		diffNums[s[i] - 'a']++
		diffNums[c - 'a']--
	}
	diff := 0
	ans := make([]int, 0)
	for _, c := range diffNums {
		if c != 0 {
			diff++
		}
	}
	if diff == 0 {
		ans = append(ans, 0)
	}
	for i, c := range s[:len(s) - len(p)] {
		if diffNums[c - 'a'] == 1 {
			diff--
		} else if diffNums[c - 'a'] == 0 {
			diff++
		}
		diffNums[c - 'a']--
		if diffNums[s[i + len(p)] - 'a'] == -1 {
			diff--
		} else if diffNums[s[i + len(p)] - 'a'] == 0 {
			diff++
		}
		diffNums[s[i + len(p)] - 'a']++
		if diff == 0 {
			ans = append(ans, i + 1)
		}
	}
	return ans
}
