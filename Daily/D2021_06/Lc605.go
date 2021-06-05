package D2021_06

import "Leetcode/Structs"

/*
203. 移除链表元素
给你一个链表的头节点head 和一个整数 val ，
请你删除链表中所有满足 Node.val == val 的节点，
并返回新的头节点 。
 */

func Lc605RemoveElements(head *Structs.ListNode, val int) *Structs.ListNode {
	if head == nil {
		return nil
	}
	h := new(Structs.ListNode)
	h.Next = head
	p, next := h, h
	for p != nil {
		next = p.Next
		for next != nil && next.Val == val {
			next = next.Next
		}
		p.Next = next
		p = p.Next
	}
	return h.Next
}
