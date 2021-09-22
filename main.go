package main

import (
	"Leetcode/Daily/D2021_09"
	"Leetcode/Structs"
	"fmt"
)

func main() {
	list := Structs.NewList([]interface{}{1, 2, 3})
	ret := D2021_09.Lc922SplitListToParts(list, 5)
	fmt.Println(ret)
}
