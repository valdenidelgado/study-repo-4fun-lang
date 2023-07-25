package main

import "fmt"

func main()  {
	
	slice := []string{"mustafa", "kemal", "atatürk"}

	fmt.Println("Fatiando o slice com [:] para mostrar todos do array")
	fmt.Println(slice[:])

	fmt.Println("Usando for tradicional para mostrar todos os elementos do array")
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("Usando for range para mostrar todos os elementos do array")
	for _, v := range slice {
		fmt.Println(v)
	}
		
}