package main

import (
	"Leetcode/Daily/D2022_01"
	"fmt"
)

func main() {
	ans := D2022_01.Lc0112IncreasingTriplet([]int{2,4,-2,-3})
	fmt.Println(ans, 0x7fffffff)
}

func isAdditiveNumber(num string) bool {
	for i := 1; i <= 17 && i < len(num); i++ {
		for j := 1; j <= 18 - i && i + j < len(num); j++ {
			if isAdditive(i, j, num, true) {
				return true
			}
		}
	}
	return false
}

func isAdditive(l1, l2 int, num string, first bool) bool {
	fmt.Println(l1, l2, num)
	n1 := num[:l1]
	n2 := num[l1:l1 + l2]
	if l1 > 1 && n1[0] == '0' || l2 > 1 && n2[0] == '0' {
		return false
	}
	left := num[l1 + l2:]
	if len(left) == 0 {
		return !first
	}
	n3 := plus(n1, n2)
	l3 := len(n3)
	if l3 > len(left) {
		return false
	}
	if left[:l3] == n3 {
		return isAdditive(l2, l3, n2 + left, false)
	}
	return false
}

func plus(a, b string) string {
	aa, bb := make([]uint8, 0), make([]uint8, 0)
	for i := len(a) - 1; i >= 0; i-- {
		aa = append(aa, a[i] - '0')
	}
	for i := len(b) - 1; i >= 0; i-- {
		bb = append(bb, b[i] - '0')
	}
	length, inc := 0, uint8(0)
	ans := make([]uint8, 0)
	if len(b) > len(a) {
		length = len(b) + 1
	} else {
		length = len(a) + 1
	}
	for i := 0; i < length; i++ {
		tmp := inc
		if len(a) > i {
			tmp += aa[i]
		}
		if len(b) > i {
			tmp += bb[i]
		}
		if tmp >= 10 {
			inc = 1
		} else {
			inc = 0
		}
		ans = append(ans, tmp % 10)
	}
	if ans[len(ans) - 1] == 0 {
		ans = ans[:len(ans) - 1]
	}
	ansStr := ""
	for i := len(ans) - 1; i >= 0; i-- {
		ansStr += string(rune(ans[i] + '0'))
	}
	return ansStr
}
