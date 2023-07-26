package main

import "fmt"

func main() {
	// Defer geralmente deixa para executar algo no final da função
	// Exemplo:
	defer fmt.Println("Olá, mundo!")
	fmt.Println("Olá, Go!")

	// muito usado para fechar arquivos, conexões com banco de dados, etc
}