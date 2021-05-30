package Daily

/*
6. Z 字形变换
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下:
P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
 */

func Lc530AnotherConvert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ans := ""
	for i := 0; i < len(s); i += 2 * numRows - 2 {
		ans += string(s[i])
	}
	l, r := 2 * numRows - 4, 2
	for i := 1; i < numRows - 1; i++ {
		for j := i; j < len(s); {
			ans += string(s[j])
			j += l
			if j >= len(s) {
				break
			}
			ans += string(s[j])
			j += r
		}
		l -= 2
		r += 2
	}
	for i := numRows - 1; i < len(s); i += 2 * numRows - 2 {
		ans += string(s[i])
	}
	return ans
}
