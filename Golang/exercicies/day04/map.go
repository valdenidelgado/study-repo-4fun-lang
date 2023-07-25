package main

import "fmt"

func main() {
	
	maps := map[string]string{"nome": "Valdeni", "sobrenome": "Delgado"}
	maps2 := make(map[string]string)

	maps2["nome"] = "Valdeni"
	maps2["sobrenome"] = "Delgado"
	maps["phone"] = "123456789"

	fmt.Println(maps)
	fmt.Println(maps["nome"])
	fmt.Println(maps2)

	for key, value := range maps {
		fmt.Println(key, value)
	}

	// para checar se uma chave existe no map
	maybe, ok := maps["email"]
	fmt.Println(maybe, ok) // ok return boolean value

	//podemos usar o if tambem para chegar e ja acrescentar uma regra
	if email, ok := maps["email"]; ok {
		fmt.Println(email)
	} else {
		fmt.Println("email não existe")
	}

	// para deletar um elemento do map
	delete(maps, "phone")
	fmt.Println(maps)
}