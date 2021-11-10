package main

import "fmt"

func main() {
	str := [3]int{1,2,3}
	fmt.Println(&str[0])
	for _, v := range str {
		fmt.Println(&v, v)
	}
}
