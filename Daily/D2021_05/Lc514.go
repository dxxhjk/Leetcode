package D2021_05

/*
12. 整数转罗马数字
罗马数字包含以下七种字符: I，V，X，L，C，D 和 M。
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如，罗马数字 2 写做 II ，即为两个并列的 1。
12 写做XII，即为 X + II 。
27 写做 XXVII, 即为 XX + V + II 。
通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做IIII，而是IV。
数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。
同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给你一个整数，将其转为罗马数字。
 */

func Lc514IntToRoman(num int) string {
	thousand, hundred, ten, one := num / 1000, (num % 1000) / 100, (num % 100) / 10, num % 10
	ans := ""
	for thousand != 0 {
		ans += "M"
		thousand--
	}
	if hundred == 9 {
		ans += "CM"
		hundred -= 9
	}
	if hundred >= 5 {
		ans += "D"
		hundred -= 5
	}
	if hundred == 4 {
		ans += "CD"
		hundred -= 4
	}
	for hundred != 0 {
		ans += "C"
		hundred--
	}
	if ten == 9 {
		ans += "XC"
		ten -= 9
	}
	if ten >= 5 {
		ans += "L"
		ten -= 5
	}
	if ten == 4 {
		ans += "XL"
		ten -= 4
	}
	for ten != 0 {
		ans += "X"
		ten--
	}
	if one == 9 {
		ans += "IX"
		one -= 9
	}
	if one >= 5 {
		ans += "V"
		one -= 5
	}
	if one == 4 {
		ans += "IV"
		one -= 4
	}
	for one != 0 {
		ans += "I"
		one--
	}
	return ans
}
