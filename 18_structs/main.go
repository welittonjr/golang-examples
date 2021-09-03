package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

type Profissional struct {
	pessoa    Pessoa
	descricao string
	salario   int
}

func main() {
	p1 := Pessoa{
		nome:  "Wellington",
		idade: 26,
	}

	p2 := Profissional{
		pessoa: Pessoa{
			nome:  "Mayana",
			idade: 30,
		},
		descricao: "arquiteta",
		salario:   15000,
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
