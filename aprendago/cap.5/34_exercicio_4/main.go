package main

import (
	"fmt"
)

func main() {
	var val1 = 200
	fmt.Printf("%d\n%b\n%x\n", val1, val1, val1)
	val2 := val1 << 1
	fmt.Printf("%d\n%b\n%x\n", val2, val2, val2)
}