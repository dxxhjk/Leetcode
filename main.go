package main

import (
	"Leetcode/Daily/D2021_06"
	"Leetcode/Structs"
	"fmt"
)

func main() {
	l1 := Structs.NewList([]interface{}{1,2,3,4})
	l2 := Structs.NewList([]interface{}{5})
	l3 := Structs.NewList([]interface{}{6,7,8})
	l1.Next.Next.Next.Next = l3
	l2.Next = l3
	ret := D2021_06.Lc604GetIntersectionNode(l1, l2)
	fmt.Println(ret.Val)

}