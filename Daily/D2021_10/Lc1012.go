package D2021_10

import "math"

/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
返回被除数 dividend 除以除数 divisor 得到的商。
整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
 */

func Lc1012Divide(dividend int, divisor int) int {
	if dividend==0 {return 0}
	if dividend==math.MinInt32 && divisor==-1 {return math.MaxInt32}

	diffSign:=false
	if (dividend<0)!=(divisor<0) {
		diffSign=true
	}

	i,j,sum := 0,0,0
	for {
		tmp:=0
		if diffSign {
			tmp=sum-divisor<<j
		}else{
			tmp=sum+divisor<<j
		}
		if (dividend>0 && tmp>dividend) || (dividend<0 && tmp<dividend) {
			if j==0 {break}
			j-- //步长减半
			continue
		}
		sum=tmp
		i+=1<<j
		j++ //步长加倍
	}

	if diffSign {
		return -i
	}
	return i
}
