package D2021_09

func Lc907BalancedStringSplit(s string) int {
	if len(s) == 0 {
		return 0
	}
	l := 0
	ans := 0
	for _, i := range s {
		if i == 'L' {
			l++
		} else {
			l--
		}
		if l == 0 {
			ans++
		}
	}
	return ans
}