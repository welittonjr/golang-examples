package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	// altera o endereço de memoria
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("inicial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// passa o endereço de memoria
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("ponteiro:", &i)
}
