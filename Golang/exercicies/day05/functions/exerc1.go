package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}
	defer fmt.Println(returnInt())
	fmt.Println(sum(1,2,3,4,5,6,7,8,9,10))
	fmt.Println(sum(slice...))
	fmt.Println(sum2(slice))

}

func returnInt() int {
	return 10
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func sum2(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
