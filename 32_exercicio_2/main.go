package main

import (
	"fmt"
)

func main() {
	a := 7
	b := 5

	r1 := a == b
	r2 := a != b
	r3 := a <= b
	r4 := a < b
	r5 := a >= b
	r6 := a > b

	fmt.Printf("%d eh igual a %d = %v\n", a, b, r1)
	fmt.Printf("%d eh diferente de %d = %v\n", a, b, r2)
	fmt.Printf("%d eh menor igual a %d = %v\n", a, b, r3)
	fmt.Printf("%d eh menor que %d = %v\n", a, b, r4)
	fmt.Printf("%d eh maior igual a %d = %v\n", a, b, r5)
	fmt.Printf("%d eh maior que %d = %v\n", a, b, r6)
}