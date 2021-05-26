package Structs

import "fmt"

type Stack []interface{}

func NewStack() *Stack {
	s := new(Stack)
	return s
}

func(s *Stack) Push(elem interface{}) {
	*s = append(*s, elem)
}

func(s *Stack) Pop() (bool, interface{}) {
	if len(*s) == 0 {
		return false, nil
	}
	ret := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s) - 1]
	return true, ret
}

func(s *Stack) Show() {
	fmt.Println(*s)
}