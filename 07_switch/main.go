package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("É fim de semana")
	default:
		fmt.Println("É um dia de semana")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("É antes do meio dia")
	default:
		fmt.Println("É depois do meio dia")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Bolean")
		case int:
			fmt.Println("Inteiro")
		default:
			fmt.Printf("Não sei tipo %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
