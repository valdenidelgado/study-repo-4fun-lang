package main

import "fmt"

func main() {
	x := struct {
		valor map[string]int
		slice []int
	}{
		valor: map[string]int{
			"João": 10,
			"Maria": 20,
		},
		slice: []int{1, 2, 3, 4, 5},
	}

	fmt.Println(x)

	for _, v := range x.valor {
		fmt.Println(v)
	}
	for _, v := range x.slice {
		fmt.Println(v)
	}
}