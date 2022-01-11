package D2021_12

func Lc1224EatenApples(apples []int, days []int) int {
	ans := 0
	hasApple := 0
	for i := 0; i < len(apples); i++ {
		if hasApple != 0 {
			hasApple--
			ans++
			if days[i] > hasApple {
				hasApple += min(apples[i], days[i] - hasApple)
			}
		} else {
			if apples[i] == 0 {
				continue
			}
			ans++
			hasApple = min(apples[i] - 1, days[i])
		}
	}
	return ans + hasApple
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}