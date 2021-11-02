package D2021_09

import "Leetcode/Structs"

/*
最小k个数
 */

func Lc903SmallestK(arr []int, k int) []int {
	heap := Structs.MakeHeap(k, func(a, b interface{}) bool {return a.(int) < b.(int)})
	for _, num := range arr {
		heap.Push(num)
		heap.Show()
	}
	interfaceAns := heap.GetAll()
	ans := make([]int, 0)
	for _, num := range interfaceAns {
		ans = append(ans, num.(int))
	}
	return ans
}