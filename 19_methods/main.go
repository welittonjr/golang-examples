package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

func (p *Pessoa) setNome(nome string) {
	p.nome = nome
}

func (p Pessoa) getNome() string {
	return p.nome
}

func (p *Pessoa) setIdade(idade int) {
	p.idade = idade
}

func (p Pessoa) getIdade() int {
	return p.idade
}

func main() {
	pessoa1 := Pessoa{}
	pessoa2 := Pessoa{}

	pessoa1.setNome("wellington")
	pessoa1.setIdade(26)

	pessoa2.setNome("mayana")
	pessoa2.setIdade(30)

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
}
