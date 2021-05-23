package main

import (
	"Leetcode/Daily"
	"fmt"
)

func main()  {
	ret := Daily.Lc523MaximizeXor([]int{5,2,4,6,6,3}, [][]int{{12,4}, {8,1}, {6,3}})
	fmt.Println(ret)
}