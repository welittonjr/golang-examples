package main

import "fmt"

func main() {
	x := 0
	for {
		if x < 10 {
			x++
			fmt.Println("x e menor que 10")
		} else {
			break
		}
	}
}
