package main

import (
	"Leetcode/Daily/D2021_09"
	"fmt"
)

func main() {
	ret := D2021_09.Lc916FindWords([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, []string{"SEE"})
	fmt.Println(ret)
}
