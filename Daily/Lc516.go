package Daily

import "fmt"

/*
421. 数组中两个数的最大异或值
给你一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。
进阶：你可以在 O(n) 的时间解决这个问题吗？
 */

/*
思路：
用一个32层的字典树存储读过的每一个数，每次读新数就去树里找异或最大值，并将自己存入字典树
 */

type TrieNode struct {
	children [2]*TrieNode
}

func(t *TrieNode) Add(num int) {
	p := t
	for i := 31; i >= 0; i-- {
		a := (num >> i) & 1
		if p.children[a] == nil {
			newNode := new(TrieNode)
			p.children[a] = newNode
		}
		p = p.children[a]
	}
}

func(t *TrieNode) Find(num int) int {
	p := t
	ans := 0
	for i := 31; i>= 0; i-- {
		a := (num >> i) & 1
		b := 1 - a
		ans <<= 1
		if p.children[b] != nil {
			ans = ans + 1
			p = p.children[b]
		} else {
			p = p.children[a]
		}
		fmt.Println(i, ans)
	}
	return ans
}

func Lc516FindMaximumXOR(nums []int) int {
	t := new(TrieNode)
	ans := 0
	for i := 0; i < len(nums); i++ {
		t.Add(nums[i])
		tempXor := t.Find(nums[i])
		if tempXor > ans {
			ans = tempXor
		}
	}
	return ans
}