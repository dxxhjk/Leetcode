package main

import (
	"Leetcode/Daily/D2021_07"
	"encoding/json"
	"fmt"
)

func main() {
	ret := D2021_07.Lc717MaxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4})
	fmt.Println(ret)
	c := make(map[uint32]string)
	bytes := []byte(`{"11":"11","22":"b","33":"b"}`)
	json.Unmarshal(bytes, &c)
	fmt.Println(c)

	bytes1 := []byte(`{"22":"d","44":"b"}`)
	json.Unmarshal(bytes1, &c)
	fmt.Println(c)
}