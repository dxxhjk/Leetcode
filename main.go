package main

import (
	"Leetcode/Daily"
	"fmt"
)

func main()  {
	ret := Daily.Lc529NumSubMatrixSumTarget([][]int{{0,1,0},{1,1,1},{0,1,0}}, 0)
	fmt.Println(ret)

}