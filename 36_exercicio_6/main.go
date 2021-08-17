package main

import (
	"fmt"
)

const (
	_  = iota //0
	a1 = 2021 + iota //1
	a2 = 2021 + iota //2
	a3 = 2021 + iota //3 
	a4 = 2021 + iota //4
	a5 = 2021 + iota //5
)

func main() {
	fmt.Println(a1, a2, a3, a4, a5)
}