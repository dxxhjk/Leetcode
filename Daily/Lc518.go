package Daily

/*
1442. 形成两个异或相等数组的三元组数目

给你一个整数数组 arr 。
现需要从数组中取三个下标 i、j 和 k ，其中 (0 <= i < j <= k < arr.length) 。
a 和 b 定义如下：
a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]
注意：^ 表示 按位异或 操作。
请返回能够令 a == b 成立的三元组 (i, j, k) 的数目。
 */

/*
思路：
因为若 a ^ c == b, 则 a == c ^ b，
因此 j 在 i 到 k 之间的哪个位置无关紧要，
题目转化为：找到区间内元素全部异或后为0的所有区间[i, k]，ans += k - i
即找 Sk - Si == 0
使用 Hash 进行存储：Sa == Sb == Sc == k, 则 map[k] = []int{a, b, c}
时间O(n), 空间O(n)
 */

func Lc518CountTriplets(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	sumMap := make(map[int][]int, 0)
	tempSum := 0
	sumMap[0] = []int{0}
	for i := 0; i < len(arr); i++ {
		tempSum ^= arr[i]
		_, ok := sumMap[tempSum]
		if ok {
			sumMap[tempSum] = append(sumMap[tempSum], i + 1)
		} else {
			newArray := make([]int, 0)
			newArray = append(newArray, i + 1)
			sumMap[tempSum] = newArray
		}
	}
	ans := 0
	for _, v := range sumMap {
		if len(v) > 1 {
			for i := 0; i < len(v) - 1; i++ {
				for j := i + 1; j < len(v); j++ {
					ans += v[j] - v[i] - 1
				}
			}
		}
	}
	return ans
}
