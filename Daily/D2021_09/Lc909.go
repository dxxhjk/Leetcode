package D2021_09

func Lc909FullJustify(words []string, maxWidth int) []string {
	ans := make([]string, 0)
	buf := make([]string, 0)
	i, tempLength := 0, 0
	for i < len(words) {
		buf = buf[:0]
		tempLength = 0
		for i < len(words) && len(words[i]) <= maxWidth - tempLength {
			buf = append(buf, words[i])
			tempLength += len(words[i]) + 1
			i++
		}
		ans = append(ans, makeLine(buf, tempLength - 1, maxWidth))
	}
	ans[len(ans) - 1] = makeLastLine(buf, maxWidth)
	return ans
}

func makeLine(buf []string, length, maxWidth int) string {
	ret := ""
	ret += buf[0]
	if len(buf) == 1 {
		for len(ret) < maxWidth {
			ret += " "
		}
		return ret
	}
	j := 1
	for i := 0; i < (maxWidth - length) % (len(buf) - 1); i++ {
		for k := 0; k < (maxWidth - length) / (len(buf) - 1) + 2; k++ {
			ret += " "
		}
		ret += buf[j]
		j++
	}
	for i := (maxWidth - length) % (len(buf) - 1); i < len(buf) - 1; i++ {
		for k := 0; k < (maxWidth - length) / (len(buf) - 1) + 1; k++ {
			ret += " "
		}
		ret += buf[j]
		j++
	}
	return ret
}

func makeLastLine(buf []string, maxWidth int) string {
	ret := buf[0]
	for i := 1; i < len(buf); i++ {
		ret += " " + buf[i]
	}
	for len(ret) < maxWidth {
		ret += " "
	}
	return ret
}
