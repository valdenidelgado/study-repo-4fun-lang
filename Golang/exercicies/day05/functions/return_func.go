package main

import "fmt"

func main() {

	x := returnFunc(10)
	fmt.Println(x(2))

	// or u call the function directly like this:
	fmt.Println(returnFunc(10)(2))
}

func returnFunc(x int) func(int) int {
	return func(y int) int {
		return y * x
	}
}