package main

import "fmt"

func main()  {
	
	slice := []string{"mustafa", "kemal", "atatürk"}
	outroSlice := []string{"bla", "two", "ble"}

	// append() é uma função nativa do Go que adiciona um elemento ao slice
	slice = append(slice, "novo elemento")

	// append() também pode adicionar mais de um elemento ao slice
	slice = append(slice, "novo elemento", "outro elemento")

	// append() também pode adicionar um slice inteiro ao slice
	// para isso, basta adicionar ... ao final do slice
	// os ... são chamados de variadic arguments que são usados para passar um número variável de argumentos para uma função
	slice = append(slice, outroSlice...)

	fmt.Println(slice)

}