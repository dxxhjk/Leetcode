package D2021_06

import "Leetcode/Structs"

/*
160. 相交链表
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。
如果两个链表没有交点，返回 null 。
 */

/*
思路：
Pa 和 Pb 两个指针同时开始遍历链表，遍历完成后指向对方的头继续遍历，
这样就可以在第二次走到相交点的时候让两个链表没有长度差。
若指向一处则相交。
若没有相交则两个指针逗为空，也是相等。
 */

func Lc604GetIntersectionNode(headA, headB *Structs.ListNode) *Structs.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}