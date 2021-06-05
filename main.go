package main

import (
	"Leetcode/Daily/D2021_06"
	"Leetcode/Structs"
	"fmt"
)

func main() {
	l1 := Structs.NewList([]interface{}{1,3,2,3,3,4,3})
	ret := D2021_06.Lc605RemoveElements(l1, 3)
	ret.Show()
	fmt.Println(ret.Val)

}