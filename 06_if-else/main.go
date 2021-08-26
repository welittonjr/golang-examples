package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 e par")
	} else {
		fmt.Println("7 e impar")
	}

	if 8%4 == 0 {
		fmt.Println("8 eh divisivel por 4")
	}

	if num := 50; num < 0 {
		fmt.Println(num, "e negativo")
	} else if num < 10 {
		fmt.Println(num, "tem 1 digito")
	} else {
		fmt.Println(num, "tem multiplos digitos")
	}
}
