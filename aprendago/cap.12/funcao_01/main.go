package main

import "fmt"

func main() {
	// basica()
	argumento("tarde")
}

func basica() {
	fmt.Println("Oi, Bom dia")
}

func argumento(s string) {
	if s == "manhã" {
		fmt.Println("Oi, bom dia!")
	} else if s == "tarde" {
		fmt.Println("Oi, boa tarde!")
	} else {
		fmt.Println("Oi, boa noite!")
	}
}
