package D2021_05

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

type trieNode516 struct {
	children [2]*trieNode516
}

func(t *trieNode516) add(num int) {
	p := t
	for i := 31; i >= 0; i-- {
		a := (num >> i) & 1
		if p.children[a] == nil {
			newNode := new(trieNode516)
			p.children[a] = newNode
		}
		p = p.children[a]
	}
}

func(t *trieNode516) find(num int) int {
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
	t := new(trieNode516)
	ans := 0
	for i := 0; i < len(nums); i++ {
		t.add(nums[i])
		tempXor := t.find(nums[i])
		if tempXor > ans {
			ans = tempXor
		}
	}
	return ans
}