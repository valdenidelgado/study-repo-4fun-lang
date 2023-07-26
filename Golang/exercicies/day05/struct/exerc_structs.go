package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	sabores []string
}

func main() {

	mapa := make(map[string]pessoa)
	
	pessoa1 := pessoa{
		nome: "João",
		sobrenome: "Silva",
		sabores: []string{"chocolate", "morango", "baunilha"},
	}
	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sabores {
		fmt.Println("\t", v)
	}

	mapa["Silva"] = pessoa1

	for _, v := range mapa {
		fmt.Println(v.nome, v.sobrenome)
		for _, s := range v.sabores {
			fmt.Println("\t", s)
		}
	}

}