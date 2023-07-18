package main

import "fmt"

func main() {
	var a int
	scanln, err := fmt.Scanln(&a)
	if err != nil {
		return
	}
	var b int
	err, err := fmt.Scanln(&b)
	if err != nil {
		return
	}

	fmt.Println("PROD =", a*b)
}
