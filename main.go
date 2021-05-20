package main

import (
	"Leetcode/Daily"
	"fmt"
)

func main()  {
	ret := Daily.Lc512XorQueries([]int{4,8,2,10}, [][]int{{2,3},{1,3},{0,0},{0,3}})
	fmt.Println(ret)
}