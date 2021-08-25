package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
	d = iota
)

func main() {
	fmt.Println(a, b, c, d)
}