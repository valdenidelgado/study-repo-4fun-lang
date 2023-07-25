package main

import "fmt"

func main() {
	
	// o make cria um slice com um tamanho inicial e um tamanho máximo
	slice := make([]int, 5, 10)

	// o tamanho inicial é o tamanho do array interno
	// o tamanho máximo é o tamanho do array interno + o tamanho do array externo
	fmt.Println(slice, len(slice), cap(slice))

	// posso acrescentar elementos no tamanho inicial e no tamanho máximo
	slice[0], slice[1], slice[2],  slice[3],  slice[4] = 1, 2, 3, 4, 5
	fmt.Println(slice, len(slice), cap(slice))

	// posso acrescentar elementos no tamanho máximo
	slice = append(slice, 6, 7, 8, 9, 10)
	fmt.Println(slice, len(slice), cap(slice))

	//ou
	slice[6], slice[7], slice[8], slice[9] = 11, 12, 13, 14
	fmt.Println(slice, len(slice), cap(slice))
}