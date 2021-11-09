package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "asd"
	fmt.Println(reflect.TypeOf(str), reflect.ValueOf(str))
	for _, b := range str {
		fmt.Println(reflect.TypeOf(b), reflect.ValueOf(b), b, byte(b))
	}
	fmt.Println(reflect.TypeOf(str[0]), reflect.ValueOf(str[0]))
	m := make(map[int]int)
	m[1] = 10
	fmt.Println(m)
	test(m)
	fmt.Println(m)
}

func test(m map[int]int) {
	m[1] -= 2
}
