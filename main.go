package main

import (
	"Leetcode/Daily"
	"fmt"
)

func main()  {
	ret := Daily.Lc521MaxUncrossedLines([]int{2,5,1,2,5}, []int{10,5,2,1,5,2})
	fmt.Println(ret)
}