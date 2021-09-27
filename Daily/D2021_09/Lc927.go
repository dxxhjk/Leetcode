package D2021_09

import "fmt"

/*
639. 解码方法 II
一条包含字母 A-Z 的消息通过以下的方式进行了编码：
'A' -> 1
'B' -> 2
...
'Z' -> 26
要 解码 一条已编码的消息，所有的数字都必须分组，然后按原来的编码方案反向映射回字母（可能存在多种方式）。例如，"11106" 可以映射为：
"AAJF" 对应分组 (1 1 10 6)
"KJF" 对应分组 (11 10 6)
注意，像 (1 11 06) 这样的分组是无效的，因为 "06" 不可以映射为 'F' ，因为 "6" 与 "06" 不同。
除了 上面描述的数字字母映射方案，编码消息中可能包含 '*' 字符，可以表示从 '1' 到 '9' 的任一数字（不包括 '0'）。
例如，编码字符串 "1*" 可以表示 "11"、"12"、"13"、"14"、"15"、"16"、"17"、"18" 或 "19" 中的任意一条消息。
对 "1*" 进行解码，相当于解码该字符串可以表示的任何编码消息。
给你一个字符串 s ，由数字和 '*' 字符组成，返回 解码 该字符串的方法 数目 。
由于答案数目可能非常大，返回对 10^9 + 7 取余 的结果。
*/

func Lc927NumDecodings(s string) int {
	if len(s) == 1 {
		if s[0] == '*' {
			return 9
		}
		return 1
	}
	a, b := 1, Lc927NumDecodings(s[:1])
	for i := 1; i < len(s); i++ {
		fmt.Println(a, b)
		oneCh, twoCh := 0, 0
		if s[i] == '*' {
			oneCh = 9
			if s[i-1] == '1' {
				twoCh = 9
			} else if s[i-1] == '2' {
				twoCh = 6
			} else if s[i-1] == '*' {
				twoCh = 15
			}
		} else {
			if s[i] != '0' {
				oneCh = 1
			}
			if s[i-1] == '1' {
				twoCh = 1
			} else if s[i-1] == '2' {
				if s[i] > '6' {
					twoCh = 0
				} else {
					twoCh = 1
				}
			} else if s[i-1] == '*' {
				if s[i] > '6' {
					twoCh = 1
				} else {
					twoCh = 2
				}
			}
		}
		fmt.Println(oneCh, twoCh)
		a, b = b, (a*twoCh+b*oneCh)%(1000000007)
	}
	return b
}
