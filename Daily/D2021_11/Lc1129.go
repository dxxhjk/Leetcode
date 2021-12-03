package D2021_11

import "container/heap"

/*
给你一个按递增顺序排序的数组 arr 和一个整数 k 。数组 arr 由 1 和若干素数组成，且其中所有整数互不相同。
对于每对满足 0 < i < j < arr.length 的 i 和 j ，可以得到分数 arr[i] / arr[j] 。
那么第 k 个最小的分数是多少呢?
以长度为 2 的整数数组返回你的答案, 这里 answer[0] == arr[i] 且 answer[1] == arr[j] 。
 */

func Lc1129KthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	h := make(hp, n-1)
	for j := 1; j < n; j++ {
		h[j-1] = frac{arr[0], arr[j], 0, j}
	}
	heap.Init(&h)
	for loop := k - 1; loop > 0; loop-- {
		f := heap.Pop(&h).(frac)
		if f.i+1 < f.j {
			heap.Push(&h, frac{arr[f.i+1], f.y, f.i + 1, f.j})
		}
	}
	return []int{h[0].x, h[0].y}
}

type frac struct{ x, y, i, j int }
type hp []frac
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].x*h[j].y < h[i].y*h[j].x }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(frac)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
