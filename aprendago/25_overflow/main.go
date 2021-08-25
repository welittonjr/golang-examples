package main

import (
	"fmt"
)

func main () {
	var i uint16
	i = 65535
	fmt.Println(i) //limit maximo
	i++
	fmt.Println(i) //transbordo, zerado
	i++
	fmt.Println(i)

	//var j uint16
	//j = 65536
	//estouro de memoria
	//fmt.Println(j)
}