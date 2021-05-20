package main

import (
	"Leetcode/Daily"
	"fmt"
)

func main()  {
	ret := Daily.Lc520TopKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2)
	fmt.Println(ret)
}