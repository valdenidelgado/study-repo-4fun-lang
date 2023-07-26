package main

import "fmt"


type cliente struct {
	name string
	sobrenome string
	fumante bool
}


func main() {
	cliente1 := cliente{
		name: "João",
		sobrenome: "Silva",
		fumante: false,
	}

	cliente2 := cliente{"João", "Silva", false}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
}