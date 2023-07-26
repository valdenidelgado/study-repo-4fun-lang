package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
}

func (p pessoa) returnPessoa()  {
	fmt.Println(p.nome, p.sobrenome)
}

func main() {
	pessoa1 := pessoa{"João", "Silva"}
	pessoa1.returnPessoa()
}