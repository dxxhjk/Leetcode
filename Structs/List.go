package Structs

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func NewList(elements []interface{}) *ListNode {
	return NewListWithHead(elements).Next
}

func NewListWithHead(elements []interface{}) *ListNode {
	list := new(ListNode)
	p := list
	for _, element := range elements {
		newNode := new(ListNode)
		newNode.Val = element
		p.Next = newNode
		p = p.Next
	}
	return list
}
