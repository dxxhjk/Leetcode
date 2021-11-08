package D2021_11

import "strconv"

/*
你在和朋友一起玩 猜数字（Bulls and Cows）游戏，该游戏规则如下：
写出一个秘密数字，并请朋友猜这个数字是多少。朋友每猜测一次，你就会给他一个包含下述信息的提示：
猜测数字中有多少位属于数字和确切位置都猜对了（称为 "Bulls", 公牛），
有多少位属于数字猜对了但是位置不对（称为 "Cows", 奶牛）。
也就是说，这次猜测中有多少位非公牛数字可以通过重新排列转换成公牛数字。
给你一个秘密数字 secret 和朋友猜测的数字 guess ，请你返回对朋友这次猜测的提示。
提示的格式为 "xAyB" ，x 是公牛个数， y 是奶牛个数，A 表示公牛，B 表示奶牛。
 */

func Lc1108GetHint(secret string, guess string) string {
	a, b := 0, 0
	ma, mb := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			a++
		} else {
			ma[secret[i]]++
			mb[guess[i]]++
		}
	}
	for k, v := range ma {
		b += getMin(mb[k], v)
	}
	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
