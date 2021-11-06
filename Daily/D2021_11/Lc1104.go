package D2021_11

/*
有效的完全平方数
给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
进阶：不要 使用任何内置的库函数，如 sqrt 。
 */

func Lc1104IsPerfectSquare(num int) bool {
	l, r := 1, num
	mid := (l + r) / 2
	for mid * mid != num {
		if mid * mid > num {
			if r == mid {
				return false
			}
			r = mid
		} else {
			if l == mid {
				return false
			}
			l = mid
		}
		mid = (l + r) / 2
	}
	return true
}
