package main

import "fmt"

func main() {
	x := struct {
		nome string
		sobrenome string
	}{
		nome: "João",
		sobrenome: "Silva",
	}

	fmt.Println(x)
}