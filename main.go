package main

import (
	"Leetcode/Daily"
	"Leetcode/Structs"
	"fmt"
)

func main()  {
	root1 := Structs.BuildTree([]interface{}{3,5,1,6,2,9,8,nil,nil,7,4})
	root2 := Structs.BuildTree([]interface{}{3,5,1,6,7,4,2,nil,nil,nil,nil,nil,nil,9,8})
	ret := Daily.Lc510LeafSimilar(root1, root2)
	fmt.Println(ret)
}