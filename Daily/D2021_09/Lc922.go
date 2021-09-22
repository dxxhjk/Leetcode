package D2021_09

import (
	"Leetcode/Structs"
)

/*
725. 分隔链表
给你一个头结点为 head 的单链表和一个整数 k ，请你设计一个算法将链表分隔为 k 个连续的部分。
每部分的长度应该尽可能的相等：任意两部分的长度差距不能超过 1 。这可能会导致有些部分为 null 。
这 k 个部分应该按照在链表中出现的顺序排列，并且排在前面的部分的长度应该大于或等于排在后面的长度。
返回一个由上述 k 部分组成的数组。
*/

func Lc922SplitListToParts(head *Structs.ListNode, k int) []*Structs.ListNode {
	length := 0
	p := head
	for p != nil {
		length++
		p = p.Next
	}
	base, extraNum := length/k, length%k
	ans := make([]*Structs.ListNode, k)
	for i := 0; i < extraNum; i++ {
		tempList := new(Structs.ListNode)
		tp := tempList
		for j := 0; j < base+1; j++ {
			tp.Next = head
			head = head.Next
			tp = tp.Next
		}
		tp.Next = nil
		ans[i] = tempList.Next
	}
	for i := extraNum; i < k; i++ {
		tempList := new(Structs.ListNode)
		tp := tempList
		for j := 0; j < base; j++ {
			tp.Next = head
			head = head.Next
			tp = tp.Next
		}
		tp.Next = nil
		ans[i] = tempList.Next
	}
	return ans
}
