package Structs

import "fmt"

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func(l *ListNode) Show() {
	p := l
	for p != nil {
		fmt.Print(p, " ")
		p = p.Next
	}
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
