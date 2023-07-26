package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
}

type profissional struct {
	pessoa
	profissao string
	salario float64
}

func main() {
	
	pessoa1 := pessoa{"João", "Silva"}
	profissional1 := profissional{pessoa1, "Programador", 1000.00}
	profissional2 := profissional{pessoa{"João", "Silva"}, "Programador", 1000.00}

	fmt.Println(pessoa1)
	fmt.Println(profissional1.nome)
	fmt.Println(profissional2.salario)
}