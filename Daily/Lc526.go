package Daily

import "Leetcode/Structs"

/*
1190. 反转每对括号间的子串
给出一个字符串 s（仅含有小写英文字母和括号）。
请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。
注意，您的结果中不应包含任何括号。
 */

/*
思路：
栈模拟，用栈存储便利过的除‘）’以外的字符，遇到则将栈中‘（’之后的字符数顺序倒置，需要用到一个额外数组
 */

func Lc526ReverseParentheses(s string) string {
	stack := Structs.NewStack()
	for i := 0; i < len(s); i++ {
		if s[i] != ')' {
			stack.Push(s[i])
		} else {
			temp := make([]byte, 0)
			_, tempB := stack.Pop()
			for tempB.(byte) != '(' {
				temp = append(temp, tempB.(byte))
				_, tempB = stack.Pop()
			}
			for _, i := range temp {
				stack.Push(i)
			}
		}
	}
	ans := ""
	flag, elem := stack.Pop()
	for flag {
		ans = string(elem.(byte)) + ans
		flag, elem = stack.Pop()
	}
	return ans
}