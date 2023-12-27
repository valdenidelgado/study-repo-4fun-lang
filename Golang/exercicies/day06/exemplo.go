package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice)
	t := slice[1:3]
	fmt.Println(t)

	fmt.Println(slice)
}