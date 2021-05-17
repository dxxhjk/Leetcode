package main

import (
	"Leetcode/answers"
	"Leetcode/structs"
	"fmt"
)

func main()  {
	root := structs.BuildTree([]interface{}{1, 2, 3, nil, 4})
	ret := answers.Lc517IsCousins(root, 2, 3)
	fmt.Println(ret)
}